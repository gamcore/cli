package utils

func MkLink(source, target string) error {
	return mklink(source, target)
}
