package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/verifier_test_helper"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/test_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

var _ ccvadapters.TestVerifierChainAdapter = (*EVMTestVerifierChainAdapter)(nil)

// EVMTestVerifierChainAdapter implements TestVerifierChainAdapter for EVM chains.
type EVMTestVerifierChainAdapter struct{}

// DeployTestVerifierChain returns the sequence that deploys test verifier infrastructure on an EVM chain.
func (a *EVMTestVerifierChainAdapter) DeployTestVerifierChain() *cldf_ops.Sequence[ccvadapters.DeployTestVerifierChainInput, sequences.OnChainOutput, ccvadapters.DeployTestVerifierChainDeps] {
	return test_verifier.DeployTestVerifierChain
}

// ConfigureTestVerifierChainForLanes returns the sequence that configures test verifier lanes on an EVM chain.
func (a *EVMTestVerifierChainAdapter) ConfigureTestVerifierChainForLanes() *cldf_ops.Sequence[ccvadapters.ConfigureTestVerifierForLanesInput, sequences.OnChainOutput, ccvadapters.ConfigureTestVerifierForLanesDeps] {
	return test_verifier.ConfigureTestVerifierChainForLanes
}

// TokenAddress returns the TESTVTR token address on the given chain in bytes.
func (a *EVMTestVerifierChainAdapter) TokenAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	ref, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:      datastore.ContractType("BurnMintERC20WithDrip"),
		Qualifier: "TESTVTR",
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return nil, fmt.Errorf("resolve TESTVTR token on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// PoolAddress returns the TESTVTR pool address on the given chain in bytes.
func (a *EVMTestVerifierChainAdapter) PoolAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	ref, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
		Version:   burn_mint_token_pool.Version,
		Qualifier: "TESTVTR",
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return nil, fmt.Errorf("resolve TESTVTR pool on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// VerifierAddress returns the test verifier address on the given chain.
func (a *EVMTestVerifierChainAdapter) VerifierAddress(d datastore.DataStore, chainSelector uint64) ([]byte, error) {
	ref, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:    datastore.ContractType(verifier_test_helper.ContractType),
		Version: verifier_test_helper.Version,
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return nil, fmt.Errorf("resolve test verifier on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// VerifierResolverAddress returns the test verifier resolver address on the given chain.
func (a *EVMTestVerifierChainAdapter) VerifierResolverAddress(d datastore.DataStore, chainSelector uint64) ([]byte, error) {
	ref, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:      datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
		Version:   versioned_verifier_resolver.Version,
		Qualifier: test_verifier.TestVerifierResolverQualifier,
	}, chainSelector, evmds.ToEVMAddressBytes)
	if err != nil {
		return nil, fmt.Errorf("resolve test verifier resolver on chain %d: %w", chainSelector, err)
	}
	return ref, nil
}

// Ensure EVMTestVerifierChainAdapter satisfies the RemoteTestVerifierChain interface.
var _ ccvadapters.RemoteTestVerifierChain = (*EVMTestVerifierChainAdapter)(nil)

// Compile-time check that the version tag is consistent.
var _ = semver.MustParse("2.0.0")

// compile-time check: common is used.
var _ = common.Address{}
