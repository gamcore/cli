package api

import (
	"strconv"
	"time"

	"github.com/hashicorp/go-version"
)

var (
	Version         string
	Timestamp       string
	Revision        string
	ProjectRepoSlug string
)

func GetVersion() version.Version {
	return *version.Must(version.NewSemver(Version))
}

func GetTimestamp() time.Time {
	timestamp, err := strconv.ParseInt(Timestamp, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.UnixMilli(timestamp)
}
