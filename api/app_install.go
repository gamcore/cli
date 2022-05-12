package api

import (
	"github.com/goo-app/cli/utils"
	"path"
	"runtime"
)

func (a App) Install() error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	v, link, err := mf.CheckUpdate()
	if err != nil {
		return err
	}
	actualPath := path.Join(a.appPath(), v)
	return a.actionInstall(v, func() error {
		err = a.doDownloadAndExtract(link, actualPath)
		if err != nil {
			return err
		}
		err = utils.MkLink(actualPath, a.currentAppPath())
		if err != nil {
			return err
		}
		return a.createLink()
	})
}

func (a App) createLink() error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	for _, x := range mf.Executable {
		execLink := path.Join(a.currentAppPath(), x)
		if runtime.GOOS == "windows" {
			execLink = execLink + ".exe"
		}
		baseName := path.Base(execLink)
		linkName := path.Join(binaryPath(), baseName)
		err = utils.MkLink(execLink, linkName)
		if err != nil {
			return err
		}
	}
	return nil
}
