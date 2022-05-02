package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/spf13/cobra"
)

var (
	cleanup = &cobra.Command{
		Use:   "cleanup [app-names..]",
		Short: "Clean unused installations and caches",
		RunE:  doCleanup,
	}
)

func init() {
	root.AddCommand(cleanup)
}

func doCleanup(_ *cobra.Command, argv []string) error {
	return internal.Cleanup(argv...)
}
