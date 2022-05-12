package cmd

import (
	"github.com/goo-app-manager/goo/api"
	"github.com/spf13/cobra"
)

var (
	cleanup = &cobra.Command{
		Use:   "cleanup [app-names..]",
		Short: "Clean unused installations and caches",
		RunE:  doCleanup,
	}
	cache = false
)

func init() {
	cleanup.Flags().BoolVarP(&cache, "cache", "c", false, "Cleanup cache")
	root.AddCommand(cleanup)
}

func doCleanup(_ *cobra.Command, argv []string) error {
	return api.Cleanup(argv, cache)
}
