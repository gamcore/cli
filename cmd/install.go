package cmd

import (
	"github.com/goo-app/cli/internal"
	"github.com/spf13/cobra"
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
