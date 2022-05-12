package api

import (
	"github.com/go-git/go-git/v5"
	"os"
	"path"
)

func AddRepository(name, url string) error {
	path.Join(reposPath(), name)
	_, err := git.PlainClone(path.Join(reposPath(), name), false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}
