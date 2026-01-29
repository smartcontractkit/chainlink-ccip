package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var GetTokensPaginationSize = uint64(20)

type ConfigImportAdapter struct {
	FeeQuoter     map[uint64]common.Address
	Router        map[uint64]common.Address
	TokenAdminReg map[uint64]common.Address
}

func (ci *ConfigImportAdapter) InitializeAdapter(e cldf.Environment, selectors []uint64) error {
	ci.FeeQuoter = make(map[uint64]common.Address)
	ci.Router = make(map[uint64]common.Address)
	for _, chainSelector := range selectors {
		fqRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:          datastore.ContractType(fqops.ContractType),
			Version:       fqops.Version,
			ChainSelector: chainSelector,
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return fmt.Errorf("failed to find fee quoter contract ref for chain %d: %w", chainSelector, err)
		}
		ci.FeeQuoter[chainSelector] = fqRef
		routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:          datastore.ContractType("Router"),
			Version:       semver.MustParse("1.2.0"),
			ChainSelector: chainSelector,
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return fmt.Errorf("failed to find router contract ref for chain %d: %w", chainSelector, err)
		}
		ci.Router[chainSelector] = routerRef
		tokenAdminRegRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
			Type:          datastore.ContractType("TokenAdminRegistry"),
			Version:       semver.MustParse("1.5.0"),
			ChainSelector: chainSelector,
		}, chainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return fmt.Errorf("failed to find token admin registry contract ref for chain %d: %w", chainSelector, err)
		}
		ci.TokenAdminReg[chainSelector] = tokenAdminRegRef
	}
	return nil
}

func (ci *ConfigImportAdapter) SupportedTokens(e cldf.Environment, chainsel uint64) ([]common.Address, error) {
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	tokenAdminRegAddr, ok := ci.TokenAdminReg[chainsel]
	if !ok {
		return nil, fmt.Errorf("token admin registry address not found for chain %d", chainsel)
	}
	// get all supported tokens from token admin registry
	tokenAdminRegC, err := token_admin_registry.NewTokenAdminRegistry(tokenAdminRegAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token admin registry contract at %s on chain %d: %w", tokenAdminRegAddr.String(), chain.Selector, err)
	}
	startIndex := uint64(0)
	allTokens := make([]common.Address, 0)
	for {
		fetchedTokens, err := tokenAdminRegC.GetAllConfiguredTokens(nil, startIndex, GetTokensPaginationSize)
		if err != nil {
			return nil, err
		}
		allTokens = append(allTokens, fetchedTokens...)
		startIndex += GetTokensPaginationSize
		if uint64(len(fetchedTokens)) < GetTokensPaginationSize {
			break
		}
	}
	return allTokens, nil
}

func (ci *ConfigImportAdapter) ConnectedChains(e cldf.Environment, chainsel uint64) ([]uint64, error) {
	chain, ok := e.BlockChains.EVMChains()[chainsel]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found in environment", chainsel)
	}
	routerAddr, ok := ci.Router[chainsel]
	if !ok {
		return nil, fmt.Errorf("router address not found for chain %d", chainsel)
	}
	// get all offRamps from router to find connected chains
	routerC, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate router contract at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	offRamps, err := routerC.GetOffRamps(&bind.CallOpts{
		Context: e.GetContext(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramps from router at %s on chain %d: %w", routerAddr.String(), chain.Selector, err)
	}
	connectedChains := make([]uint64, 0)
	for _, offRamp := range offRamps {
		if offRamp.OffRamp == (common.Address{}) {
			continue // skip uninitialized off-ramps
		}
		connectedChains = append(connectedChains, offRamp.SourceChainSelector)
	}
	return connectedChains, nil
}

func (ci *ConfigImportAdapter) SequenceImportConfig() *cldf_ops.Sequence[api.ImportConfigPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"seq-config-import",
		semver.MustParse("1.0.0"),
		"Imports configuration for specified chains",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.ImportConfigPerChainInput) (output sequences.OnChainOutput, err error) {
			evmChain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", in.ChainSelector)
			}
			chainSelector := in.ChainSelector
			b.Logger.Infof("Importing configuration for chain %d (%s)", chainSelector, evmChain.Name())
			var contractMetadata []datastore.ContractMetadata
			// read FQ config from onchain
			fqAddress, ok := ci.FeeQuoter[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("fee quoter address not found for chain %d", chainSelector)
			}
			destChainConfigs := make(map[uint64]fee_quoter.FeeQuoterDestChainConfig)
			for _, remoteChain := range in.RemoteChains {
				opsOutput, err := cldf_ops.ExecuteOperation(b, fqops.GetDestChainConfig, evmChain, contract.FunctionInput[uint64]{
					Address:       fqAddress,
					ChainSelector: chainSelector,
					Args:          remoteChain,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for "+
						"remote chain %d from feequoter %s on chain %d: %w",
						remoteChain, fqAddress.Hex(), chainSelector, err)
				}
				destChainConfigs[remoteChain] = opsOutput.Output
			}
			for _, token := range in.Tokens {

			}
			contractMetadata = append(contractMetadata, datastore.ContractMetadata{
				Address:       fqAddress.Hex(),
				ChainSelector: chainSelector,
				Metadata: struct {
					DestChainConfigs map[uint64]fee_quoter.FeeQuoterDestChainConfig
				}{
					DestChainConfigs: destChainConfigs,
				},
			})
			return sequences.OnChainOutput{}, nil
		},
	)
}
