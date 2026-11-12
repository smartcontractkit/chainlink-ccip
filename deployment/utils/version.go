package utils

import "github.com/Masterminds/semver/v3"

// SameMajor reports whether two versions are on the same major line. A program upgraded in place
// (for example 2.0.0 to 2.0.1) stays on its major line, so lookups that only care about the line
// should compare majors rather than whole versions.
func SameMajor(a, b *semver.Version) bool {
	return a != nil && b != nil && a.Major() == b.Major()
}

// ResolveRegisteredVersion picks the registered version that serves a wanted one. An exact
// registration always wins. Otherwise, for versions from 2.0.0 up, the highest registered version
// on the same major line serves it, so a 2.0.1 ref resolves to the adapter registered for 2.0.0.
// Versions below 2.0.0 never fall back: 1.5.0 and 1.6.0 are different lines with different adapters.
func ResolveRegisteredVersion(registered []*semver.Version, want *semver.Version) (*semver.Version, bool) {
	if want == nil {
		return nil, false
	}
	var best *semver.Version
	for _, v := range registered {
		if v == nil {
			continue
		}
		if v.Equal(want) {
			return v, true
		}
		if want.Major() >= Version_2_0_0.Major() && SameMajor(v, want) && (best == nil || v.GreaterThan(best)) {
			best = v
		}
	}
	return best, best != nil
}
