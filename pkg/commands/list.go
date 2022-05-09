package commands

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "ls-ver [REGEXP]",
		Aliases: []string{"lsv"},
		Short:   `List Kusion versions to install, alias: "lsv"`,
		Long: `List available Kusion versions matching a regexp filter for installation. If no filter is provided,
list all available versions.`,
		Example: `
  kusionup ls-ver
  kusionup ls-ver latest
  kusionup ls-ver 0.2
`,
		RunE: runList,
	}
}

func runList(cmd *cobra.Command, args []string) error {
	var regexp string
	if len(args) > 0 {
		regexp = args[0]
	}

	vers, err := listKusionVersions(regexp)
	if err != nil {
		return err
	}

	for _, ver := range vers {
		fmt.Println(ver)
	}

	return nil
}

func listKusionVersions(re string) ([]string, error) {
	if re == "" {
		re = ".+"
	} else {
		re = fmt.Sprintf(`.*%s.*`, re)
	}

	r := regexp.MustCompile(re)
	vers := []string{}

	for _, rs := range registedReleaseSources {
		for _, ver := range rs.GetVersions() {
			title := GetSourceVersionTitle(rs.GetName(), ver)
			if r.Match([]byte(title)) {
				vers = append(vers, title)
			}
		}
	}

	return vers, nil
}
