package internal

func (a Application) HasUpdate() (bool, error) {
	var update = false

	return update, nil
}

func (a Application) Update() error {
	if a.IsInstalled() {
		return a.update()
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) update() error {
	return nil
}
