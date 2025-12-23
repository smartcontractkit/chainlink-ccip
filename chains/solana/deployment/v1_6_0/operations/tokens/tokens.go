package tokens

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	soltokens "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var LinkContractType cldf_deployment.ContractType = "LINK"
var Version *semver.Version = semver.MustParse("1.6.0")

type Params struct {
	ExistingAddresses      []datastore.AddressRef
	TokenProgramName       cldf_deployment.ContractType
	TokenPrivKey           solana.PrivateKey
	TokenDecimals          uint8
	TokenSymbol            string
	ATAList                []solana.PublicKey
	DisableFreezeAuthority bool
}

var DeployLINK = operations.NewOperation(
	"link:deploy",
	Version,
	"Deploys the LINK token contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (datastore.AddressRef, error) {
		instructions, err := tokens.CreateToken(
			context.Background(),
			solana.TokenProgramID,
			input.TokenPrivKey.PublicKey(),
			chain.DeployerKey.PublicKey(),
			input.TokenDecimals,
			chain.Client,
			cldf_solana.SolDefaultCommitment,
		)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		err = chain.Confirm(instructions, common.AddSigners(input.TokenPrivKey))
		if err != nil {
			return datastore.AddressRef{}, err
		}
		return datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       input.TokenPrivKey.PublicKey().String(),
			Type:          datastore.ContractType(LinkContractType),
		}, nil
	},
)

var DeploySolanaToken = operations.NewOperation(
	"solana-token:deploy",
	Version,
	"Deploys and configures an SPL token contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (datastore.AddressRef, error) {
		// CREATE TOKEN
		tokenProgramID, err := utils.GetTokenProgramID(input.TokenProgramName)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		freezeAuthority := utils.GetTimelockSignerPDA(input.ExistingAddresses, chain.Selector, common_utils.CLLQualifier)
		tokenAdminPubKey := chain.DeployerKey.PublicKey()
		// if we're disabling the freeze authority, we first set it to the deployer key so it can
		// immediately revoke it
		if input.DisableFreezeAuthority {
			freezeAuthority = chain.DeployerKey.PublicKey()
		}
		var mint solana.PublicKey
		privKey := input.TokenPrivKey
		if privKey.IsValid() {
			mint = privKey.PublicKey()
		} else {
			privKey, err = solana.NewRandomPrivateKey()
			if err != nil {
				return datastore.AddressRef{}, err
			}
			mint = privKey.PublicKey()
		}
		instructions, err := soltokens.CreateTokenWith(
			context.Background(),
			tokenProgramID,
			mint,
			tokenAdminPubKey,
			freezeAuthority,
			input.TokenDecimals,
			chain.Client,
			cldf_solana.SolDefaultCommitment,
			false,
		)
		if err != nil {
			return datastore.AddressRef{}, err
		}
		err = chain.Confirm(instructions, common.AddSigners(privKey))
		if err != nil {
			return datastore.AddressRef{}, err
		}
		// CREATE ATAs
		for _, ata := range input.ATAList {
			createATAIx, _, err := soltokens.CreateAssociatedTokenAccount(
				tokenProgramID,
				mint,
				ata,
				chain.DeployerKey.PublicKey(),
			)
			if err != nil {
				return datastore.AddressRef{}, err
			}
			if err := chain.Confirm([]solana.Instruction{createATAIx}); err != nil {
				return datastore.AddressRef{}, err
			}
		}
		// DISABLE FREEZE AUTHORITY
		if input.DisableFreezeAuthority {
			err = utils.DisableFreezeAuthority(chain, []solana.PublicKey{mint})
			if err != nil {
				return datastore.AddressRef{}, err
			}
		}
		return datastore.AddressRef{
			ChainSelector: chain.Selector,
			Address:       mint.String(),
			Type:          datastore.ContractType(input.TokenProgramName),
			Qualifier:     input.TokenSymbol,
		}, nil
	},
)
