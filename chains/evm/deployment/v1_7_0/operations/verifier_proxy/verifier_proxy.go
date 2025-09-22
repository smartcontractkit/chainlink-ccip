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

var SetRamp = contract.NewWrite(
	"verifier-proxy:set-ramp",
	semver.MustParse("1.7.0"),
	"Set the verifier address on the verifier proxy",
	ContractType,
	verifier_proxy.VerifierProxyABI,
	verifier_proxy.NewVerifierProxy,
	contract.OnlyOwner,
	func(common.Address) error { return nil },
	func(verifierProxy *verifier_proxy.VerifierProxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return verifierProxy.SetVerifier(opts, args)
	},
)
