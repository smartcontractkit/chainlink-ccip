package utils

import (
	"regexp"
)

// IsCustomImage checks if DEVSPACE_IMAGE is overridden to use custom image
func IsCustomImage(devspaceNamespace string, devspaceImage string) bool {
	cribLocalPattern := `^localhost:5001/chainlink-[a-z]*-devspace(:[a-zA-Z0-9._-]+)?$`
	awsECRPattern := `^323150190480\.dkr\.ecr\.us-west-2\.amazonaws\.com/chainlink-[a-z]*-devspace(:[a-zA-Z0-9._-]+)?$`

	cribLocalRegex := regexp.MustCompile(cribLocalPattern)
	awsECRRegex := regexp.MustCompile(awsECRPattern)
	if devspaceNamespace == "crib-local" && cribLocalRegex.MatchString(devspaceImage) {
		return true
	} else if awsECRRegex.MatchString(devspaceImage) {
		return true
	}

	return false
}
