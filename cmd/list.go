package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
	"github.com/stachu540/goo/internal/logger"
	"github.com/stachu540/goo/internal/utils"
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
			logger.Errorf(`could not obtain update information for "%s": %s`, name, err)
		} else {
			if hasUpdate {
				name = pterm.Bold.Sprint(name)
			}
		}
		data = append(data, []string{
			name, *version,
		})
	})
	_ = pterm.DefaultTable.WithData(data).Render()
}
