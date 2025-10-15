package fee_quoter

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var ProgramName = "fee_quoter"
var ProgramSize = 5 * 1024 * 1024
var Version *semver.Version = semver.MustParse("1.6.0")

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
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) ([]solana.Instruction, error) {
		programData, err := utils.GetSolProgramData(chain, input.FeeQuoter)
		if err != nil {
			return nil, fmt.Errorf("failed to get program data: %w", err)
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
			return nil, fmt.Errorf("failed to build initialize instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{instruction})
		if err != nil {
			return nil, fmt.Errorf("failed to confirm initialization: %w", err)
		}
		return []solana.Instruction{instruction}, nil
	},
)

var AddPriceUpdater = operations.NewOperation(
	"fee-quoter:add-price-updater",
	Version,
	"Adds a price updater to the FeeQuoter 1.6.0 contract",
	func(b operations.Bundle, chain cldf_solana.Chain, input Params) ([]solana.Instruction, error) {
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(input.FeeQuoter)
		offRampBillingSignerPDA, _, _ := state.FindOfframpBillingSignerPDA(input.OffRamp)
		fqAllowedPriceUpdaterOfframpPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(offRampBillingSignerPDA, input.FeeQuoter)
		instruction, err := fee_quoter.NewAddPriceUpdaterInstruction(
			offRampBillingSignerPDA,
			fqAllowedPriceUpdaterOfframpPDA,
			feeQuoterConfigPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return nil, fmt.Errorf("failed to build add price updater instruction: %w", err)
		}
		return []solana.Instruction{instruction}, nil
	},
)

type Params struct {
	MaxFeeJuelsPerMsg bin.Uint128
	FeeQuoter         solana.PublicKey
	Router            solana.PublicKey
	OffRamp           solana.PublicKey
	LinkToken         solana.PublicKey
	Authority         solana.PublicKey
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
