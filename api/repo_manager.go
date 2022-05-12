package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/goo-app/cli/log"
	"github.com/goo-app/cli/utils"
	"github.com/manifoldco/promptui"
	"net/url"
	"os"
	"path"
	"strings"
)

type RepoManager struct {
	Repo
}

func (r RepoManager) Validate() error {
	for _, app := range r.Applications() {
		mf, err := app.Manifest()
		if err != nil {
			return err
		}
		err = mf.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r RepoManager) AddApplication() error {
	appNames := utils.MapSlice(r.Applications(), func(a App) string {
		return a.Name
	})
	name := prompt("Application Name", func(input string) error {
		if err := required(input); err != nil {
			return err
		}
		if utils.AnySlice(appNames, func(a string) bool {
			return strings.ToLower(a) == strings.ToLower(input)
		}) {
			return ErrInputAlreadyExist
		}
		return nil
	})
	description := prompt("Application description", required)
	homepage := prompt("Application description", func(input string) error {
		err := required(input)
		if err != nil {
			return err
		}
		if _, err := url.Parse(input); err != nil {
			return err
		}
		return nil
	})
	license := prompt("Application license [Default: MIT]", skip)

	confirm("Do you wanna to add it? [Default: Y]", func(input string) error {
		if strings.TrimSpace(input) != "" || strings.ToLower(input) != "y" {
			log.Warn("Aborted!")
			os.Exit(0)
		}
		name = strings.ToLower(name)
		return nil
	})

	mf := AppManifest{
		Description: description,
		Homepage:    homepage,
		License: License{
			Name: license,
		},
		Executable: []string{},
		Updates:    AppUpdateSchema{},
	}

	data, err := json.Marshal(mf)
	if err != nil {
		return err
	}
	jName := fmt.Sprintf("%s.json", strings.ToLower(name))
	jPath := path.Join(r.appsRepoPath(), jName)
	err = os.WriteFile(jPath, data, os.ModeType)
	if err != nil {
		return err
	}

	log.WarnF(`"%s" has been added. Please update manifest first before start verification`, strings.ToLower(name))
	return nil
}

func (r RepoManager) RemoveApplications(names ...string) error {
	for _, a := range r.Applications().GetAll(names...) {
		err := os.Remove(path.Join(r.repoPath(), "apps", a.Name+".json"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (r RepoManager) Init() error {
	repo, err := git.PlainInit(r.repoPath(), false)
	if err != nil {
		return err
	}
	// TODO: Create .gitignore
	err = os.Mkdir(r.appsRepoPath(), os.ModeDir)
	if err != nil {
		return err
	}
	// TODO: CI CD Support
	// TODO: Generate Readme
	tree, err := repo.Worktree()
	if err != nil {
		return err
	}
	_, err = tree.Add(".")
	if err != nil {
		return err
	}
	return nil
}

func prompt(label string, validation func(string) error) string {
	p := promptui.Prompt{
		Label:    label,
		Validate: validation,
	}
	input, err := p.Run()

	for err != nil {
		if err == promptui.ErrAbort || err == promptui.ErrInterrupt {
			os.Exit(0)
		}
		log.Warn(err.Error())
		return prompt(label, validation)
	}
	return input
}

func required(input string) error {
	if strings.TrimSpace(input) == "" {
		return ErrInputEmpty
	}
	return nil
}
func skip(_ string) error {
	return nil
}

func confirm(label string, validation func(string) error) bool {
	p := promptui.Prompt{
		Label:    label,
		Validate: validation,
	}
	c, _ := p.Run()
	return strings.TrimSpace(c) == "" || strings.ToLower(c) == "y"
}
