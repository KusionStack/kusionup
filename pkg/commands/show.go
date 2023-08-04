package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func showCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show installed Kusion",
		Long:  "Show installed Kusion versions.",
		RunE:  runShow,
	}
}

func runShow(_ *cobra.Command, _ []string) error {
	vers, err := listKusionVers()
	if err != nil {
		return err
	}

	if len(vers) == 0 {
		showKusionIfExist()
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Version", "Active"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, ver := range vers {
		if ver.Current {
			table.Append([]string{ver.Ver, "*"})
		} else {
			table.Append([]string{ver.Ver, ""})
		}
	}

	table.Render()

	return nil
}

func showKusionIfExist() {
	goBin, err := exec.LookPath("go")
	if err == nil {
		fmt.Printf("No Kusion is installed by kusionup. Using system Kusion %q.\n", goBin)
	} else {
		fmt.Println("No Kusion is installed by kusionup.")
	}
}

type goVer struct {
	Ver     string
	Current bool
}

func listKusionVers() ([]goVer, error) {
	files, err := os.ReadDir(KusionupDir())
	if err != nil {
		return nil, err
	}

	current, err := currentKusionVersion()
	if err != nil {
		return nil, err
	}

	var vers []goVer

	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), "kusion") {
			if _, err := os.Stat(KusionupDir(file.Name(), unpackedOkay)); err != nil {
				// only list successfully unpacked verions
				continue
			}
			vers = append(vers, goVer{
				Ver:     strings.TrimPrefix(file.Name(), "kusion-"),
				Current: current == file.Name(),
			})
		}
	}

	return vers, nil
}

func currentKusionVersion() (string, error) {
	current := KusionupCurrentDir()

	goroot, err := os.Readlink(current)
	if err != nil {
		return "", err
	}

	return filepath.Base(goroot), nil
}
