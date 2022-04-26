package internal

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"strings"
)

type (
	Deprecated struct {
		Reason      string  `json:"reason" yaml:"reason"`
		Since       string  `json:"since" yaml:"since"`
		Alternative *string `json:"alternative,omitempty" yaml:"alternative,omitempty"`
	}
	RManifest struct {
		SupportSince string      `json:"support_since" json:"support_since"`
		Deprecated   *Deprecated `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	}
	AManifest struct {
		Description string             `json:"description"`
		Homepage    string             `json:"homepage"`
		License     License            `json:"license"`
		Actions     *ActionScripts     `json:"actions,omitempty"`
		Updates     UpdateSchema       `json:"updates"`
		Deprecated  *Deprecated        `json:"deprecated,omitempty"`
		Appendix    *map[string]string `json:"appendix,omitempty"`
	}
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	UpdateSchema struct {
		Type    UpdateType `json:"type"`
		URL     string     `json:"url"`
		Pattern string     `json:"pattern"`
	}
	ActionScripts struct {
		PreInstall    []string `json:"pre_install,omitempty"`
		PostInstall   []string `json:"post_install,omitempty"`
		PreUpdate     []string `json:"pre_update,omitempty"`
		PostUpdate    []string `json:"post_update,omitempty"`
		PreUninstall  []string `json:"pre_uninstall,omitempty"`
		PostUninstall []string `json:"post_uninstall,omitempty"`
	}
	UpdateType int
)

func (d *Deprecated) Get(name string) *string {
	v := GetVersion()
	sb := new(strings.Builder)
	if d != nil && v.GreaterThanOrEqual(version.Must(version.NewSemver(d.Since))) {
		sb.WriteString(name)
		sb.WriteString(" has been deprecated: \"")
		sb.WriteString(d.Reason)
		sb.WriteString("\"")
		if d.Alternative != nil {
			sb.WriteString(" Please use " + *d.Alternative)
		}
	}
	result := strings.TrimSpace(sb.String())
	if result != "" {
		return &result
	} else {
		return nil
	}
}

const (
	UpdateHTML UpdateType = iota
	UpdateXML
	UpdateJSON
	UpdateGitHub
)

func (t UpdateType) String() string {
	switch t {
	case UpdateHTML:
		return "html"
	case UpdateXML:
		return "xml"
	case UpdateJSON:
		return "json"
	case UpdateGitHub:
		return "github"
	}
	return ""
}

func (t UpdateType) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *UpdateType) UnmarshalJSON(data []byte) error {
	var ut UpdateType
	switch string(data) {
	case "html":
		ut = UpdateHTML
		break
	case "xml":
		ut = UpdateXML
		break
	case "json":
		ut = UpdateJSON
		break
	case "github":
		ut = UpdateGitHub
	default:
		return fmt.Errorf(`unknown update type "%s", supports: "html", "xml", "json" and "github"`, string(data))
	}

	t = &ut

	return nil
}
