package commands

import (
	"github.com/spf13/cobra"
)

func reinstallCmd() *cobra.Command {
	reinstallCmd := &cobra.Command{
		Use:     "reinstall [VERSION]",
		Aliases: []string{"ri"},
		Short:   `Reinstall Kusion with a version, alias: "ri"`,
		Long:    `Reinstall Kusion by providing a version. If no version is provided, reinstall the latest Kusion (internal@latest).`,
		Example: `
  kusionup reinstall
  kusionup reinstall internal@1.15.2
`,
		PersistentPreRunE: preRunReinstall,
		RunE:              runReinstall,
		Args:              cobra.MinimumNArgs(1),
	}

	return reinstallCmd
}

func preRunReinstall(cmd *cobra.Command, args []string) error {
	return installCmd().PersistentPreRunE(cmd, args)
}

func runReinstall(cmd *cobra.Command, args []string) error {
	logger.Printf("Reinstalling %s ...", args[0])

	uninstallCmd := uninstallCmd()

	err := uninstallCmd.Flags().Lookup("ignore-hint").Value.Set("true")
	if err != nil {
		return err
	}

	err = uninstallCmd.RunE(uninstallCmd, args)
	if err != nil {
		return err
	}

	return installCmd().RunE(cmd, args)
}
