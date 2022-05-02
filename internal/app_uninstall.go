package internal

func (a Application) Uninstall() error {
	if a.IsInstalled() {
		return a.uninstall()
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) uninstall() error {
	return nil
}
