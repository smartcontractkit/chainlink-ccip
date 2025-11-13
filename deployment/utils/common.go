package utils

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type ConfigType string

const (
	ConfigTypeActive           ConfigType        = "active"
	ConfigTypeCandidate        ConfigType        = "candidate"
	BypasserManyChainMultisig  cldf.ContractType = "BypasserManyChainMultiSig"
	CancellerManyChainMultisig cldf.ContractType = "CancellerManyChainMultiSig"
	ProposerManyChainMultisig  cldf.ContractType = "ProposerManyChainMultiSig"
	RBACTimelock               cldf.ContractType = "RBACTimelock"
	CallProxy                  cldf.ContractType = "CallProxy"
	CapabilitiesRegistry       cldf.ContractType = "CapabilitiesRegistry"
	CCIPHome                   cldf.ContractType = "CCIPHome"
	// CLL Identifiers
	CLLQualifier = "CLLCCIP"

	// https://github.com/smartcontractkit/chainlink/blob/1423e2581e8640d9e5cd06f745c6067bb2893af2/contracts/src/v0.8/ccip/libraries/Internal.sol#L275-L279
	/*
				```Solidity
					// bytes4(keccak256("CCIP ChainFamilySelector EVM"))
					bytes4 public constant CHAIN_FAMILY_SELECTOR_EVM = 0x2812d52c;
					// bytes4(keccak256("CCIP ChainFamilySelector SVM"));
		  		bytes4 public constant CHAIN_FAMILY_SELECTOR_SVM = 0x1e10bdc4;
				```
	*/
	EVMFamilySelector   = "2812d52c"
	SVMFamilySelector   = "1e10bdc4"
	AptosFamilySelector = "ac77ffec"
	TVMFamilySelector   = "647e2ba9"
)

var (
	ErrZeroAddress         = errors.New("address cannot be zero address")
	ErrNoAdapterRegistered = func(family string, version *semver.Version) error {
		return fmt.Errorf("no adapter registered for chain family %s and version %s", family, version.String())
	}
)

var (
	Version_1_6_0 = semver.MustParse("1.6.0")
)

func NewRegistererID(chainFamily string, version *semver.Version) string {
	return fmt.Sprintf("%s-%s", chainFamily, version.String())
}
