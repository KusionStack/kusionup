package commands

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	// KusionupEnvFileContent is the content of the kusionup environment file
	KusionupEnvFileContent = `export PATH=$HOME/.kusionup/bin:$HOME/.kusionup/current/bin:$HOME/.kusionup/current/kclvm/bin:$PATH
export KUSION_PATH=$HOME/.kusionup/current`
	// ProfileFileSourceContent is the content of the kusionup profile file
	ProfileFileSourceContent = `source "$HOME/.kusionup/env"`

	welcomeTmpl = `Welcome to kusionup!

Kusion installed through kusionup will be located at:

  {{ .KusionupDir }}

To get started you need Kusion's bin directory 
({{ .CurrentKusionBinDir }})
in your PATH environment variable. These two paths will be added to your PATH 
environment variable by modifying the profile files located at:
{{ range $index, $elem :=.ProfileFiles }}
  {{ $elem -}}
{{ end }}

Next time you log in this will be done automatically. To configure your current 
shell run source $HOME/.bash_profile.
`
)

var (
	initCmdSkipInstallFlag bool
	initCmdSkipPromptFlag  bool
)

func initCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the kusionup environment file",
		RunE:  runInit,
	}

	initCmd.PersistentFlags().BoolVar(&initCmdSkipInstallFlag, "skip-install", false, "Skip installing Kusion")
	initCmd.PersistentFlags().BoolVar(&initCmdSkipPromptFlag, "skip-prompt", false, "Skip confirmation prompt")

	return initCmd
}

func runInit(cmd *cobra.Command, args []string) error {
	tmpl, err := template.New("").Parse(welcomeTmpl)
	if err != nil {
		return err
	}

	params := struct {
		KusionupDir         string
		KusionupBinDir      string
		CurrentKusionBinDir string
		ProfileFiles        []string
	}{
		KusionupDir:         KusionupDir(),
		KusionupBinDir:      KusionupBinDir(),
		CurrentKusionBinDir: KusionupCurrentBinDir(),
		ProfileFiles:        ProfileFiles,
	}
	if err := tmpl.Execute(os.Stdout, params); err != nil {
		return err
	}

	if !initCmdSkipPromptFlag {
		// Add a line break
		fmt.Println("")

		prompt := promptui.Prompt{
			Label:     "Would you like to proceed with the installation",
			IsConfirm: true,
		}

		if _, err := prompt.Run(); err != nil {
			return fmt.Errorf("interrupted")
		}
	}

	ef := KusionupEnvFile()
	if err := os.MkdirAll(filepath.Dir(ef), 0o755); err != nil {
		return err
	}

	// ignore error, similar to rm -f
	os.Remove(ef)

	if err := ioutil.WriteFile(ef, []byte(KusionupEnvFileContent), 0o664); err != nil {
		return err
	}

	if err := appendSourceToProfiles(ProfileFiles); err != nil {
		return err
	}

	if !initCmdSkipInstallFlag {
		// Add a line break
		fmt.Printf("\nNow, start installing the latest version of Kusion (internal@latest):\n")

		if err := runInstall(cmd, args); err != nil {
			return err
		}
	}

	return nil
}

func appendSourceToProfiles(profiles []string) error {
	for _, profile := range profiles {
		if err := appendToFile(profile, ProfileFileSourceContent); err != nil {
			return err
		}
	}

	return nil
}

func appendToFile(filename, value string) error {
	ok, err := checkStringExistsFile(filename, value)
	if err != nil {
		return err
	}

	if ok {
		return nil
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString("\n" + value + "\n"); err != nil {
		return err
	}

	return err
}

func checkStringExistsFile(filename, value string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0o600)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == value {
			return true, nil
		}
	}

	return false, scanner.Err()
}
