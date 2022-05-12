package api

import "os"

func (r Repo) Delete() error {
	return os.RemoveAll(r.Path())
}
