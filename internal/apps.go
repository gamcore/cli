package internal

import (
	"encoding/json"
	"fmt"
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

type Application struct {
	Name       string
	repository Repository
}

func (a Application) Manifest() (manifest AManifest) {
	mFile := path.Join(a.repository.Path(), "apps", a.Name+".json")
	mFileData, _ := os.ReadFile(mFile)
	err := json.Unmarshal(mFileData, &manifest)
	if err != nil {
		Logger.Debugf("error read manifest file: %s", err)
	}
	return
}

func (a Application) Repository() Repository {
	return a.repository
}

func (a Application) CurrentVersion() (*string, error) {
	if a.IsInstalled() {
		ln, err := os.Readlink(path.Join(a.path(), "current"))
		if err != nil {
			return nil, fmt.Errorf("cannot read a current version: %s", err)
		}
		if f, err := os.Stat(ln); os.IsExist(err) {
			name := f.Name()
			return &name, nil
		} else if os.IsNotExist(err) {
			return nil, fmt.Errorf("cannot read a current version: %s", err)
		}
	}
	return nil, ErrIsNotInstalled
}

func (a Application) Versions() []string {
	var versions []string
	if a.IsInstalled() {
		files, _ := ioutil.ReadDir(a.path())
		for _, file := range files {
			if file.IsDir() && file.Name() != "current" {
				versions = append(versions, file.Name())
			}
		}
	}
	return versions
}

func (a Application) IsInstalled() bool {
	_, err1 := os.Stat(a.path())
	_, err2 := os.Stat(path.Join(a.path(), "current"))

	return os.IsExist(err1) && os.IsExist(err2)
}

func (a Application) Path() *string {
	var p = a.path()
	if a.IsInstalled() {
		return &p
	} else {
		return nil
	}
}

func (a Application) Install() error {
	if a.IsInstalled() {
		return ErrAlreadyInstalled
	} else {
		return a.install()
	}
}

func (a Application) Update() error {
	if a.IsInstalled() {
		return a.update()
	} else {
		return ErrIsNotInstalled
	}

}

func (a Application) Use(version string) error {
	if a.IsInstalled() {
		chdir, err := os.Getwd()
		if err != nil {
			return err
		}
		err = os.Chdir(a.path())
		if err != nil {
			return err
		}
		err = os.Symlink(fmt.Sprintf("./%s", version), "current")
		if err != nil {
			return err
		}
		return os.Chdir(chdir)
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) Uninstall() error {
	if a.IsInstalled() {
		return a.uninstall()
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) Clean() error {
	if a.IsInstalled() {
		return a.clean()
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) install() error {
	//a.download()
	return nil
}

func (a Application) update() error {
	return nil
}

func (a Application) uninstall() error {
	return nil
}

func (a Application) clean() error {
	return nil
}

func (a Application) path() string {
	return path.Join(Path, "apps", a.Name)
}
