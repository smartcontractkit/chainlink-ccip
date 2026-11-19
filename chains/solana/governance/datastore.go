package governance

import (
	"github.com/Masterminds/semver/v3"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// GetAddressRef returns the first ref matching the chain, type and version, and the qualifier
// when one is given (an empty qualifier matches any). It returns an empty ref when none matches.
//
// This has the same logic as chainlink-ccip/deployment/utils/datastore.GetAddressRef, which this
// module cannot import without pulling in the whole deployment module.
func GetAddressRef(
	input []cldf_datastore.AddressRef,
	selector uint64,
	contractType cldf_datastore.ContractType,
	contractVersion *semver.Version,
	contractQualifier string,
) cldf_datastore.AddressRef {
	for _, ref := range input {
		if ref.ChainSelector == selector &&
			ref.Type == contractType &&
			ref.Version.Equal(contractVersion) &&
			(contractQualifier == "" || ref.Qualifier == contractQualifier) {
			return ref
		}
	}
	return cldf_datastore.AddressRef{}
}

// GetAllMCMS returns the MCMS refs of a chain in a fixed order: access controller program,
// proposer/executor/canceller/bypasser access controller accounts, timelock, then the proposer,
// canceller and bypasser MCMs. A ref that is not in the datastore comes back as an empty ref in
// its place.
func GetAllMCMS(
	version *semver.Version,
	chainSelector uint64,
	qualifier string,
	existingAddresses []cldf_datastore.AddressRef) []cldf_datastore.AddressRef {
	accessControllerProgram := GetAddressRef(
		existingAddresses, chainSelector, AccessControllerProgramType, version, "")
	proposerAccount := GetAddressRef(
		existingAddresses, chainSelector, ProposerAccessControllerAccount, version, qualifier)
	executorAccount := GetAddressRef(
		existingAddresses, chainSelector, ExecutorAccessControllerAccount, version, qualifier)
	cancellerAccount := GetAddressRef(
		existingAddresses, chainSelector, CancellerAccessControllerAccount, version, qualifier)
	bypasserAccount := GetAddressRef(
		existingAddresses, chainSelector, BypasserAccessControllerAccount, version, qualifier)
	timelockProgram := GetAddressRef(
		existingAddresses, chainSelector, RBACTimelock, version, qualifier)
	proposerMCMSAccount := GetAddressRef(
		existingAddresses, chainSelector, ProposerManyChainMultisig, version, qualifier)
	cancellerMCMSAccount := GetAddressRef(
		existingAddresses, chainSelector, CancellerManyChainMultisig, version, qualifier)
	bypasserMCMSAccount := GetAddressRef(
		existingAddresses, chainSelector, BypasserManyChainMultisig, version, qualifier)
	return []cldf_datastore.AddressRef{
		accessControllerProgram,
		proposerAccount,
		executorAccount,
		cancellerAccount,
		bypasserAccount,
		timelockProgram,
		proposerMCMSAccount,
		cancellerMCMSAccount,
		bypasserMCMSAccount,
	}
}
