package api

import (
	"github.com/go-git/go-git/v5"
	"github.com/goo-app/cli/log"
	"github.com/goo-app/cli/utils"
	"os"
	"path"
	"strings"
)

type (
	Repos []Repo
	Repo  struct {
		Name string
	}
)

func AddRepo(name, url string) error {
	name = strings.ToLower(name)
	repoPath := path.Join(reposPath(), name)
	if err := os.MkdirAll(repoPath, os.ModeDir); os.IsExist(err) {
		return ErrRepoAlreadyExist
	}
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}
	return nil
}

func RemoveRepo(name string) error {
	if utils.NoneSlice(Repositories(), func(r Repo) bool {
		return r.Name == strings.ToLower(name)
	}) {
		return ErrRepoNotExist
	}
	repoPath := path.Join(reposPath(), strings.ToLower(name))
	return os.RemoveAll(repoPath)
}

func Repositories() Repos {
	repos := Repos{}
	if reposDirs, err := os.ReadDir(reposPath()); err != nil {
		log.Error(err.Error())
	} else {
		for _, d := range reposDirs {
			if d.IsDir() {
				repos = append(repos, Repo{Name: d.Name()})
			}
		}
	}
	return repos
}

func Repository(name string) *Repo {
	return utils.FindFirstSlice(Repositories(), func(r Repo) bool {
		return r.Name == name
	})
}

func (r Repo) Path() string {
	return r.repoPath()
}

func reposPath() string {
	return path.Join(Path, "repos")
}

func (r Repo) repoPath() string {
	return path.Join(reposPath(), r.Name)
}
