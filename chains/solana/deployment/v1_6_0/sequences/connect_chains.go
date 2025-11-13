package sequences

import (
	"encoding/binary"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ConfigureLaneLegAsSource = operations.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.6.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Info("SVM Configuring lane leg as source:", input)
		feeQuoterAddress := solana.PublicKeyFromBytes(input.Source.FeeQuoter)
		offRampAddress := solana.PublicKeyFromBytes(input.Source.OffRamp)
		ccipRouterProgram := solana.PublicKeyFromBytes(input.Source.Router)

		// Add FeeQuoter
		fqOut, err := operations.ExecuteOperation(b, fqops.ConnectChains, chains.SolanaChains()[input.Source.Selector], fqops.ConnectChainsParams{
			FeeQuoter:           feeQuoterAddress,
			OffRamp:             offRampAddress,
			RemoteChainSelector: input.Dest.Selector,
			DestChainConfig:     TranslateFQ(input.Dest.FeeQuoterDestChainConfig),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to add OffRamp to Router: %w", err)
		}
		result.Addresses = append(result.Addresses, fqOut.Output.Addresses...)
		result.BatchOps = append(result.BatchOps, fqOut.Output.BatchOps...)

		// Add Router
		routerOut, err := operations.ExecuteOperation(b, routerops.ConnectChains, chains.SolanaChains()[input.Source.Selector], routerops.ConnectChainsParams{
			Router:              ccipRouterProgram,
			OffRamp:             offRampAddress,
			RemoteChainSelector: input.Dest.Selector,
			AllowlistEnabled:    input.Source.AllowListEnabled,
			AllowedSenders:      TranslateAllowlist(input.Source.AllowList),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize OffRamp: %w", err)
		}
		result.Addresses = append(result.Addresses, routerOut.Output.Addresses...)
		result.BatchOps = append(result.BatchOps, routerOut.Output.BatchOps...)

		return result, nil
	},
)

var ConfigureLaneLegAsDest = operations.NewSequence(
	"ConfigureLaneLegAsDest",
	semver.MustParse("1.6.0"),
	"Configures lane leg as destination on CCIP 1.6.0",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Info("SVM Configuring lane leg as destination:", input)
		offRampAddress := solana.PublicKeyFromBytes(input.Dest.OffRamp)
		ccipRouterProgram := solana.PublicKeyFromBytes(input.Dest.Router)

		// OffRamp must be added to Router before initialization
		routerOut, err := operations.ExecuteOperation(b, routerops.AddOffRamp, chains.SolanaChains()[input.Dest.Selector], routerops.ConnectChainsParams{
			Router:              ccipRouterProgram,
			OffRamp:             offRampAddress,
			RemoteChainSelector: input.Source.Selector,
			AllowlistEnabled:    input.Source.AllowListEnabled,
			AllowedSenders:      TranslateAllowlist(input.Source.AllowList),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to add OffRamp to Router: %w", err)
		}
		result.Addresses = append(result.Addresses, routerOut.Output.Addresses...)
		result.BatchOps = append(result.BatchOps, routerOut.Output.BatchOps...)

		// Add DestChain to OffRamp
		offRampOut, err := operations.ExecuteOperation(b, offrampops.ConnectChains, chains.SolanaChains()[input.Dest.Selector], offrampops.ConnectChainsParams{
			RemoteChainSelector: input.Source.Selector,
			OffRamp:             offRampAddress,
			SourceOnRamp:        input.Source.OffRamp,
			EnabledAsSource:     !input.IsDisabled,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize OffRamp: %w", err)
		}
		result.Addresses = append(result.Addresses, offRampOut.Output.Addresses...)
		result.BatchOps = append(result.BatchOps, offRampOut.Output.BatchOps...)

		return result, nil
	},
)

// high level API
func (a *SolanaAdapter) ConfigureLaneLegAsSource() *cldf_ops.Sequence[ccipapi.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return ConfigureLaneLegAsSource
}

func (a *SolanaAdapter) ConfigureLaneLegAsDest() *cldf_ops.Sequence[ccipapi.UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return ConfigureLaneLegAsDest
}

func TranslateFQ(fqc lanes.FeeQuoterDestChainConfig) fee_quoter.DestChainConfig {
	return fee_quoter.DestChainConfig{
		IsEnabled:                         fqc.IsEnabled,
		MaxNumberOfTokensPerMsg:           fqc.MaxNumberOfTokensPerMsg,
		MaxDataBytes:                      fqc.MaxDataBytes,
		MaxPerMsgGasLimit:                 fqc.MaxPerMsgGasLimit,
		DestGasOverhead:                   fqc.DestGasOverhead,
		DestGasPerPayloadByteBase:         uint32(fqc.DestGasPerPayloadByteBase),
		DestGasPerPayloadByteHigh:         uint32(fqc.DestGasPerPayloadByteHigh),
		DestGasPerPayloadByteThreshold:    uint32(fqc.DestGasPerPayloadByteThreshold),
		DestDataAvailabilityOverheadGas:   fqc.DestDataAvailabilityOverheadGas,
		DestGasPerDataAvailabilityByte:    fqc.DestGasPerDataAvailabilityByte,
		DestDataAvailabilityMultiplierBps: fqc.DestDataAvailabilityMultiplierBps,
		ChainFamilySelector:               [4]byte(binary.BigEndian.AppendUint32(nil, fqc.ChainFamilySelector)),
		EnforceOutOfOrder:                 fqc.EnforceOutOfOrder,
		DefaultTokenFeeUsdcents:           fqc.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead:       fqc.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:                 fqc.DefaultTxGasLimit,
		GasMultiplierWeiPerEth:            fqc.GasMultiplierWeiPerEth,
		GasPriceStalenessThreshold:        fqc.GasPriceStalenessThreshold,
		NetworkFeeUsdcents:                fqc.NetworkFeeUSDCents,
	}
}

func TranslateAllowlist(allowlist []string) []solana.PublicKey {
	var pkList []solana.PublicKey
	for _, addr := range allowlist {
		pkList = append(pkList, solana.MustPublicKeyFromBase58(addr))
	}
	return pkList
}
