package internal

import "github.com/goo-app/cli/internal/logger"

func Install(apps ...string) error {
	toInstall := GetApplications().Fetch(apps...)
	for _, a := range toInstall {
		err := a.Install()
		if err != nil {
			if err == ErrAlreadyInstalled {
				if len(apps) == 1 {
					return err
				}
				logger.WarnF(`"%s" is already installed`, a.Name)
			} else {
				return err
			}
		}
	}
	return nil
}
