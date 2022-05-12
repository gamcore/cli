package cmd

import (
	"fmt"
	"github.com/goo-app-manager/goo/api"
	"github.com/spf13/cobra"
	"net/url"
)

var (
	repo = &cobra.Command{
		Use:   "repo",
		Short: "Repository containment",
	}
	repoAdd = &cobra.Command{
		Use:   "add [name] [git-url]",
		Short: "add repository",
		RunE:  doAddRepo,
	}
	repoDelete = &cobra.Command{
		Use:     "delete [name]",
		Aliases: []string{"remove"},
		Short:   "Delete repository",
		Long:    "Only if none apps are installed from specific repository",
		RunE:    doDelRepo,
	}
)

func init() {
	repo.AddCommand(repoAdd, repoDelete)
	root.AddCommand(repo)
}

func doAddRepo(_ *cobra.Command, argv []string) error {
	name := argv[0]
	gitUrl := argv[1]

	if _, err := url.Parse(name); err == nil {
		return fmt.Errorf("first argument !!!MUST BE!!! a name of repository")
	}
	if _, err := url.Parse(gitUrl); err != nil {
		return fmt.Errorf("second argument !!!MUST BE!!! a git repository url")
	}

	return api.AddRepo(name, gitUrl)
}

func doDelRepo(_ *cobra.Command, argv []string) error {
	name := argv[0]
	return api.RemoveRepo(name)
}
