package api

import "net/url"

func (m AppManifest) Validate() error {
	if m.Description == "" {
		return ErrManifestDescriptionEmpty
	}
	if len(m.Executable) == 0 {
		return ErrManifestRequiredAnyExec
	}
	if m.Homepage == "" {
		return ErrManifestUrlEmpty
	}
	if _, err := url.Parse(m.Homepage); err != nil {
		return err
	}

	if err := m.License.Validate(); err != nil {
		return err
	}

	if err := m.Updates.Validate(); err != nil {
		return err
	}

	return nil
}

func (l License) Validate() error {
	if l.Name == "" {
		return ErrManifestLicenseNameEmpty
	}
	if l.URL == "" {
		return ErrManifestUrlEmpty
	}
	if _, err := url.Parse(l.URL); err != nil {
		return err
	}

	return nil
}

func (u AppUpdateSchema) Validate() error {
	switch u.Type {
	case AppUpdateGitHub:
		return u.validateGithub()
	case AppUpdateHTML:
		return u.validateHtml()
	case AppUpdateXML, AppUpdateJSON:
		return u.validateXmlJson()
	default:
		return ErrSchemaInvalidType
	}
}

func (u AppUpdateSchema) validateGithub() error {
	if u.URL == "" {
		return ErrManifestUrlEmpty
	}
	if _, err := url.Parse(u.URL); err != nil {
		return err
	}
	if !githubUrlRegex.MatchString(u.URL) {
		return ErrSchemaInvalidURL
	}
	return nil
}
func (u AppUpdateSchema) validateHtml() error {
	if u.URL == "" {
		return ErrManifestUrlEmpty
	}
	if _, err := url.Parse(u.URL); err != nil {
		return err
	}
	if u.GetPattern() == "" {
		return ErrManifestUrlPatternRequired
	}
	if _, err := url.Parse(u.GetPattern()); err != nil {
		return err
	}
	return nil
}

func (u AppUpdateSchema) validateXmlJson() error {
	if u.URL == "" {
		return ErrManifestUrlEmpty
	}
	if _, err := url.Parse(u.URL); err != nil {
		return err
	}
	if u.Path == nil {
		return ErrSchemaPathRequired
	}
	if u.GetPattern() == "" {
		return ErrManifestUrlPatternRequired
	}
	if _, err := url.Parse(u.GetPattern()); err != nil {
		return err
	}

	return nil
}
