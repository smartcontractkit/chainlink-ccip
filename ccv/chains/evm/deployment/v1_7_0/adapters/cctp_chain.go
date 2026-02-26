package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/cctp"
	cctp_through_ccv_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_through_ccv_token_pool"
	cctp_message_transmitter_proxy_v1_6_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/cctp_message_transmitter_proxy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ adapters.CCTPChain = &CCTPChainAdapter{}

// CCTPChainAdapter is the adapter for CCTP chains.
type CCTPChainAdapter struct{}

// DeployCCTPChain returns the sequence for deploying a CCTP chain.
func (c *CCTPChainAdapter) DeployCCTPChain() *operations.Sequence[adapters.DeployCCTPInput, seq_core.OnChainOutput, adapters.DeployCCTPChainDeps] {
	return cctp.DeployCCTPChain
}

// ConfigureCCTPChainForLanes returns the sequence for configuring a CCTP chain for lanes.
func (c *CCTPChainAdapter) ConfigureCCTPChainForLanes() *operations.Sequence[adapters.ConfigureCCTPChainForLanesInput, seq_core.OnChainOutput, adapters.ConfigureCCTPChainForLanesDeps] {
	return cctp.ConfigureCCTPChainForLanes
}

// CCTPV2AllowedCallerOnDest returns the address allowed to trigger message reception on the remote domain.
// On dest, the caller of CCTPV2 is the CCTPMessageTransmitterProxy 2.0.0.
func (c *CCTPChainAdapter) CCTPV2AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	allowedCallerOnDestAddressRef, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_message_transmitter_proxy.ContractType),
		Version: cctp_message_transmitter_proxy.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find allowed caller on dest address: %w", err)
	}
	return common.FromHex(allowedCallerOnDestAddressRef.Address), nil
}

// CCTPV1AllowedCallerOnDest returns the address allowed to trigger message reception on the remote domain.
// On dest, the caller of CCTPV1 is the CCTPMessageTransmitterProxy v1.6.2.
func (c *CCTPChainAdapter) CCTPV1AllowedCallerOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	allowedCallerOnDestAddressRef, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_message_transmitter_proxy_v1_6_2.ContractType),
		Version: cctp_message_transmitter_proxy_v1_6_2.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find allowed caller on dest address: %w", err)
	}
	return common.FromHex(allowedCallerOnDestAddressRef.Address), nil
}

// AllowedCallerOnSource returns the address allowed to deposit tokens for burn on the remote chain.
// On EVM, the caller of CCTP is the CCTPVerifier.
func (c *CCTPChainAdapter) AllowedCallerOnSource(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	allowedCallerOnSourceAddressRef, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_verifier.ContractType),
		Version: cctp_verifier.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find allowed caller on source address: %w", err)
	}
	return common.FromHex(allowedCallerOnSourceAddressRef.Address), nil
}

// MintRecipientOnDest returns the address that will receive tokens on the remote domain.
// On EVM, there is no mint recipient.
func (c *CCTPChainAdapter) MintRecipientOnDest(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil
}

// USDCType returns the type of the USDC on the chain.
func (c *CCTPChainAdapter) USDCType() adapters.USDCType {
	return adapters.Canonical
}

// PoolAddress returns the address of the token pool on the remote chain in bytes.
func (c *CCTPChainAdapter) PoolAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64, registeredPoolRef datastore.AddressRef) ([]byte, error) {
	registeredPoolAddress, err := datastore_utils.FindAndFormatRef(d, registeredPoolRef, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find registered pool address: %w", err)
	}
	return common.FromHex(registeredPoolAddress.Address), nil
}

// TokenAddress returns the address of the token on the remote chain in bytes.
func (c *CCTPChainAdapter) TokenAddress(d datastore.DataStore, b chain.BlockChains, chainSelector uint64) ([]byte, error) {
	poolAddressRef, err := datastore_utils.FindAndFormatRef(d, datastore.AddressRef{
		Type:    datastore.ContractType(cctp_through_ccv_token_pool.ContractType),
		Version: cctp_through_ccv_token_pool.Version,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return nil, fmt.Errorf("failed to find pool address: %w", err)
	}
	chain, ok := b.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	boundTokenPool, err := cctp_through_ccv_token_pool_bindings.NewCCTPThroughCCVTokenPool(common.HexToAddress(poolAddressRef.Address), chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind token pool: %w", err)
	}
	tokenAddress, err := boundTokenPool.GetToken(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get token address: %w", err)
	}
	return tokenAddress.Bytes(), nil
}
