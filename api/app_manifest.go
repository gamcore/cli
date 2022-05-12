package api

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

type (
	AppManifest struct {
		Description string            `json:"description"`          // Application Description
		Homepage    string            `json:"homepage"`             // Application Homepage
		License     License           `json:"license"`              // Application License
		Actions     *ActionScripts    `json:"actions,omitempty"`    // Action Scripts for each installation interaction
		Executable  []string          `json:"executable,omitempty"` // Executable relative path (without extension)
		Updates     AppUpdateSchema   `json:"updates"`              // Application Updates Schema
		Deprecated  *Deprecated       `json:"deprecated,omitempty"` // Deprecation for Application
		Appendix    map[string]string `json:"appendix,omitempty"`   // Appendix for OS / Architecture
	}
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	AppUpdateSchema struct {
		OneFile      bool             `json:"one_file"`
		Type         AppUpdateType    `json:"type"`
		URL          string           `json:"url"`
		Pattern      AppUpdatePattern `json:"pattern"`
		Path         *string          `json:"path,omitempty"`
		VersionRegex *string          `json:"version_regex,omitempty"`
	}
	AppUpdatePattern struct {
		Windows string `json:"windows,omitempty"`
		Linux   string `json:"linux,omitempty"`
		Macos   string `json:"macos,omitempty"`
	}
	ActionScripts struct {
		PreInstall    []string `json:"pre_install,omitempty"`
		PostInstall   []string `json:"post_install,omitempty"`
		PreUpdate     []string `json:"pre_update,omitempty"`
		PostUpdate    []string `json:"post_update,omitempty"`
		PreUninstall  []string `json:"pre_uninstall,omitempty"`
		PostUninstall []string `json:"post_uninstall,omitempty"`
	}
	actionFormatter struct {
		Name    string
		Version *string
		Paths   actionFormatterPaths
	}
	actionFormatterPaths struct {
		App   string
		Goo   string
		Temp  string
		Shims string
	}
)

func (a App) actionInstall(version string, actionConsumer func() error) error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	formatter := doActionFormatter(a, &version)
	err = doAction(mf.Actions.PreInstall, formatter)
	if err != nil {
		return err
	}
	err = actionConsumer()
	if err != nil {
		return err
	}
	return doAction(mf.Actions.PostInstall, formatter)
}

func (a App) actionUpdate(version string, actionConsumer func() error) error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	formatter := doActionFormatter(a, &version)
	err = doAction(mf.Actions.PreUpdate, formatter)
	if err != nil {
		return err
	}
	err = actionConsumer()
	if err != nil {
		return err
	}
	return doAction(mf.Actions.PostUpdate, formatter)
}

func (a App) actionUninstall(actionConsumer func() error) error {
	mf, err := a.Manifest()
	if err != nil {
		return err
	}
	formatter := doActionFormatter(a, nil)
	err = doAction(mf.Actions.PreUninstall, formatter)
	if err != nil {
		return err
	}
	err = actionConsumer()
	if err != nil {
		return err
	}
	return doAction(mf.Actions.PostUninstall, formatter)
}

func doAction(actions []string, formatter actionFormatter) error {
	for _, action := range actions {
		tmpl, err := template.New("ActionSpec").Parse(action)
		if err != nil {
			return err
		}
		out := strings.Builder{}
		err = tmpl.Execute(&out, formatter)
		if err != nil {
			return err
		}
		cmdArgv := strings.Split(strings.TrimSpace(out.String()), " ")
		x := exec.Command(cmdArgv[0], cmdArgv[1:]...)
		x.Stdout = os.Stdout
		x.Stderr = os.Stderr
		err = x.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func doActionFormatter(a App, version *string) actionFormatter {
	return actionFormatter{
		Name:    a.Name,
		Version: version,
		Paths: actionFormatterPaths{
			Goo:   Path,
			Temp:  path.Join(Path, "tmp"),
			App:   path.Join(Path, "apps", a.Name),
			Shims: path.Join(Path, "shims", a.Name),
		},
	}
}
