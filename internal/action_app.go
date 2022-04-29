package internal

import (
	"fmt"
	"github.com/stachu540/goo/internal/utils"
	"strings"
)

func GetApplications() Applications {
	apps := Applications{}
	for _, r := range GetRepositories() {
		apps = append(apps, r.Applications()...)
	}
	return apps
}
func GetInstalledApplications() Applications {
	return utils.FilterSlice(GetApplications(), func(app Application) bool { return app.IsInstalled() })
}

func GetApplication(name string) (*Application, error) {
	for _, a := range GetApplications() {
		if strings.ToLower(a.Name) == strings.ToLower(name) {
			return &a, nil
		}
	}

	return nil, fmt.Errorf(`could not find a app name "%s"`, name)
}
