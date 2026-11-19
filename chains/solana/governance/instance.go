package governance

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

// ErrMCMSInstanceNotFound is returned when the datastore has no ref for a Solana timelock or MCM instance.
var ErrMCMSInstanceNotFound = errors.New("solana mcms instance not found in datastore")

// mcmsInstanceRefTypes maps the "<program>.<seed>" ref type of a Solana MCMS instance to its
// seed-only ref type and its program ref type. Deploys write both instance refs, but some
// datastores only have the "<program>.<seed>" one.
var mcmsInstanceRefTypes = map[cldf_datastore.ContractType]struct {
	seed, program cldf_datastore.ContractType
}{
	RBACTimelock:               {RBACTimelockSeed, TimelockProgramType},
	ProposerManyChainMultisig:  {ProposerSeed, McmProgramType},
	CancellerManyChainMultisig: {CancellerSeed, McmProgramType},
	BypasserManyChainMultisig:  {BypasserSeed, McmProgramType},
}

// MCMSInstance is a Solana timelock or MCM instance resolved from the datastore.
type MCMSInstance struct {
	ProgramID solana.PublicKey
	Seed      state.PDASeed
	// MissingRefs are the refs for this instance the datastore lacks, derived from the ref it has.
	MissingRefs []cldf_datastore.AddressRef
}

// ResolveMCMSInstance resolves a Solana timelock or MCM instance at the given version.
// instanceType is the "<program>.<seed>" ref type (RBACTimelock or one of the *ManyChainMultiSig
// types). When that ref is missing, the instance is rebuilt from the program ref and the seed ref.
// It returns ErrMCMSInstanceNotFound when neither the instance ref nor the seed ref exists.
func ResolveMCMSInstance(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	instanceType cldf_datastore.ContractType,
	version *semver.Version,
	qualifier string) (MCMSInstance, error) {
	refTypes, ok := mcmsInstanceRefTypes[instanceType]
	if !ok {
		return MCMSInstance{}, fmt.Errorf("unsupported solana mcms instance type %q", instanceType)
	}
	instanceRef := GetAddressRef(existingAddresses, chainSelector, instanceType, version, qualifier)
	seedRef := GetAddressRef(existingAddresses, chainSelector, refTypes.seed, version, qualifier)

	if instanceRef.Address != "" {
		id, seed, err := mcms_solana.ParseContractAddress(instanceRef.Address)
		if err != nil {
			return MCMSInstance{}, fmt.Errorf("invalid %s ref %q on chain %d: %w", instanceType, instanceRef.Address, chainSelector, err)
		}
		out := MCMSInstance{ProgramID: id, Seed: state.PDASeed(seed)}
		seedStr := string(bytes.TrimRight(seed[:], "\x00"))
		switch seedRef.Address {
		case "":
			out.MissingRefs = append(out.MissingRefs, cldf_datastore.AddressRef{
				Address:       seedStr,
				ChainSelector: chainSelector,
				Type:          refTypes.seed,
				Version:       version,
				Qualifier:     instanceRef.Qualifier,
			})
		case seedStr:
		default:
			return MCMSInstance{}, fmt.Errorf("%s ref %q and %s ref %q disagree on chain %d", instanceType, instanceRef.Address, refTypes.seed, seedRef.Address, chainSelector)
		}
		return out, nil
	}

	if seedRef.Address == "" {
		return MCMSInstance{}, fmt.Errorf("%w: no %s or %s ref on chain %d with qualifier %q", ErrMCMSInstanceNotFound, instanceType, refTypes.seed, chainSelector, qualifier)
	}
	programRef := GetAddressRef(existingAddresses, chainSelector, refTypes.program, version, "")
	if programRef.Address == "" {
		return MCMSInstance{}, fmt.Errorf("found %s ref but no %s ref on chain %d", refTypes.seed, refTypes.program, chainSelector)
	}
	id, err := solana.PublicKeyFromBase58(programRef.Address)
	if err != nil {
		return MCMSInstance{}, fmt.Errorf("invalid %s ref %q on chain %d: %w", refTypes.program, programRef.Address, chainSelector, err)
	}
	var seed state.PDASeed
	if len(seedRef.Address) > len(seed) {
		return MCMSInstance{}, fmt.Errorf("%s ref %q on chain %d is longer than %d bytes", refTypes.seed, seedRef.Address, chainSelector, len(seed))
	}
	copy(seed[:], seedRef.Address)
	return MCMSInstance{
		ProgramID: id,
		Seed:      seed,
		MissingRefs: []cldf_datastore.AddressRef{{
			Address:       mcms_solana.ContractAddress(id, mcms_solana.PDASeed(seed)),
			ChainSelector: chainSelector,
			Type:          instanceType,
			Version:       version,
			Qualifier:     seedRef.Qualifier,
		}},
	}, nil
}

// GetTimelockSignerPDA resolves the timelock instance at the given version and returns its signer PDA.
func GetTimelockSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	version *semver.Version,
	qualifier string) (solana.PublicKey, error) {
	timelock, err := ResolveMCMSInstance(existingAddresses, chainSelector, RBACTimelock, version, qualifier)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return state.GetTimelockSignerPDA(timelock.ProgramID, timelock.Seed), nil
}

// GetMCMSignerPDA resolves an MCM instance at the given version and returns its signer PDA.
func GetMCMSignerPDA(
	existingAddresses []cldf_datastore.AddressRef,
	chainSelector uint64,
	signerType cldf_datastore.ContractType,
	version *semver.Version,
	qualifier string) (solana.PublicKey, error) {
	mcm, err := ResolveMCMSInstance(existingAddresses, chainSelector, signerType, version, qualifier)
	if err != nil {
		return solana.PublicKey{}, err
	}
	return state.GetMCMSignerPDA(mcm.ProgramID, mcm.Seed), nil
}
