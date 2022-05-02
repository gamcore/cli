package internal

import (
	"encoding/json"
	"github.com/goo-app/cli/internal/logger"
	"os"
	"path"
)

type Application struct {
	Name       string
	repository Repository
}

func (a Application) Manifest() (manifest AManifest) {
	mFile := path.Join(a.repository.Path(), "apps", a.Name+".json")
	mFileData, _ := os.ReadFile(mFile)
	err := json.Unmarshal(mFileData, &manifest)
	if err != nil {
		logger.DebugF("error read manifest file: %s", err)
	}
	return
}

func (a Application) Repository() Repository {
	return a.repository
}

func (a Application) IsInstalled() bool {
	_, err1 := os.Stat(a.path())
	_, err2 := os.Stat(path.Join(a.path(), "current"))

	return os.IsExist(err1) && os.IsExist(err2)
}

func (a Application) Path() *string {
	var p = a.path()
	if a.IsInstalled() {
		return &p
	} else {
		return nil
	}
}

func (a Application) path() string {
	return path.Join(Path, "apps", a.Name)
}
