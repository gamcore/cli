package internal

import (
	"github.com/goo-app/cli/internal/logger"
	"github.com/goo-app/cli/internal/utils"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func Cleanup(apps ...string) error {
	installed := GetInstalledApplications()

	if apps[0] != "*" {
		installed = utils.FilterSlice(installed, func(a Application) bool {
			for _, app := range apps {
				if strings.ToLower(app) == strings.ToLower(a.Name) {
					return true
				}
			}
			return false
		})
	}

	for _, a := range installed {
		if err := a.Clean(); err != nil {
			return err
		}
	}

	return doCleanupCache()
}

func doCleanupCache() error {
	cachePath := path.Join(Path, "tmp")
	cacheSlice, err := ioutil.ReadDir(cachePath)
	if err != nil {
		return err
	}
	logger.Debug("Cleaning temp files")
	for _, f := range cacheSlice {
		err = os.RemoveAll(path.Join(cachePath, f.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}
