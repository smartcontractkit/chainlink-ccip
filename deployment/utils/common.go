package utils

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

func NewRegistererID(chainFamily string, version *semver.Version) string {
	return fmt.Sprintf("%s-%s", chainFamily, version.String())
}
