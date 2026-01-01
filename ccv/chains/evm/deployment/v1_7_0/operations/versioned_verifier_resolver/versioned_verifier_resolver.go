package versioned_verifier_resolver

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "VersionedVerifierResolver"

var Version = semver.MustParse("1.7.0")

type InboundImplementationArgs = versioned_verifier_resolver.VersionedVerifierResolverInboundImplementationArgs

type OutboundImplementationArgs = versioned_verifier_resolver.VersionedVerifierResolverOutboundImplementationArgs

type AcceptOwnershipArgs struct {
	IsProposedOwner bool
}

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "versioned-verifier-resolver:deploy",
	Version:          Version,
	Description:      "Deploys the VersionedVerifierResolver contract",
	ContractMetadata: versioned_verifier_resolver.VersionedVerifierResolverMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(cctp_verifier.ResolverType, *Version).String(): {
			EVM: common.FromHex(versioned_verifier_resolver.VersionedVerifierResolverBin),
		},
		cldf_deployment.NewTypeAndVersion(committee_verifier.ResolverType, *Version).String(): {
			EVM: common.FromHex(versioned_verifier_resolver.VersionedVerifierResolverBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyInboundImplementationUpdates = contract.NewWrite(contract.WriteParams[[]InboundImplementationArgs, *versioned_verifier_resolver.VersionedVerifierResolver]{
	Name:            "versioned-verifier-resolver:apply-inbound-implementation-updates",
	Version:         Version,
	Description:     "Updates verifier implementations for inbound traffic",
	ContractType:    ContractType,
	ContractABI:     versioned_verifier_resolver.VersionedVerifierResolverABI,
	NewContract:     versioned_verifier_resolver.NewVersionedVerifierResolver,
	IsAllowedCaller: contract.OnlyOwner[*versioned_verifier_resolver.VersionedVerifierResolver, []InboundImplementationArgs],
	Validate: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, backend bind.ContractBackend, opts *bind.CallOpts, args []InboundImplementationArgs) error {
		for _, arg := range args {
			if arg.Version == [4]byte{} {
				return errors.New("version cannot be the zero")
			}
		}
		return nil
	},
	IsNoop: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.CallOpts, args []InboundImplementationArgs) (bool, error) {
		for _, arg := range args {
			versionBytes := make([]byte, 4)
			copy(versionBytes, arg.Version[:])
			actualInboundImplementation, err := resolver.GetInboundImplementation(opts, versionBytes)
			if err != nil {
				return false, fmt.Errorf("failed to get inbound implementation for version %x: %w", arg.Version, err)
			}
			if actualInboundImplementation != arg.Verifier {
				return false, nil
			}
		}
		return false, nil
	},
	CallContract: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.TransactOpts, args []InboundImplementationArgs) (*types.Transaction, error) {
		return resolver.ApplyInboundImplementationUpdates(opts, args)
	},
})

var ApplyOutboundImplementationUpdates = contract.NewWrite(contract.WriteParams[[]OutboundImplementationArgs, *versioned_verifier_resolver.VersionedVerifierResolver]{
	Name:            "versioned-verifier-resolver:apply-outbound-implementation-updates",
	Version:         Version,
	Description:     "Updates verifier implementations for outbound traffic",
	ContractType:    ContractType,
	ContractABI:     versioned_verifier_resolver.VersionedVerifierResolverABI,
	NewContract:     versioned_verifier_resolver.NewVersionedVerifierResolver,
	IsAllowedCaller: contract.OnlyOwner[*versioned_verifier_resolver.VersionedVerifierResolver, []OutboundImplementationArgs],
	Validate: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, backend bind.ContractBackend, opts *bind.CallOpts, args []OutboundImplementationArgs) error {
		for _, arg := range args {
			if arg.DestChainSelector == 0 {
				return errors.New("dest chain selector cannot be 0")
			}
		}
		return nil
	},
	IsNoop: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.CallOpts, args []OutboundImplementationArgs) (bool, error) {
		for _, arg := range args {
			actualOutboundImplementation, err := resolver.GetOutboundImplementation(opts, arg.DestChainSelector, []byte{})
			if err != nil {
				return false, fmt.Errorf("failed to get outbound implementation for dest chain selector %d: %w", arg.DestChainSelector, err)
			}
			if actualOutboundImplementation != arg.Verifier {
				return false, nil
			}
		}
		return false, nil
	},
	CallContract: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.TransactOpts, args []OutboundImplementationArgs) (*types.Transaction, error) {
		return resolver.ApplyOutboundImplementationUpdates(opts, args)
	},
})

var AcceptOwnership = contract.NewWrite(contract.WriteParams[AcceptOwnershipArgs, *versioned_verifier_resolver.VersionedVerifierResolver]{
	Name:         "versioned-verifier-resolver:accept-ownership",
	Version:      Version,
	Description:  "Accept ownership of the versioned verifier resolver",
	ContractType: ContractType,
	ContractABI:  versioned_verifier_resolver.VersionedVerifierResolverABI,
	NewContract:  versioned_verifier_resolver.NewVersionedVerifierResolver,
	IsAllowedCaller: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.CallOpts, caller common.Address, args AcceptOwnershipArgs) (bool, error) {
		return args.IsProposedOwner, nil
	},
	Validate: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, backend bind.ContractBackend, opts *bind.CallOpts, args AcceptOwnershipArgs) error {
		return nil
	},
	IsNoop: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.CallOpts, args AcceptOwnershipArgs) (bool, error) {
		return false, nil
	},
	CallContract: func(resolver *versioned_verifier_resolver.VersionedVerifierResolver, opts *bind.TransactOpts, _ AcceptOwnershipArgs) (*types.Transaction, error) {
		return resolver.AcceptOwnership(opts)
	},
})
