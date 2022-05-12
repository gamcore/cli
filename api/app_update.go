package api

import (
	"github.com/hashicorp/go-version"
	"os"
)

type (
	AppUpdateType int
)

var (
	appUpdatesRaw = map[AppUpdateType]string{
		AppUpdateHTML:   "html",
		AppUpdateXML:    "xml",
		AppUpdateJSON:   "json",
		AppUpdateGitHub: "github",
	}
	appUpdates = map[string]AppUpdateType{
		"html":   AppUpdateHTML,
		"xml":    AppUpdateXML,
		"json":   AppUpdateJSON,
		"github": AppUpdateGitHub,
	}
)

const (
	AppUpdateHTML AppUpdateType = iota
	AppUpdateXML
	AppUpdateJSON
	AppUpdateGitHub
)

func (t AppUpdateType) String() string {
	return appUpdatesRaw[t]
}

func (t AppUpdateType) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *AppUpdateType) UnmarshalJSON(data []byte) error {
	*t = appUpdates[string(data)]

	return nil
}

func (a App) HasUpdate() (bool, error) {
	a.IsInstalled()
	mf, err := a.Manifest()
	if err != nil {
		return false, err
	}
	v, _, err := mf.CheckUpdate()
	if err != nil {
		return false, err
	}
	linkedApp, err := a.currentLinkedAppPath()
	if err != nil {
		return false, err
	}
	linkedPath, err := os.Stat(linkedApp)
	if err != nil && os.IsNotExist(err) {
		return false, err
	}
	currentV, err := version.NewVersion(linkedPath.Name())
	if err != nil {
		return false, err
	}
	latestV, err := version.NewVersion(v)
	if err != nil {
		return false, err
	}
	return currentV.GreaterThan(latestV), nil
}
