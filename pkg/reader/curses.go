package reader

import (
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// CurseInfo contains cursing information that are fetched from the rmn remote contract.
type CurseInfo struct {
	// CursedSourceChains contains the cursed source chains.
	CursedSourceChains map[ccipocr3.ChainSelector]bool
	// CursedDestination indicates that the destination chain is cursed.
	CursedDestination bool
	// GlobalCurse indicates that all chains are cursed.
	GlobalCurse bool
}

func (ci CurseInfo) NonCursedSourceChains(inputChains []ccipocr3.ChainSelector) []ccipocr3.ChainSelector {
	if ci.GlobalCurse {
		return nil
	}

	sourceChains := make([]ccipocr3.ChainSelector, 0, len(inputChains))
	for _, ch := range inputChains {
		if !ci.CursedSourceChains[ch] {
			sourceChains = append(sourceChains, ch)
		}
	}
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })

	return sourceChains
}

// GlobalCurseSubject Defined as a const in RMNRemote.sol
// Docs of RMNRemote:
// An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject
// for issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of
// using the local chain selector as a subject.
var GlobalCurseSubject = [16]byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
}
