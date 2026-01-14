package utils

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

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
	RMNHome                    cldf.ContractType = "RMNHome"
	BurnMintTokenPool          cldf.ContractType = "BurnMintTokenPool"
	LockReleaseTokenPool       cldf.ContractType = "LockReleaseTokenPool"
	TokenPoolLookupTable       cldf.ContractType = "TokenPoolLookupTable"
	BurnWithFromMintTokenPool  cldf.ContractType = "BurnWithFromMintTokenPool"
	BurnFromMintTokenPool      cldf.ContractType = "BurnFromMintTokenPool"
	// CLL Identifiers
	CLLQualifier         = "CLLCCIP"
	RMNTimelockQualifier = "RMNMCMS"

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
	SuiFamilySelector   = "c4e05953"
)

func GetSelectorHex(selector uint64) []byte {
	destFamily, _ := chain_selectors.GetSelectorFamily(selector)
	var familySelector []byte
	switch destFamily {
	case chain_selectors.FamilyEVM:
		familySelector, _ = hex.DecodeString(EVMFamilySelector)
	case chain_selectors.FamilySolana:
		familySelector, _ = hex.DecodeString(SVMFamilySelector)
	case chain_selectors.FamilyAptos:
		familySelector, _ = hex.DecodeString(AptosFamilySelector)
	case chain_selectors.FamilyTon:
		familySelector, _ = hex.DecodeString(TVMFamilySelector)
	case chain_selectors.FamilySui:
		familySelector, _ = hex.DecodeString(SuiFamilySelector)
	}
	return familySelector
}

var (
	ErrZeroAddress         = errors.New("address cannot be zero address")
	ErrNoAdapterRegistered = func(family string, version *semver.Version) error {
		return fmt.Errorf("no adapter registered for chain family %s and version %s", family, version.String())
	}
)

var (
	Version_1_0_0 = semver.MustParse("1.0.0")
	Version_1_5_0 = semver.MustParse("1.5.0")
	Version_1_5_1 = semver.MustParse("1.5.1")
	Version_1_6_0 = semver.MustParse("1.6.0")
	Version_1_6_1 = semver.MustParse("1.6.1")
)

func NewRegistererID(chainFamily string, version *semver.Version) string {
	return fmt.Sprintf("%s-%s", chainFamily, version.String())
}

const (
	EXECUTION_STATE_UNTOUCHED  = 0
	EXECUTION_STATE_INPROGRESS = 1
	EXECUTION_STATE_SUCCESS    = 2
	EXECUTION_STATE_FAILURE    = 3
)