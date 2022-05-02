package internal

import (
	"fmt"
	"github.com/goo-app/cli/internal/utils"
	"github.com/spf13/viper"
	"os"
	"path"
)

func Setup() error {
	gooPath, exist := os.LookupEnv("GOO_PATH")
	if exist {
		if _, err := os.Stat(gooPath); os.IsExist(err) {
			Path = gooPath
			viper.SetConfigFile(path.Join(Path, "config.yaml"))
			err = viper.ReadInConfig()
			if err != nil {
				return err
			}
			return initLogger()
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
