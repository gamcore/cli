package internal

import (
	"github.com/stachu540/goo/internal/utils"
	"strings"
)

func Uninstall(apps ...string) error {
	toUninstall := utils.FilterSlice(GetApplications(), func(a Application) bool {
		if a.IsInstalled() {
			for _, app := range apps {
				if strings.ToLower(app) == strings.ToLower(a.Name) {
					return true
				}
			}
		}
		return false
	})

	for _, a := range toUninstall {
		err := a.Uninstall()
		if err != nil {
			return err
		}
	}

	return nil
}
