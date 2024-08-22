package mocks

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/mock"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type ChainSupport struct {
	*mock.Mock
}

func NewChainSupport() *ChainSupport {
	return &ChainSupport{
		Mock: &mock.Mock{},
	}
}

func (c ChainSupport) KnownSourceChainsSlice() ([]cciptypes.ChainSelector, error) {
	args := c.Called()
	return args.Get(0).([]cciptypes.ChainSelector), args.Error(1)
}

// SupportedChains returns the set of chains that the given Oracle is configured to access
func (c ChainSupport) SupportedChains(oracleID commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	args := c.Called(oracleID)
	return args.Get(0).(mapset.Set[cciptypes.ChainSelector]), args.Error(1)
}

// SupportsDestChain returns true if the given oracle supports the dest chain, returns false otherwise
func (c ChainSupport) SupportsDestChain(oracleID commontypes.OracleID) (bool, error) {
	args := c.Called(oracleID)
	return args.Get(0).(bool), args.Error(1)
}

// SupportsChain returns true if the given oracle supports the given chain, returns false otherwise
func (c ChainSupport) SupportsChain(oracleID commontypes.OracleID, selector cciptypes.ChainSelector) (bool, error) {
	args := c.Called(oracleID, selector)
	return args.Get(0).(bool), args.Error(1)
}
