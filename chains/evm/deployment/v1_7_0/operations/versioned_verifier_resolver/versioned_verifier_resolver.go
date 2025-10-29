package versioned_verifier_resolver

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "VersionedVerifierResolver"

type InboundImplementationArgs = versioned_verifier_resolver.VersionedVerifierResolverInboundImplementationArgs

type OutboundImplementationArgs = versioned_verifier_resolver.VersionedVerifierResolverOutboundImplementationArgs

var ApplyInboundImplementationUpdates = contract.NewWrite(contract.WriteParams[[]InboundImplementationArgs, *versioned_verifier_resolver.VersionedVerifierResolver]{
	Name:            "versioned-verifier-resolver:apply-inbound-implementation-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates verifier implementations for inbound traffic",
	ContractType:    ContractType,
	ContractABI:     versioned_verifier_resolver.VersionedVerifierResolverABI,
	NewContract:     versioned_verifier_resolver.NewVersionedVerifierResolver,
	IsAllowedCaller: contract.OnlyOwner[*versioned_verifier_resolver.VersionedVerifierResolver, []InboundImplementationArgs],
	Validate:        func([]InboundImplementationArgs) error { return nil },
	CallContract: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.TransactOpts, args []InboundImplementationArgs) (*types.Transaction, error) {
		return resolver.ApplyInboundImplementationUpdates(opts, args)
	},
})

var ApplyOutboundImplementationUpdates = contract.NewWrite(contract.WriteParams[[]OutboundImplementationArgs, *versioned_verifier_resolver.VersionedVerifierResolver]{
	Name:            "versioned-verifier-resolver:apply-outbound-implementation-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates verifier implementations for outbound traffic",
	ContractType:    ContractType,
	ContractABI:     versioned_verifier_resolver.VersionedVerifierResolverABI,
	NewContract:     versioned_verifier_resolver.NewVersionedVerifierResolver,
	IsAllowedCaller: contract.OnlyOwner[*versioned_verifier_resolver.VersionedVerifierResolver, []OutboundImplementationArgs],
	Validate:        func([]OutboundImplementationArgs) error { return nil },
	CallContract: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.TransactOpts, args []OutboundImplementationArgs) (*types.Transaction, error) {
		return resolver.ApplyOutboundImplementationUpdates(opts, args)
	},
})
