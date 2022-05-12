package api

import (
	"os"
	"path"
	"runtime"
)

func (a App) Uninstall() error {
	return a.actionUninstall(func() error {
		err := os.RemoveAll(a.appPath())
		if err != nil {
			return err
		}
		return a.purgeLink()
	})
}

func (a App) Clean() error {
	cv, err := a.currentLinkedAppPath()
	if err != nil {
		return err
	}
	dirs, err := os.ReadDir(a.appPath())
	if err != nil {
		return err
	}
	for _, d := range dirs {
		if d.Name() != "current" || d.Name() != path.Base(cv) {
			// log removal
			err = os.RemoveAll(path.Join(a.appPath(), d.Name()))
			if err != nil {
				return err
			}
			// log removal ok
		}
	}
	return nil
}

func (a App) purgeLink() error {
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
		err = os.Remove(linkName)
		if err != nil {
			return err
		}
	}
	return nil
}
