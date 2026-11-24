package hooks

import "github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"

// extraArgsSource falls back to a family-only source because adapters registered as ForkCCIPSendTestAdapter are
// not required to implement Family (e.g. CLD's legacy send wrapper embeds the interface). With only the family,
// BuildExtraArgs uses the destination's GetExtraArgs, as before.
func extraArgsSource(src testadapters.ForkCCIPSendTestAdapter, family string) testadapters.ExtraArgsSource {
	if s, ok := src.(testadapters.ExtraArgsSource); ok {
		return s
	}
	return familyOnlySource(family)
}

var _ testadapters.ExtraArgsSource = familyOnlySource("")

type familyOnlySource string

func (f familyOnlySource) Family() string { return string(f) }
