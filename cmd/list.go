package cmd

import (
	"github.com/goo-app-manager/goo/api"
	"github.com/goo-app-manager/goo/log"
	"github.com/goo-app-manager/goo/utils"
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
	apps := api.Applications().Installed()
	var data = [][]string{{"Name", "Version"}}
	utils.ForEachSlice(apps, func(app api.App) {
		version, _ := app.CurrentVersion()
		if version == "" {
			version = "???"
		}
		name := app.Name
		hasUpdate, err := app.HasUpdate()
		if err != nil {
			log.ErrorF(`could not obtain update information for "%s": %s`, name, err)
		} else {
			if hasUpdate {
				version = pterm.Bold.Sprint(version)
			}
		}
		data = append(data, []string{
			name, version,
		})
	})
	_ = pterm.DefaultTable.WithData(data).Render()
}
