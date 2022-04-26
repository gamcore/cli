package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/stachu540/goo/internal/log"
	"github.com/stachu540/goo/internal/utils"
	"io"
	"os"
	"path"
	"strings"
)

func Setup(level log.Level) error {
	gooPath, exist := os.LookupEnv("GOO_PATH")
	if exist {
		lf := path.Join(gooPath, "goo.log")
		if _, err := os.Stat(gooPath); os.IsExist(err) {
			log.New(level, lf)
			return nil
		} else {
			return err
		}
	}
	return ErrNotInitialized
}

func Init(force bool) error {
	init := force

	if _, err := os.Stat(path.Join(Path, "repos")); !force || os.IsExist(err) {
		if empty, _ := utils.IsDirEmpty(path.Join(Path, "repos")); empty {
			init = true
		}
	}

	if init {
		return registerDefaultRepos()
	} else {
		return ErrAlreadyInitialized
	}
}

func GetRepositories() (repos Repositories) {
	dirs, _ := os.ReadDir(path.Join(Path, "repos"))
	for _, dir := range dirs {
		if dir.IsDir() {
			r, err := Read(Path, dir.Name())
			if err != nil {
				Logger.Errorf("error reading repository %s", dir.Name(), err)
			} else {
				repos = append(repos, *r)
			}
		}
	}
	return
}

func GetApplications() (apps Applications) {
	for _, r := range GetRepositories() {
		apps = append(apps, r.Applications()...)
	}
	return
}

func GetApplication(name string) (*Application, error) {
	for _, a := range GetApplications() {
		if strings.ToLower(a.Name) == strings.ToLower(name) {
			return &a, nil
		}
	}

	return nil, fmt.Errorf(`could not find a app name "%s"`, name)
}

func Update() error {
	for _, repo := range GetRepositories() {
		err := repo.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func Install(apps ...string) error {
	toInstall := GetApplications().Fetch(apps...)
	for _, a := range toInstall {
		err := a.Install()
		if err != nil {
			if err == ErrAlreadyInstalled {
				if len(apps) == 1 {
					return err
				}
				Logger.Warnf(`"%s" is already installed`, a.Name)
			} else {
				return err
			}
		}
	}
	return nil
}

func Upgrade(apps ...string) error {
	toUpgrade := utils.FilterSlice(GetApplications(), func(a Application) bool {
		return a.IsInstalled()
	})

	if apps[0] != "*" {
		toUpgrade = utils.FilterSlice(toUpgrade, func(a Application) bool {
			for _, app := range apps {
				if strings.ToLower(app) == strings.ToLower(a.Name) {
					return true
				}
			}
			return false
		})
	}

	for _, a := range toUpgrade {
		err := a.Update()
		if err != nil {
			if err != ErrIsUpToDate {
				return err
			}
		}
	}
	return nil
}

func Uninstall(apps ...string) error {
	toUninstall := utils.FilterSlice(GetApplications(), func(a Application) bool {
		if a.IsInstalled() {
			for _, app := range apps {
				if strings.ToLower(app) == strings.ToLower(a.Name) {
					return true
				}
			}
		}
		return false
	})

	for _, a := range toUninstall {
		err := a.Uninstall()
		if err != nil {
			return err
		}
	}

	return nil
}

func AddRepository(name, url string) error {
	repoPath := path.Join(Path, "repos", name)
	if err := os.Mkdir(repoPath, os.ModeDir); os.IsNotExist(err) {
		_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      url,
			Progress: io.MultiWriter(Logger.Out(), Logger.LogFile()),
		})
		return err
	}

	return ErrRepositoryAlreadyExist
}

func registerDefaultRepos() error {
	coreRepo := path.Join(Path, "repos", "core")
	if _, err := os.Stat(coreRepo); os.IsExist(err) {
		if err = os.RemoveAll(coreRepo); err != nil {
			return err
		}
	}

	return AddRepository("core", fmt.Sprintf("https://github.com/%s.git", CoreRepoSlug))
}
