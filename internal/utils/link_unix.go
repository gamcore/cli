//go:build !windows

package utils

import "os"

func mklink(source, target string) error {
	return os.Symlink(source, target)
}
