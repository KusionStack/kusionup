package commands

import (
	"os"
	"path/filepath"

	"github.com/KusionStack/kusionup/pkg/sources"
	"github.com/KusionStack/kusionup/pkg/sources/cdn"
	"github.com/KusionStack/kusionup/pkg/sources/github"
	"github.com/KusionStack/kusionup/third_party/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	homedir string
	logger  *logrus.Logger

	ProfileFiles []string

	rootCmdVerboseFlag     bool
	customSourcesFile      string
	registedReleaseSources map[string]sources.ReleaseSource
)

func init() {
	logger = logrus.New()

	var err error

	homedir, err = os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}

	// Init release sources
	registedReleaseSources = map[string]sources.ReleaseSource{
		github.GithubReleaseSource.GetName(): github.GithubReleaseSource,
		cdn.CDNReleaseSource.GetName():       cdn.CDNReleaseSource,
	}

	// Init custom sources from configuration file
	defaultCustomSourcesFile := KusionupDir(DefaultCustomSourceConfigFilename)
	if config.IsValidConfigFilename(defaultCustomSourcesFile) {
		sources := &CustomSources{}
		if err := config.FromFile(afero.NewOsFs(), defaultCustomSourcesFile, sources); err == nil {
			for i := 0; i < len(sources.Sources); i++ {
				registedReleaseSources[sources.Sources[i].Name] = &sources.Sources[i]
			}
		}
	}

	ProfileFiles = []string{
		filepath.Join(homedir, ".profile"),
		filepath.Join(homedir, ".zprofile"),
		filepath.Join(homedir, ".bash_profile"),
		filepath.Join(homedir, ".zshrc"),
	}
}

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "kusionup",
		Short:             "The Kusion installer",
		PersistentPreRunE: preRunRoot,
		RunE:              runDefault,
	}

	rootCmd.PersistentFlags().BoolVarP(&rootCmdVerboseFlag, "verbose", "v", false, "Verbose")
	rootCmd.PersistentFlags().StringVarP(&customSourcesFile, "custom-sources-file", "c", "", "Custom sources file")

	rootCmd.AddCommand(defaultCmd())
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(installCmd())
	rootCmd.AddCommand(uninstallCmd())
	rootCmd.AddCommand(reinstallCmd())
	rootCmd.AddCommand(showCmd())
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(versionCmd())

	return rootCmd
}

func preRunRoot(_ *cobra.Command, _ []string) error {
	if rootCmdVerboseFlag {
		logger.SetLevel(logrus.DebugLevel)
	}

	if customSourcesFile != "" {
		if config.IsValidConfigFilename(customSourcesFile) {
			sources := &CustomSources{}

			err := config.FromFile(afero.NewOsFs(), customSourcesFile, sources)
			if err != nil {
				logger.Printf("Failed to load custom sources from %s: %s\n", customSourcesFile, err)
			} else {
				for i := 0; i < len(sources.Sources); i++ {
					registedReleaseSources[sources.Sources[i].Name] = &sources.Sources[i]
				}
			}
		} else {
			logger.Printf("Invalid custom sources file: %s\n", customSourcesFile)
		}
	}

	return nil
}
