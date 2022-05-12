package api

import (
	"context"
	jsonq "github.com/antchfx/jsonquery"
	xmlq "github.com/antchfx/xmlquery"
	"github.com/goo-app/cli/utils"
	"github.com/google/go-github/v44/github"
	"regexp"
	"runtime"
	"strings"
	"text/template"
)

type (
	updatePatternSpec struct {
		Os      string
		Arch    string
		Version string
	}
)

var (
	versionRegex   = regexp.MustCompile("^v?(?P<version>(\\d+\\.)+\\d+)$")
	githubUrlRegex = regexp.MustCompile("^https://github.com/(?P<owner>.+)/(?P<repo_name>.+)$")
)

func (m AppManifest) CheckUpdate() (string, string, error) {
	switch m.Updates.Type {
	case AppUpdateHTML:
		return m.checkUpdateHtml()
	case AppUpdateXML:
		return m.checkUpdateXml()
	case AppUpdateJSON:
		return m.checkUpdateJson()
	case AppUpdateGitHub:
		return m.checkUpdateGithub()
	}
	return "", "", ErrSchemaInvalidType
}

func (m AppManifest) checkUpdateHtml() (string, string, error) {
	return "", "", ErrSchemaIsNotSupportedYet
}

func (m AppManifest) checkUpdateXml() (string, string, error) {
	res, err := httpClient.Get(m.Updates.URL)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	doc, err := xmlq.Parse(res.Body)
	if err != nil {
		return "", "", err
	}

	n := xmlq.FindOne(doc, *m.Updates.Path)

	if n != nil && n.Type == xmlq.TextNode {
		rv := n.InnerText()
		v := m.extractVersion(rv)
		p, err := m.formatUrl(m.Updates.GetPattern(), v)
		if err != nil {
			return "", "", err
		}
		return rv, p, nil
	}

	return "", "", ErrSchemaInvalidPath
}

func (m AppManifest) checkUpdateJson() (string, string, error) {
	res, err := httpClient.Get(m.Updates.URL)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	doc, err := jsonq.Parse(res.Body)
	if err != nil {
		return "", "", err
	}

	n := jsonq.FindOne(doc, *m.Updates.Path)

	if n != nil && n.Type == jsonq.TextNode {
		rv := n.InnerText()
		v := m.extractVersion(rv)
		p, err := m.formatUrl(m.Updates.GetPattern(), v)
		if err != nil {
			return "", "", err
		}
		return rv, p, nil
	}

	return "", "", ErrSchemaInvalidPath
}

func (m AppManifest) checkUpdateGithub() (string, string, error) {
	tagPattern := versionRegex
	if m.Updates.VersionRegex != nil {
		tagPattern = regexp.MustCompile(*m.Updates.VersionRegex)
	}
	if githubUrlRegex.MatchString(m.Updates.URL) {
		matches := githubUrlRegex.FindStringSubmatch(m.Updates.URL)
		var owner, repo string
		for i, name := range githubUrlRegex.SubexpNames() {
			if i != 0 && name != "" {
				switch name {
				case "owner":
					owner = matches[i]
				case "repo_name":
					repo = matches[i]
				}
			}
		}
		data, res, err := githubClient.Repositories.GetLatestRelease(context.Background(), owner, repo)
		if err != nil {
			return "", "", err
		}
		defer res.Body.Close()

		if tagPattern.MatchString(data.GetTagName()) {
			asset := m.Updates.getAsset(data.Assets, m.Appendix)
			return data.GetTagName(), asset, nil
		} else {
			return "", "", ErrSchemaInvalidPattern
		}
	} else {
		return "", "", ErrSchemaInvalidURL
	}
}

func (m AppManifest) formatUrl(raw, version string) (string, error) {
	return m.formatUrlOsArch(raw, version, runtime.GOOS, runtime.GOARCH)
}

func (m AppManifest) formatUrlOsArch(raw, version, os, arch string) (string, error) {
	if val, ok := m.Appendix[os]; ok {
		os = val
	}
	if val, ok := m.Appendix[arch]; ok {
		arch = val
	}
	tmpl, err := template.New("url").Parse(raw)
	if err != nil {
		return "", err
	}
	sb := strings.Builder{}
	err = tmpl.Execute(&sb, updatePatternSpec{os, arch, version})
	if err != nil {
		return "", err
	}
	return sb.String(), nil
}

func (m AppManifest) extractVersion(raw string) string {
	vRegex := versionRegex
	if m.Updates.VersionRegex != nil {
		vRegex = regexp.MustCompile(*m.Updates.VersionRegex)
	}
	if vRegex.MatchString(raw) {
		return vRegex.ReplaceAllString(raw, "$1")
	}
	return raw
}

func (u AppUpdateSchema) getAsset(assets []*github.ReleaseAsset, appendix map[string]string) string {
	archAlt := map[string]string{
		"amd64": "386",
		"arm64": "arm",
	}
	osAlt := map[string]string{
		"windows": "win",
		"macos":   "mac",
		"linux":   "unix",
	}
	os := runtime.GOOS
	arch := runtime.GOARCH

	sliceByOs := utils.FilterSlice(assets, func(a *github.ReleaseAsset) bool {
		if strings.Contains(a.GetName(), os) {
			return true
		}
		if x, ok := appendix[os]; ok && strings.Contains(a.GetName(), x) {
			return true
		}
		inOs := osAlt[os]
		if strings.Contains(a.GetName(), inOs) {
			return true
		}
		if x, ok := appendix[inOs]; ok && strings.Contains(a.GetName(), x) {
			return true
		}
		return false
	})

	if len(sliceByOs) == 1 {
		return sliceByOs[0].GetBrowserDownloadURL()
	} else {
		asset := *utils.FindFirstSlice(sliceByOs, func(a *github.ReleaseAsset) bool {
			if strings.Contains(a.GetName(), arch) {
				return true
			}
			if x, ok := appendix[arch]; ok && strings.Contains(a.GetName(), x) {
				return true
			}
			inArch := archAlt[arch]
			if strings.Contains(a.GetName(), inArch) {
				return true
			}
			if x, ok := appendix[inArch]; ok && strings.Contains(a.GetName(), x) {
				return true
			}
			return false
		})
		return asset.GetBrowserDownloadURL()
	}
}
