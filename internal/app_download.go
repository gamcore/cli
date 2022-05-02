package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

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

func (a Application) download() error {
	return nil
}
