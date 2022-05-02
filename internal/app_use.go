package internal

import (
	"fmt"
	"github.com/goo-app/cli/internal/utils"
	"os"
)

func (a Application) Use(version string) error {
	if a.IsInstalled() {
		chdir, err := os.Getwd()
		if err != nil {
			return err
		}
		err = os.Chdir(a.path())
		if err != nil {
			return err
		}
		err = utils.MkLink(fmt.Sprintf("./%s", version), "current")
		if err != nil {
			return err
		}
		return os.Chdir(chdir)
	} else {
		return ErrIsNotInstalled
	}
}
