package internal

import (
	"github.com/goo-app/cli/internal/utils"
	"strings"
)

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
