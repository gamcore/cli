package internal

func (a Application) Install() error {
	if a.IsInstalled() {
		return ErrAlreadyInstalled
	} else {
		return a.install()
	}
}

func (a Application) install() error {
	return nil
}
