package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stachu540/goo/internal"
)

var (
	update = &cobra.Command{
		Use:   "update",
		Short: "Update repository indexes",
		RunE:  doUpdate,
	}
)

func init() {
	root.AddCommand(update)
}

func doUpdate(_ *cobra.Command, _ []string) error {
	return internal.Update()
}
