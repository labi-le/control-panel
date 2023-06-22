package internal

import (
	"fmt"
)

var (
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

func BuildVersion() string {
	return fmt.Sprintf("Version: %s\nCommit hash: %s\nBuild time: %s", Version, CommitHash, BuildTime)
}
