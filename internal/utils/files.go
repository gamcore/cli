package utils

import (
	"fmt"
	"io"
	"os"
	p "path"
	"strings"
)

func GetFirstFileExtFilter(path, basename string, ext ...string) string {
	nameWithExt := MapSlice(ext, func(e string) string {
		return fmt.Sprintf("%s.%s", basename, e)
	})
	entries, _ := os.ReadDir(path)
	for _, entry := range entries {
		if !entry.IsDir() && AnySlice(nameWithExt, func(e string) bool {
			return strings.ToLower(e) == strings.ToLower(entry.Name())
		}) {
			return p.Join(path, entry.Name())
		}
	}

	return ""
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
