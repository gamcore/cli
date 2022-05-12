package api

import (
	"github.com/goo-app/cli/utils"
	"path"
)

func (a App) Upgrade() error {
	if u, err := a.HasUpdate(); err == nil {
		if u {
			return a.doUpgrade()
		}
	} else {
		return err
	}
	return nil
}

func (a App) doUpgrade() error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	v, link, err := mf.CheckUpdate()
	if err != nil {
		return err
	}
	actualPath := path.Join(a.appPath(), v)
	return a.actionUpdate(v, func() error {
		err = a.doDownloadAndExtract(link, actualPath)
		if err != nil {
			return err
		}
		return utils.MkLink(actualPath, a.currentAppPath())
	})
}
