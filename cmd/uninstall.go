package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
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
