package ccv_verifier_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/verifier_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "VerifierProxy"

var SetRamp = contract.NewWrite(contract.WriteParams[common.Address, *verifier_proxy.VerifierProxy]{
	Name:            "verifier-proxy:set-ramp",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Set the verifier address on the verifier proxy",
	ContractType:    ContractType,
	ContractABI:     verifier_proxy.VerifierProxyABI,
	NewContract:     verifier_proxy.NewVerifierProxy,
	IsAllowedCaller: contract.OnlyOwner[*verifier_proxy.VerifierProxy],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(verifierProxy *verifier_proxy.VerifierProxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return verifierProxy.SetVerifier(opts, args)
	},
})
