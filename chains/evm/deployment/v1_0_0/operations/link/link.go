package link

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/link_token"
)

var ContractType cldf_deployment.ContractType = "LinkToken"
var Version = semver.MustParse("1.0.0")

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "link:deploy",
	Version:          Version,
	Description:      "Deploys the LINK token contract",
	ContractMetadata: link_token.LinkTokenMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(link_token.LinkTokenBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

// GrantMintRoleArgs contains the arguments for granting mint role.
type GrantMintRoleArgs struct {
	Minter common.Address
}

func NewWriteGrantMintRole(c *link_token.LinkToken) *cld_ops.Operation[contract.FunctionInput[GrantMintRoleArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[GrantMintRoleArgs, *link_token.LinkToken]{
		Name:            "link:grant-mint-role",
		Version:         Version,
		Description:     "Grants minting permission to an address",
		ContractType:    ContractType,
		ContractABI:     link_token.LinkTokenABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*link_token.LinkToken, GrantMintRoleArgs],
		Validate:        func(GrantMintRoleArgs) error { return nil },
		CallContract: func(linkToken *link_token.LinkToken, opts *bind.TransactOpts, args GrantMintRoleArgs) (*types.Transaction, error) {
			return linkToken.GrantMintRole(opts, args.Minter)
		},
	})
}

// MintArgs contains the arguments for minting LINK tokens.
type MintArgs struct {
	To     common.Address
	Amount *big.Int
}

func NewWriteMint(c *link_token.LinkToken) *cld_ops.Operation[contract.FunctionInput[MintArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[MintArgs, *link_token.LinkToken]{
		Name:            "link:mint",
		Version:         Version,
		Description:     "Mints LINK tokens to the specified address",
		ContractType:    ContractType,
		ContractABI:     link_token.LinkTokenABI,
		Contract:        c,
		IsAllowedCaller: contract.AllCallersAllowed[*link_token.LinkToken, MintArgs], // Minter check is done on-chain
		Validate:        func(MintArgs) error { return nil },
		CallContract: func(linkToken *link_token.LinkToken, opts *bind.TransactOpts, args MintArgs) (*types.Transaction, error) {
			return linkToken.Mint(opts, args.To, args.Amount)
		},
	})
}
