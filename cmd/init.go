package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
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
	return internal.Init(force)
}
