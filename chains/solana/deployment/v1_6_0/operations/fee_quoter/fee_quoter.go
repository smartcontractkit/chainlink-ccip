package fee_quoter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var ProgramName = "fee_quoter"
var ProgramSize = 5 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

type ConnectChainsParams struct {
	FeeQuoter           solana.PublicKey
	OffRamp             solana.PublicKey
	RemoteChainSelector uint64
	DestChainConfig     fee_quoter.DestChainConfig
}

var Deploy = operations.NewOperation(
	"fee-quoter:deploy",
	Version,
	"Deploys the FeeQuoter program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			ContractType,
			Version,
			"",
			ProgramName,
			ProgramSize)
	},
)

var Initialize = operations.NewOperation(
	"fee-quoter:initialize",
	Version,
	"Initializes the FeeQuoter 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		programData, err := utils.GetSolProgramData(chain, input.FeeQuoter)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get program data: %w", err)
		}
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(input.FeeQuoter)
		instruction, err := fee_quoter.NewInitializeInstruction(
			input.MaxFeeJuelsPerMsg,
			input.Router,
			feeQuoterConfigPDA,
			input.LinkToken,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
			input.FeeQuoter,
			programData.Address,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm initialization: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var AddPriceUpdater = operations.NewOperation(
	"fee-quoter:add-price-updater",
	Version,
	"Adds a price updater to the FeeQuoter 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) (sequences.OnChainOutput, error) {
		authority := GetAuthority(chain, input.FeeQuoter)
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(input.FeeQuoter)
		offRampBillingSignerPDA, _, _ := state.FindOfframpBillingSignerPDA(input.OffRamp)
		fqAllowedPriceUpdaterOfframpPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(offRampBillingSignerPDA, input.FeeQuoter)
		instruction, err := fee_quoter.NewAddPriceUpdaterInstruction(
			offRampBillingSignerPDA,
			fqAllowedPriceUpdaterOfframpPDA,
			feeQuoterConfigPDA,
			authority,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build add price updater instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

var ConnectChains = operations.NewOperation(
	"fee-quoter:connect-chains",
	Version,
	"Connects the FeeQuoter 1.6.0 contract to other chains",
	func(b operations.Bundle, chain cldf_solana.Chain, input ConnectChainsParams) (sequences.OnChainOutput, error) {
		isUpdate := false
		authority := GetAuthority(chain, input.FeeQuoter)
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(input.FeeQuoter)
		fqRemoteChainPDA, _, _ := state.FindFqDestChainPDA(input.RemoteChainSelector, input.FeeQuoter)
		var destChainStateAccount fee_quoter.DestChain
		err := chain.GetAccountDataBorshInto(context.Background(), fqRemoteChainPDA, &destChainStateAccount)
		if err == nil {
			fmt.Println("Remote chain state account found:", destChainStateAccount)
			isUpdate = true
		}
		var ixn solana.Instruction
		if isUpdate {
			ixn, err = fee_quoter.NewUpdateDestChainConfigInstruction(
				input.RemoteChainSelector,
				input.DestChainConfig,
				feeQuoterConfigPDA,
				fqRemoteChainPDA,
				chain.DeployerKey.PublicKey(),
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build update dest chain instruction: %w", err)
			}
		} else {
			ixn, err = fee_quoter.NewAddDestChainInstruction(
				input.RemoteChainSelector,
				input.DestChainConfig,
				feeQuoterConfigPDA,
				fqRemoteChainPDA,
				authority,
				solana.SystemProgramID,
			).ValidateAndBuild()
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to build add dest chain instruction: %w", err)
			}
			err = utils.ExtendLookupTable(chain, input.OffRamp, []solana.PublicKey{fqRemoteChainPDA})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to extend OffRamp lookup table: %w", err)
			}
		}
		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to confirm add price updater: %w", err)
		}
		return sequences.OnChainOutput{}, nil
	},
)

func GetAuthority(chain cldf_solana.Chain, program solana.PublicKey) solana.PublicKey {
	return chain.DeployerKey.PublicKey()
}

type Params struct {
	MaxFeeJuelsPerMsg bin.Uint128
	FeeQuoter         solana.PublicKey
	Router            solana.PublicKey
	OffRamp           solana.PublicKey
	LinkToken         solana.PublicKey
}

func DefaultParams() Params {
	defaultLow, defaultHigh := GetHighLowBits(big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)))
	return Params{
		MaxFeeJuelsPerMsg: bin.Uint128{
			Lo:         defaultLow,
			Hi:         defaultHigh,
			Endianness: nil,
		},
	}
}

func GetHighLowBits(n *big.Int) (low, high uint64) {
	mask := big.NewInt(0).SetUint64(0xFFFFFFFFFFFFFFFF) // 64-bit mask

	lowBig := big.NewInt(0).And(n, mask)
	low = lowBig.Uint64()

	highBig := big.NewInt(0).Rsh(n, 64) // Shift right by 64 bits
	high = highBig.Uint64()

	return low, high
}
