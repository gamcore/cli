package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/goo-app/cli/internal/logger"
	"github.com/goo-app/cli/internal/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	list = &cobra.Command{
		Use:   "list",
		Short: "List installed applications",
		Run:   doList,
	}
)

func init() {
	root.AddCommand(list)
}

func doList(_ *cobra.Command, _ []string) {
	apps := internal.GetInstalledApplications()
	var data [][]string
	utils.ForEachSlice(apps, func(app internal.Application) {
		version, _ := app.CurrentVersion()
		name := app.Name
		hasUpdate, err := app.HasUpdate()
		if err != nil {
			logger.ErrorF(`could not obtain update information for "%s": %s`, name, err)
		} else {
			if hasUpdate {
				name = pterm.Italic.Sprint(name)
			}
		}
		data = append(data, []string{
			name, *version,
		})
	})
	_ = pterm.DefaultTable.WithData(data).Render()
}
