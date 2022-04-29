package internal

import (
	"github.com/go-git/go-git/v5"
	"github.com/stachu540/goo/internal/logger"
	"os"
	"path"
)

func AddRepository(name, url string) error {
	repoPath := path.Join(Path, "repos", name)
	if err := os.Mkdir(repoPath, os.ModeDir); os.IsNotExist(err) {
		var out = logger.Out()
		_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      url,
			Progress: out,
		})
		return err
	}

	return ErrRepositoryAlreadyExist
}

func GetRepositories() (repos Repositories) {
	dirs, _ := os.ReadDir(path.Join(Path, "repos"))
	for _, dir := range dirs {
		if dir.IsDir() {
			r, err := Read(Path, dir.Name())
			if err != nil {
				logger.Errorf("error reading repository %s", dir.Name(), err)
			} else {
				repos = append(repos, *r)
			}
		}
	}
	return
}

func RemoveRepository(name string) error {
	for _, repo := range GetRepositories() {
		if repo.Name == name {
			return os.RemoveAll(repo.Path())
		}
	}

	return ErrRepositoryNotExist
}
