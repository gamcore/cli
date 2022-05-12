package cmd

import (
	"github.com/goo-app/cli/api"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:    "init",
		Short:  "Initialize application",
		Hidden: true,
		RunE:   doInit,
	}
	force bool
)

func init() {
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force initialization")

	root.AddCommand(initCmd)
}

func doInit(_ *cobra.Command, _ []string) error {
	return api.Init(force)
}
