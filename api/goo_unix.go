//go:build !windows

package api

import "runtime"

func shellScript() string {
	data := `#!/bin/sh
if [ -d "$GOO_PATH/apps/goo/new" ]; then
	if pgrep goo; then
	  killall -9 goo
	fi
	mv "$GOO_PATH/apps/goo/current" "$GOO_PATH/apps/goo/old"
	mv "$GOO_PATH/apps/goo/new" "$GOO_PATH/apps/goo/current"
fi

$GOO_PATH/apps/goo/current/goo $@
`

	return data
}

func (u AppUpdateSchema) GetPattern() string {
	if runtime.GOOS == "darwin" {
		return u.Pattern.Macos
	} else {
		return u.Pattern.Linux
	}
}
