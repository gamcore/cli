package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/spf13/cobra"
)

var (
	uninstall = &cobra.Command{
		Use:   "uninstall [app-name...]",
		Short: "Uninstall specific applications",
		RunE:  doUninstall,
	}
)

func init() {
	root.AddCommand(uninstall)
}

func doUninstall(_ *cobra.Command, argv []string) error {
	return internal.Uninstall(argv...)
}
