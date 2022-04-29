package internal

import (
	"github.com/stachu540/goo/internal/utils"
	"strings"
)

func Update() error {
	for _, repo := range GetRepositories() {
		err := repo.Update()
		if err != nil {
			return err
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
