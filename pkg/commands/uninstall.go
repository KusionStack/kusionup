package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var uninstallCmdIgnoreHint bool

func uninstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "uninstall <VERSION>...",
		Aliases: []string{"ui"},
		Short:   `Uninstall Kusion with a version, alias: "ui"`,
		Long:    "Uninstall Kusion by providing a version.",
		Example: `
  kusionup uninstall open@latest
  kusionup uninstall open@v0.2.15
`,
		RunE: runUninstall,
	}

	cmd.Flags().BoolVar(&uninstallCmdIgnoreHint, "ignore-hint", false, "Ignore uninstall hint")

	return cmd
}

func runUninstall(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no version is specified")
	}

	for _, ver := range args {
		if ignoreHint, err := cmd.Flags().GetBool("ignore-hint"); err == nil && !ignoreHint {
			logger.Printf("Uninstalling %s", ver)
		}

		if err := os.RemoveAll(KusionupVersionDir(ver)); err != nil {
			return err
		}
	}

	return nil
}
