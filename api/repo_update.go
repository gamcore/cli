package api

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (r Repo) Update() error {
	repo, err := git.PlainOpen(r.Path())
	if err != nil {
		return fmt.Errorf(`could not update repository "%s": %s`, r.Name, err)
	}

	tree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf(`could not update repository "%s": %s`, r.Name, err)
	}

	return tree.Pull(&git.PullOptions{RemoteName: "origin"})
}
