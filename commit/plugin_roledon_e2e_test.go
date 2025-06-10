package commit

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
)

type roleDonTestSetup struct {
	sourceChains []uint64
	destChain    uint64
	oracles      []commontypes.OracleID
	fRoleDon     int
	fChain       map[uint64]int
	chainOracles map[uint64][]commontypes.OracleID
}

func (s roleDonTestSetup) String() string {
	return fmt.Sprintf("Role DON Test Setup:\n"+
		"Source Chains: %v\n"+
		"Dest Chain: %v\n"+
		"Oracles: %v (%d)\n"+
		"FRoleDon: %v  // floor(len(oracles)-1/2)\n"+
		"FChain: %v\n"+
		"ChainOracles: %v\n",
		s.sourceChains,
		s.destChain,
		s.oracles, len(s.oracles),
		s.fRoleDon,
		s.fChain,
		s.chainOracles,
	)
}

func TestPlugin_RoleDonE2E(t *testing.T) {
	s := roleDonTestSetup{}

	s.sourceChains = []uint64{
		chainsel.TEST_1000.Selector,
		chainsel.TEST_1338.Selector,
	}

	s.destChain = chainsel.TEST_76578.Selector

	s.oracles = make([]commontypes.OracleID, 7)
	for i := range s.oracles {
		s.oracles[i] = commontypes.OracleID(i + 1)
	}

	s.fRoleDon = int(math.Floor((float64(len(s.oracles)) - 1.0) / 2.0))

	s.fChain = map[uint64]int{
		s.sourceChains[0]: 1,
		s.sourceChains[1]: 1,
		s.destChain:       1,
	}

	s.chainOracles = map[uint64][]commontypes.OracleID{}
	for chainSel, f := range s.fChain {
		numRequiredOracles := 2*f + 1
		s.chainOracles[chainSel] = getRandomPermutation(s.oracles, numRequiredOracles+rand.Intn(1))
	}

	t.Logf("%s", s)
}

func getRandomPermutation[T any](sl []T, lim int) []T {
	var cp []T
	for i := 0; i < lim; i++ {
		cp = append(cp, sl[i])
	}
	rand.Shuffle(len(sl), func(i, j int) { sl[i], sl[j] = sl[j], sl[i] })
	return cp[:lim]
}
