package api

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/goo-app/cli/utils"
)

func (r Repo) Applications() Apps {
	mfs, _ := ioutil.ReadDir(r.appsRepoPath())

	return utils.MapSlice(mfs, func(f os.FileInfo) App {
		name := f.Name()
		name = name[:strings.LastIndex(name, ".")]

		return App{Name: name, repo: r}
	})
}

func (r Repo) appsRepoPath() string {
	return path.Join(r.repoPath(), "apps")
}
