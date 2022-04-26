package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
)

var (
	upgrade = &cobra.Command{
		Use:   "upgrade [app-names...]",
		Short: "Upgrade specific applications",
		RunE:  doUpgrade,
	}
)

func init() {
	root.AddCommand(upgrade)
}

func doUpgrade(_ *cobra.Command, argv []string) error {
	return internal.Upgrade(argv...)
}
