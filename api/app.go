package api

import (
	"encoding/json"
	"github.com/goo-app/cli/log"
	"github.com/goo-app/cli/utils"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	manifestCache = map[string]AppManifest{}
)

type (
	Apps []App
	App  struct {
		Name string
		repo Repo
	}
)

func Applications() Apps {
	return Repositories().Applications()
}

func (r Repos) Applications() Apps {
	return utils.FlatMapSlice(Repositories(), func(r Repo) []App {
		return r.Applications()
	})
}

func (apps Apps) Installed() []App {
	return utils.FilterSlice(apps, func(a App) bool {
		return a.IsInstalled()
	})
}

func (apps Apps) Filter(app string, regex bool) (list []App) {
	return utils.FilterSlice(apps, func(a App) bool {
		if regex {
			x, err := regexp.Compile(strings.ToLower(app))
			if err != nil {
				log.Error(err.Error())
				return false
			}
			return x.MatchString(strings.ToLower(a.Name))
		}
		return strings.Contains(strings.ToLower(a.Name), strings.ToLower(app))
	})
}
func (apps Apps) GetAll(app ...string) []App {
	return utils.FilterSlice(apps, func(a App) bool {
		return utils.AnySlice(app, func(an string) bool {
			return a.Name == an
		})
	})
}

func (a App) LoadManifest() error {
	if mf, ok := manifestCache[a.Name]; !ok {
		defer func() { manifestCache[a.Name] = mf }()
		mfFile := path.Join(a.repo.appsRepoPath(), a.Name+".json")
		mf = AppManifest{}
		data, err := os.ReadFile(mfFile)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &mf)
		if err != nil {
			return err
		}
		return mf.Validate()
	}
	return nil
}

func (a App) Manifest() (*AppManifest, error) {
	if mf, ok := manifestCache[a.Name]; ok {
		return &mf, nil
	} else {
		err := a.LoadManifest()
		if err != nil {
			return nil, err
		}
		return a.Manifest()
	}
}

func (a App) IsDeprecated() (bool, error) {
	mf, err := a.Manifest()
	if err != nil {
		return false, err
	}
	return mf.Deprecated != nil && mf.Deprecated.IsDeprecated(), nil
}

func (a App) GetDeprecation() *string {
	if d, _ := a.IsDeprecated(); d {
		return nil
	}
	mf, _ := a.Manifest()
	return mf.Deprecated.get(a.Name)
}

func (a App) IsInstalled() bool {
	_, err := os.Stat(a.appPath())
	if os.IsExist(err) {
		_, err = os.Stat(a.currentAppPath())
		return os.IsExist(err)
	}
	return false
}

func (a App) CurrentVersion() (string, error) {
	link, err := a.currentLinkedAppPath()
	return path.Base(link), err
}

func (a App) currentLinkedAppPath() (string, error) {
	return filepath.EvalSymlinks(a.currentAppPath())
}

func (a App) currentAppPath() string {
	return path.Join(a.appPath(), "current")
}

func (a App) appPath() string {
	return path.Join(appsPath(), a.Name)
}

func appsPath() string {
	return path.Join(Path, "apps")
}
