package api

import (
	"github.com/goo-app/cli/utils"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func Cleanup(apps []string, cache bool) error {
	var appList Apps
	if len(apps) == 1 && apps[0] == "*" {
		appList = Applications().Installed()
	} else {
		appList = utils.FilterSlice(appList, func(a App) bool {
			return utils.AnySlice(apps, func(app string) bool {
				return a.Name == app
			})
		})
	}

	for _, a := range appList {
		appV, err := os.ReadDir(a.appPath())
		if err != nil {
			return err
		}
		linked, err := a.currentLinkedAppPath()
		if err != nil {
			return err
		}
		linked, err = filepath.Abs(linked)
		if err != nil {
			return err
		}
		for _, d := range appV {
			cd := path.Join(a.appPath(), d.Name())
			cd, err = filepath.Abs(cd)
			if err != nil {
				return err
			}
			if d.Name() != "current" && cd != linked {
				err = os.RemoveAll(cd)
				if err != nil {
					return err
				}
			}
		}
	}

	if cache {
		entries, err := os.ReadDir(tempPath())
		if err != nil {
			return err
		}
		for _, e := range entries {
			err = os.RemoveAll(path.Join(tempPath(), e.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Init(force bool) error {
	shellCmd := path.Join(binaryPath(), "goo")
	configPath := path.Join(Path, "config.yaml")
	if runtime.GOOS == "windows" {
		shellCmd = shellCmd + ".cmd"
	}
	_, errShell := os.Stat(shellCmd)
	_, errCfg := os.Stat(configPath)
	if os.IsExist(errShell) && os.IsExist(errCfg) && Repository("main") == nil && !force {
		return ErrAlreadyInitialized
	}
	dataCfg, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(configPath, dataCfg, os.ModeType)
	if err != nil {
		return err
	}
	err = AddRepository("main", "https://github.com/goo-app/package-main.git")
	if err != nil {
		return err
	}
	shellData := shellScript()
	if _, err = os.Stat(shellCmd); os.IsExist(err) {
		// update new shell script
		err = os.Remove(shellCmd)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(shellCmd, []byte(shellData), 0544)
}

func Install(apps ...string) error {
	toInstall := Applications().GetAll(apps...)
	for _, i := range toInstall {
		err := i.Install()
		if err != nil && err != ErrAlreadyInstalled {
			return err
		}
	}
	return nil
}

func Uninstall(apps ...string) error {
	toRemove := Applications().GetAll(apps...)
	for _, i := range toRemove {
		err := i.Uninstall()
		if err != nil && err != ErrIsNotInstalled {
			return err
		}
	}
	return nil
}

func Update() error {
	for _, r := range Repositories() {
		err := r.Update()
		if err != nil {
			return err
		}
	}

	return nil
}

func Upgrade(apps ...string) error {
	toUpgrade := Applications().GetAll(apps...)
	if len(apps) == 1 && apps[0] == "*" {
		toUpgrade = Applications().Installed()
	}
	for _, i := range toUpgrade {
		err := i.Upgrade()
		if err != nil && err != ErrIsNotInstalled {
			return err
		}
	}
	return nil
}
