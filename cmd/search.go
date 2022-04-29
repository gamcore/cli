package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
	"github.com/stachu540/goo/internal/logger"
	"github.com/stachu540/goo/internal/utils"
)

var (
	search = &cobra.Command{
		Use:   "search [query]",
		Short: "Search applications",
		Run:   doSearch,
	}
	regex = false
)

func init() {
	search.Flags().BoolVarP(&regex, "regex", "X", false, "Regular Expression query")
	root.AddCommand(search)
}

func doSearch(_ *cobra.Command, argv []string) {
	apps := internal.GetApplications().Find(argv[0], regex)

	if len(apps) > 0 {
		logger.Info(`List of apps:`)
		pterm.NewBulletListFromStrings(utils.MapSlice(apps, func(app internal.Application) string {
			return app.Name
		}), " ")
	} else {
		logger.Warn("No entries have found")
	}
}
