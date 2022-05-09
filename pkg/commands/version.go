package commands

import (
	"fmt"

	"github.com/KusionStack/kusionup/pkg/version"
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   `Show kusionup version, alias: "v"`,
		RunE: func(c *cobra.Command, args []string) error {
			_, err := fmt.Printf("kusionup version %s\n", version.ReleaseVersion())
			return err
		},
	}
}
