package cmd

import (
	"github.com/goo-app-manager/goo/api"
	"github.com/goo-app-manager/goo/log"
	"github.com/goo-app-manager/goo/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	search = &cobra.Command{
		Use:   "search [query]",
		Short: "Search applications",
		RunE:  doSearch,
	}
	regex = false
)

func init() {
	search.Flags().BoolVarP(&regex, "regex", "X", false, "Regular Expression query")
	root.AddCommand(search)
}

func doSearch(_ *cobra.Command, argv []string) error {
	apps := api.Repositories().Applications().Filter(argv[0], regex)

	if len(apps) > 0 {
		log.Info(`List of apps:`)
		err := pterm.NewBulletListFromStrings(utils.MapSlice(apps, func(app api.App) string {
			return app.Name
		}), " ").Render()
		if err != nil {
			return err
		}
	} else {
		log.Warn("No entries have found")
	}
	return nil
}
