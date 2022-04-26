//go:build windows

package utils

import (
	win "github.com/nyaosorg/go-windows-junction"
	"os"
)

func mklink(source, target string) error {
	if src, _ := os.Stat(source); src.IsDir() {
		return junction(source, target)
	} else {
		return link(source, target)
	}
}

func junction(source, target string) error {
	return win.Create(source, target)
}

func link(source, target string) error {
	return os.Link(source, target)
}
