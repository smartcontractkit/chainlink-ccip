package hooks

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider/rpcclient"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config"
	cfgnet "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	seq2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	cciphooks "github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

var _ cciphooks.ContractOwnership = (*EVMContractOwnership)(nil)

func init() {
	cciphooks.GetContractOwnershipRegistry().Register(chain_selectors.FamilyEVM, &EVMContractOwnership{})
}

var laneMigratorContractTypesForOwnershipCheck = map[datastore.ContractType]struct{}{
	datastore.ContractType(committee_verifier.ContractType):   {},
	datastore.ContractType(executor.ContractType):             {},
	datastore.ContractType(seq2_0.ExecutorProxyType):          {},
	datastore.ContractType(onrampops.ContractType):            {},
	datastore.ContractType(offrampops.ContractType):           {},
	datastore.ContractType(fqops.ContractType):                {},
	datastore.ContractType(routerops.ContractType):            {},
	datastore.ContractType(rmn_remote.ContractType):           {},
	datastore.ContractType(rmn_proxy.ContractType):            {},
	datastore.ContractType(token_admin_registry.ContractType): {},
}

// EVMContractOwnership validates that contracts are owned by expected timelocks.
type EVMContractOwnership struct {
	cllccipTimelockAddr map[uint64]common.Address
	rmntimelockAddr     map[uint64]common.Address
}

func (e *EVMContractOwnership) initializeTimelocksInOwnershipCheck(ds datastore.DataStore, chainSelector uint64) error {
	if e.rmntimelockAddr == nil {
		e.rmntimelockAddr = make(map[uint64]common.Address)
	}
	if e.cllccipTimelockAddr == nil {
		e.cllccipTimelockAddr = make(map[uint64]common.Address)
	}
	cllTL, clltlExists := e.cllccipTimelockAddr[chainSelector]
	rmnTL, rmntlExists := e.rmntimelockAddr[chainSelector]
	if clltlExists && rmntlExists && cllTL != (common.Address{}) && rmnTL != (common.Address{}) {
		return nil
	}
	cllccipTimelockAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(common_utils.RBACTimelock),
		Qualifier: common_utils.CLLQualifier,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("ownership transfer requires CLLCCIP RBACTimelock in ExistingAddresses: %w", err)
	}

	rmnTimelockAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:      datastore.ContractType(common_utils.RBACTimelock),
		Qualifier: common_utils.RMNTimelockQualifier,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return fmt.Errorf("ownership transfer requires RMNMCMS RBACTimelock in ExistingAddresses: %w", err)
	}
	e.cllccipTimelockAddr[chainSelector] = cllccipTimelockAddr
	e.rmntimelockAddr[chainSelector] = rmnTimelockAddr
	return nil
}

func (e *EVMContractOwnership) FilterNetworks(envName string, dom domain.Domain, lggr logger.Logger) (*cfgnet.Config, error) {
	networkCfg, err := config.LoadNetworks(envName, dom, lggr)
	if err != nil {
		return nil, err
	}
	return networkCfg.FilterWith(cfgnet.ChainFamilyFilter(chain_selectors.FamilyEVM)), nil
}

func (e *EVMContractOwnership) NeedsOwnershipCheck(ref datastore.AddressRef) bool {
	_, exists := laneMigratorContractTypesForOwnershipCheck[ref.Type]
	return exists
}

func (e *EVMContractOwnership) expectedOwnerForRef(ref datastore.AddressRef) common.Address {
	switch ref.Type {
	case datastore.ContractType(rmn_remote.ContractType):
		return e.rmntimelockAddr[ref.ChainSelector]
	default:
		return e.cllccipTimelockAddr[ref.ChainSelector]
	}
}

func (e *EVMContractOwnership) VerifyContractOwnership(
	ctx context.Context,
	lggr logger.Logger,
	network cfgnet.Network,
	refsToCheck []datastore.AddressRef,
) error {
	if len(network.RPCs) == 0 || network.RPCs[0].HTTPURL == "" {
		return fmt.Errorf("network %d has no HTTP RPC configured", network.ChainSelector)
	}
	// TODO use blockchains from Env when chains are included in hookEnv
	rpcCfg := rpcclient.RPCConfig{
		ChainSelector: network.ChainSelector,
	}
	for _, rpc := range network.RPCs {
		p, err := rpcclient.URLSchemePreferenceFromString(rpc.PreferredURLScheme)
		if err != nil {
			return fmt.Errorf("invalid preferred URL scheme for RPC %s on network %d: %w", rpc.RPCName, network.ChainSelector, err)
		}
		rpcCfg.RPCs = append(rpcCfg.RPCs, rpcclient.RPC{
			Name:               rpc.RPCName,
			WSURL:              rpc.WSURL,
			HTTPURL:            rpc.HTTPURL,
			PreferredURLScheme: p,
		})
	}
	client, err := rpcclient.NewMultiClient(lggr, rpcCfg)
	if err != nil {
		return fmt.Errorf("dial RPC for chain %d: %w", network.ChainSelector, err)
	}
	defer client.Close()

	for _, ref := range refsToCheck {
		addr, err := evm_datastore_utils.ToEVMAddress(ref)
		if err != nil {
			return fmt.Errorf("error formatting address ref %s for contract type %s version %s on chain %d: %w",
				ref.Address, ref.Type, ref.Version, network.ChainSelector, err)
		}
		currentOwner, _, err := mcms_seq.LoadOwnableContract(addr, client)
		if err != nil {
			return fmt.Errorf("failed to load ownable contract %s (%s): %w", addr, ref.Type, err)
		}
		if currentOwner != e.expectedOwnerForRef(ref) {
			return fmt.Errorf("ownership check failed for contract %s (%s): expected owner %s, got %s",
				addr, ref.Type, e.expectedOwnerForRef(ref), currentOwner)
		}
		lggr.Infof("ownership check passed for contract %s (%s): owner is %s", addr, ref.Type, currentOwner)
	}
	return nil
}
