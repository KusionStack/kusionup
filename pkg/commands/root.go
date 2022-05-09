package commands

import (
	"os"
	"path/filepath"

	"github.com/KusionStack/kusionup/pkg/sources"
	"github.com/KusionStack/kusionup/pkg/sources/open"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	homedir string
	logger  *logrus.Logger

	ProfileFiles []string

	rootCmdVerboseFlag     bool
	registedReleaseSources map[string]sources.ReleaseSource
)

func init() {
	logger = logrus.New()

	// Init release sources
	registedReleaseSources = map[string]sources.ReleaseSource{
		open.OpenReleaseSource.GetName(): open.OpenReleaseSource,
	}

	var err error
	homedir, err = os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
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

func preRunRoot(cmd *cobra.Command, args []string) error {
	if rootCmdVerboseFlag {
		logger.SetLevel(logrus.DebugLevel)
	}

	return nil
}
