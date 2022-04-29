package internal

import (
	"github.com/stachu540/goo/internal/utils"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type Applications []Application

func ApplicationSlice(r Repository) Applications {
	appsRepoDir := path.Join(r.Path(), "apps")
	manifests, _ := ioutil.ReadDir(appsRepoDir)

	return utils.MapSlice(manifests, func(f os.FileInfo) Application {
		name := f.Name()
		name = name[:strings.LastIndex(name, ".")]

		var app = Application{
			Name:       name,
			repository: r,
		}
		return app
	})
}

func (a Applications) Find(query string, regex bool) (apps Applications) {
	return utils.FilterSlice(a, func(a Application) bool {
		if regex {
			q := regexp.MustCompile(query)
			return q.MatchString(a.Name)
		} else {
			return strings.Contains(a.Name, query)
		}
	})
}

func (a Applications) Fetch(apps ...string) Applications {
	return utils.FilterSlice(a, func(app Application) bool {
		for _, a := range apps {
			if strings.ToLower(app.Name) == strings.ToLower(a) {
				return true
			}
		}
		return false
	})
}
