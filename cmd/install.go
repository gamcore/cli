package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
)

var (
	install = &cobra.Command{
		Use:   "install [app-name...]",
		Short: "Install applications",
		RunE:  doInstall,
	}
)

func init() {
	root.AddCommand(install)
}

func doInstall(_ *cobra.Command, argv []string) error {
	return internal.Install(argv...)
}
