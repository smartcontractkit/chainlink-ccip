package tokenimpl

import (
	bnmERC20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	dripV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	dripV1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var tokenImpls = map[deployment.ContractType]Token{
	dripV1_5_0.ContractType: tokenBurnMintERC20WithDripV1_5_0{},
	dripV1_0_0.ContractType: tokenBurnMintERC20WithDripV1_0_0{},
	utils.ERC677TokenHelper: tokenBurnMintERC677{},
	utils.BurnMintToken:     tokenBurnMintERC677{},
	bnmERC20.ContractType:   tokenBurnMintERC20{},
	erc20.ContractType:      tokenERC20{},
	tip20.ContractType:      tokenTIP20{},
}

// Get returns the token implementation for an EVM token contract type.
func Get(ct deployment.ContractType) (Token, bool) {
	s, ok := tokenImpls[ct]
	return s, ok
}

// Capabilities returns the capability set for an EVM token contract type, or the zero value if the token implementation does not exist.
func Capabilities(ct deployment.ContractType) CapabilitySet {
	if s, ok := Get(ct); ok {
		return s.Capabilities()
	}
	return CapabilitySet{}
}
