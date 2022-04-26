package internal

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/stachu540/goo/internal/utils"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"strings"
)

type Repositories []Repository

type Repository struct {
	Name     string
	Manifest RManifest
	path     string
}

func Read(gpath string, basename string) (*Repository, error) {
	rpath := path.Join(gpath, "repos", basename)
	manifestFile := utils.GetFirstFileExtFilter(rpath, "manifest", "json", "yml", "yaml")
	if manifestFile == "" {
		return nil, fmt.Errorf("no manifest found for", basename)
	}

	manifest := RManifest{}
	ext := path.Ext(manifestFile)
	data, err := os.ReadFile(manifestFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read manifest file for", basename, err)
	}

	switch strings.ToLower(ext) {
	case "json":
		err = json.Unmarshal(data, &manifest)
		if err != nil {
			return nil, fmt.Errorf("cannot parse manifest file variables for", basename, err)
		}
		break
	case "yml", "yaml":
		err = yaml.Unmarshal(data, &manifest)
		if err != nil {
			return nil, fmt.Errorf("cannot parse manifest file variables for", basename, err)
		}
		break
	default:
		return nil, fmt.Errorf("unsupported manifest file format for %s \"%s\"", basename, ext)
	}

	return &Repository{
		Name: basename, Manifest: manifest, path: rpath,
	}, nil
}

func (r Repository) Path() string {
	return r.path
}

func (r Repository) Applications() Applications {
	return ApplicationSlice(r)
}

func (r Repository) Update() error {
	gitRepo, err := git.PlainOpen(r.Path())
	if err != nil {
		return fmt.Errorf("could not uppdate repository %s: %s", r.Name, err)
	}

	wTree, err := gitRepo.Worktree()
	if err != nil {
		return err
	}

	return wTree.Pull(&git.PullOptions{RemoteName: "origin"})
}
