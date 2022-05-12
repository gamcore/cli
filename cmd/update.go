package cmd

import (
	"github.com/goo-app/cli/api"
	"github.com/spf13/cobra"
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
	return api.Update()
}
