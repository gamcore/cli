package cmd

import (
	"github.com/goo-app-manager/goo/api"
	"github.com/spf13/cobra"
	"os"
	"path"
)

var (
	manage = &cobra.Command{
		Use:   "manage",
		Short: "Repository containment management",
		PersistentPreRunE: func(cmd *cobra.Command, argv []string) error {
			base := path.Base(cwd)
			cwdRepo = &api.RepoManager{Repo: api.Repo{Name: base}}
			return nil
		},
	}
	manageAdd = &cobra.Command{
		Use:     "add [name]",
		Aliases: []string{"plus", "+", "a"},
		Short:   "add application",
		RunE:    doManageAddApp,
	}
	manageDelete = &cobra.Command{
		Use:     "delete [name]",
		Aliases: []string{"remove", "d", "-"},
		Short:   "delete application",
		RunE:    doManageDelApp,
	}
	manageVerify = &cobra.Command{
		Use:     "verify",
		Aliases: []string{"v", "check"},
		Short:   "Validate current working directory repository",
		RunE:    doManageVerify,
	}
	manageInit = &cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		Hidden:  true,
		Short:   "Initialize new repository containment",
		RunE:    doManageInit,
	}
	cwdRepo *api.RepoManager
	cwd     string
)

func init() {
	cwd, _ = os.Getwd()
	manageInit.Flags().StringVarP(&cwd, "path", "p", cwd, "Repository path")
	manage.AddCommand(manageAdd, manageDelete, manageVerify, manageInit)
	root.AddCommand(manage)
}

func doManageVerify(_ *cobra.Command, _ []string) error {
	return cwdRepo.Validate()
}

func doManageAddApp(_ *cobra.Command, _ []string) error {
	return cwdRepo.AddApplication()
}

func doManageDelApp(_ *cobra.Command, argv []string) error {
	return cwdRepo.RemoveApplications(argv...)
}

func doManageInit(_ *cobra.Command, _ []string) error {
	return cwdRepo.Init()
}
