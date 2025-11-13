package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	rmnremoteops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	tokensops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) DeployChainContracts() *cldf_ops.Sequence[deployops.ContractDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployChainContracts
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.6.0"),
	"Deploys all required contracts for CCIP 1.6.0 to a Solana chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.ContractDeploymentConfigPerChainWithAddress) (output sequences.OnChainOutput, err error) {
		chain := chains.SolanaChains()[input.ChainSelector]
		addresses := make([]datastore.AddressRef, 0)

		// Deploy LINK
		linkRef, err := operations.ExecuteOperation(b, tokensops.DeployLINK, chain, tokensops.Params{
			TokenPrivKey:  solana.MustPrivateKeyFromBase58(input.TokenPrivKey),
			TokenDecimals: input.TokenDecimals,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LINK: %w", err)
		}
		addresses = append(addresses, linkRef.Output)

		// Deploy Router
		routerRef, err := operations.ExecuteOperation(b, routerops.Deploy, chain, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Router: %w", err)
		}
		addresses = append(addresses, routerRef.Output)

		// Deploy FeeQuoter
		feeQuoterRef, err := operations.ExecuteOperation(b, fqops.Deploy, chain, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy FeeQuoter: %w", err)
		}
		addresses = append(addresses, feeQuoterRef.Output)

		// Deploy OffRamp
		offRampRef, err := operations.ExecuteOperation(b, offrampops.Deploy, chain, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef.Output)

		// Deploy RMN Remote
		rmnRemoteRef, err := operations.ExecuteOperation(b, rmnremoteops.Deploy, chain, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy RMN Remote: %w", err)
		}
		addresses = append(addresses, rmnRemoteRef.Output)
		linkTokenAddress := solana.MustPublicKeyFromBase58(linkRef.Output.Address)
		feeQuoterAddress := solana.MustPublicKeyFromBase58(feeQuoterRef.Output.Address)
		offRampAddress := solana.MustPublicKeyFromBase58(offRampRef.Output.Address)
		rmnRemoteAddress := solana.MustPublicKeyFromBase58(rmnRemoteRef.Output.Address)
		ccipRouterProgram := solana.MustPublicKeyFromBase58(routerRef.Output.Address)

		// Initialize FeeQuoter
		lowBits, highBits := GetHighLowBits(input.MaxFeeJuelsPerMsg)
		_, err = operations.ExecuteOperation(b, fqops.Initialize, chain, fqops.Params{
			MaxFeeJuelsPerMsg: bin.Uint128{
				Lo:         lowBits,
				Hi:         highBits,
				Endianness: nil,
			},
			FeeQuoter: feeQuoterAddress,
			Router:    ccipRouterProgram,
			OffRamp:   offRampAddress,
			LinkToken: linkTokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize FeeQuoter: %w", err)
		}

		_, err = operations.ExecuteOperation(b, fqops.AddPriceUpdater, chain, fqops.Params{
			OffRamp:   offRampAddress,
			FeeQuoter: feeQuoterAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to add LINK billing token to FeeQuoter: %w", err)
		}

		// Initialize Router
		_, err = operations.ExecuteOperation(b, routerops.Initialize, chain, routerops.Params{
			FeeQuoter: feeQuoterAddress,
			Router:    ccipRouterProgram,
			LinkToken: linkTokenAddress,
			RMNRemote: rmnRemoteAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize Router: %w", err)
		}

		// Initialize OffRamp
		_, err = operations.ExecuteOperation(b, offrampops.Initialize, chain, offrampops.Params{
			FeeQuoter: feeQuoterAddress,
			OffRamp:   offRampAddress,
			Router:    ccipRouterProgram,
			RMNRemote: rmnRemoteAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize OffRamp: %w", err)
		}

		_, err = operations.ExecuteOperation(b, offrampops.InitializeConfig, chain, offrampops.Params{
			EnableExecutionAfter: int64(input.PermissionLessExecutionThresholdSeconds),
			OffRamp:              offRampAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize OffRamp config: %w", err)
		}

		// Initialize RMN Remote
		_, err = operations.ExecuteOperation(b, rmnremoteops.Initialize, chain, rmnremoteops.Params{
			RMNRemote: solana.MustPublicKeyFromBase58(rmnRemoteRef.Output.Address),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize RMN Remote: %w", err)
		}

		// LOOKUP TABLE

		// off ramp
		offRampReferenceAddressesPDA, _, _ := state.FindOfframpReferenceAddressesPDA(offRampAddress)
		offRampBillingSignerPDA, _, _ := state.FindOfframpBillingSignerPDA(offRampAddress)
		offRampConfigPDA, _, _ := state.FindOfframpConfigPDA(offRampAddress)
		// fee quoter
		linkFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(linkTokenAddress, feeQuoterAddress)
		wsolFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(solana.WrappedSol, feeQuoterAddress)
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(feeQuoterAddress)
		// router
		feeBillingSignerPDA, _, _ := state.FindFeeBillingSignerPDA(ccipRouterProgram)
		routerConfigPDA, _, _ := state.FindConfigPDA(ccipRouterProgram)
		// rmn remote
		rmnRemoteCursePDA, _, _ := state.FindRMNRemoteCursesPDA(rmnRemoteAddress)
		rmnRemoteConfigPDA, _, _ := state.FindRMNRemoteConfigPDA(rmnRemoteAddress)
		lookupTableKeys := []solana.PublicKey{
			// offramp
			offRampAddress,
			offRampConfigPDA,
			offRampReferenceAddressesPDA,
			offRampBillingSignerPDA,
			// fee quoter
			feeQuoterConfigPDA,
			feeQuoterAddress,
			linkFqBillingConfigPDA,
			wsolFqBillingConfigPDA,
			// router
			ccipRouterProgram,
			routerConfigPDA,
			feeBillingSignerPDA,
			// rmn remote
			rmnRemoteAddress,
			rmnRemoteConfigPDA,
			rmnRemoteCursePDA,
		}

		err = utils.ExtendLookupTable(chain, offRampAddress, lookupTableKeys)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to extend OffRamp lookup table: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
		}, nil
	},
)

func GetHighLowBits(n *big.Int) (low, high uint64) {
	mask := big.NewInt(0).SetUint64(0xFFFFFFFFFFFFFFFF) // 64-bit mask

	lowBig := big.NewInt(0).And(n, mask)
	low = lowBig.Uint64()

	highBig := big.NewInt(0).Rsh(n, 64) // Shift right by 64 bits
	high = highBig.Uint64()

	return low, high
}
