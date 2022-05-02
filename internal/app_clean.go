package internal

import "github.com/goo-app/cli/internal/logger"

func (a Application) Clean() error {
	if a.IsInstalled() {
		return a.clean()
	} else {
		return ErrIsNotInstalled
	}
}

func (a Application) clean() error {
	logger.InfoF(`Removing old version form "%s"`, a.Name)
	return nil
}
