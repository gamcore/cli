package internal

import (
	"fmt"
	"github.com/stachu540/goo/internal/logger"
	"github.com/stachu540/goo/internal/utils"
	"os"
	"path"
)

func Setup(level logger.Level) error {
	gooPath, exist := os.LookupEnv("GOO_PATH")
	if exist {
		Path = gooPath
		if _, err := os.Stat(gooPath); os.IsExist(err) {
			logger.InitLogger(level, Path)
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

func registerDefaultRepos() error {
	coreRepo := path.Join(Path, "repos", "core")
	if _, err := os.Stat(coreRepo); os.IsExist(err) {
		if err = os.RemoveAll(coreRepo); err != nil {
			return err
		}
	}

	return AddRepository("core", fmt.Sprintf("https://github.com/%s.git", CoreRepoSlug))
}
