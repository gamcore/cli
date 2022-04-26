package internal

import (
	"github.com/hashicorp/go-version"
	"strconv"
	"time"
)

var (
	Version         string
	Timestamp       string
	Revision        string
	ProjectRepoSlug string
	CoreRepoSlug    string
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
