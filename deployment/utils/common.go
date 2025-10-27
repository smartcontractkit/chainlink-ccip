package utils

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

const (
	BypasserManyChainMultisig  cldf.ContractType = "BypasserManyChainMultiSig"
	CancellerManyChainMultisig cldf.ContractType = "CancellerManyChainMultiSig"
	ProposerManyChainMultisig  cldf.ContractType = "ProposerManyChainMultiSig"
	RBACTimelock               cldf.ContractType = "RBACTimelock"
	CallProxy                  cldf.ContractType = "CallProxy"
	// CLL Identifiers
	CLLQualifier = "CLLCCIP"
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
