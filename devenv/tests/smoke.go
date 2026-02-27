package tests

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

var supportedTokenFamilies = map[string]bool{
	chainsel.FamilyEVM:    true,
	chainsel.FamilySolana: true,
}

func RunSmokeTests(t *testing.T, e *deployment.Environment, selectors []uint64) {
	selectorsToImpl := make(map[uint64]ccip.CCIP16ProductConfiguration)
	for _, selector := range selectors {
		family, err := chainsel.GetSelectorFamily(selector)
		require.NoError(t, err)
		chainID, err := chainsel.GetChainIDFromSelector(selector)
		require.NoError(t, err)
		i, err := ccip.NewCCIPImplFromNetwork(family, chainID)
		require.NoError(t, err)
		i.SetCLDF(e)
		selectorsToImpl[selector] = i
	}

	if os.Getenv("PARALLEL_E2E_TESTS") == "true" {
		t.Parallel()
	}

	type testpair struct {
		fromChain ccip.CCIP16ProductConfiguration
		toChain   ccip.CCIP16ProductConfiguration
	}
	matrix := []testpair{}
	for _, i := range selectors {
		for _, j := range selectors {
			if i == j {
				continue
			}
			matrix = append(matrix, testpair{
				fromChain: selectorsToImpl[i],
				toChain:   selectorsToImpl[j],
			})
		}
	}

	for _, tc := range matrix {
		fromImpl := tc.fromChain
		toImpl := tc.toChain
		laneTag := fmt.Sprintf("%s->%s", fromImpl.Family(), toImpl.Family())

		t.Run(fmt.Sprintf("%s message to contract", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello contract"),
				FeeToken:          "",
				ExtraArgs:         extraArgs,
				TokenAmounts:      nil,
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.NoError(t, err)
			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
			toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})
		})

		t.Run(fmt.Sprintf("%s message to eoa", laneTag), func(t *testing.T) {
			receiver := toImpl.EOAReceiver(t)
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello eoa"),
				FeeToken:          "",
				ExtraArgs:         extraArgs,
				TokenAmounts:      nil,
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.NoError(t, err)
			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
			toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})
		})

		t.Run(fmt.Sprintf("%s token transfer", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			_, fromSupported := supportedTokenFamilies[fromImpl.Family()]
			_, toSupported := supportedTokenFamilies[toImpl.Family()]
			if !fromSupported || !toSupported {
				t.Skip("Token transfers not supported on " + laneTag)
			}

			srcChainSel, srcTokenCfg := fromImpl.ChainSelector(), fromImpl.GetTokenExpansionConfig().DeployTokenInput
			dstChainSel, dstTokenCfg := toImpl.ChainSelector(), toImpl.GetTokenExpansionConfig().DeployTokenInput

			srcTokenFilterDS := datastore.AddressRef{ChainSelector: srcChainSel, Qualifier: srcTokenCfg.Symbol, Type: datastore.ContractType(srcTokenCfg.Type)}
			srcTokenRef, err := datastore_utils.FindAndFormatRef(e.DataStore, srcTokenFilterDS, srcChainSel, datastore_utils.FullRef)
			require.NoError(t, err)

			dstTokenFilterDS := datastore.AddressRef{ChainSelector: dstChainSel, Qualifier: dstTokenCfg.Symbol, Type: datastore.ContractType(dstTokenCfg.Type)}
			dstTokenRef, err := datastore_utils.FindAndFormatRef(e.DataStore, dstTokenFilterDS, dstChainSel, datastore_utils.FullRef)
			require.NoError(t, err)

			// Here, we avoid using a fractional token amount to simplify the test logic. In
			// this case, we transfer 10^src_decimals units on the *src* chain, which is the
			// the equivalent of one whole token on the source. This results in the receiver
			// getting the equivalent of 10^dst_decimals units which is also one whole token
			// on the *destination* chain. If we want to test fractional amounts later, then
			// we'd need to scale the amounts according to both the src/dst token decimals.
			sendAmnt := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(srcTokenCfg.Decimals)), nil)

			// We expect the receiver to get 1 whole token on the destination chain.
			recvAmnt := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(dstTokenCfg.Decimals)), nil)

			// Query the initial balance of the receiver account on the destination chain
			initAmnt, err := toImpl.GetTokenBalance(t.Context(), dstTokenRef.Address, receiver)
			require.NoError(t, err)

			// Calculate the total balance that the receiver should have after execution
			trgtAmnt := new(big.Int).Add(initAmnt, recvAmnt)

			// This balance check function will be polled at regular intervals. It returns
			// true when the receiver's current balance matches the expected target amount
			balanceCheck := func() bool {
				t.Helper()

				balance, err := toImpl.GetTokenBalance(t.Context(), dstTokenRef.Address, receiver)
				require.NoError(t, err)

				t.Log(fmt.Sprintf("Fetched receiver token balance on chain %d (%s)", toImpl.ChainSelector(), toImpl.Family()),
					"token.qualifier="+dstTokenRef.Qualifier,
					"token.address="+dstTokenRef.Address,
					"token.type="+dstTokenRef.Type,
					"balance.target="+trgtAmnt.String(),
					"balance.actual="+balance.String(),
				)
				return balance.Cmp(trgtAmnt) == 0
			}

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello eoa"),
				FeeToken:          "",
				ExtraArgs:         extraArgs,
				TokenAmounts: []testadapters.TokenAmount{
					{Amount: sendAmnt, Token: srcTokenRef.Address},
				},
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.NoError(t, err)
			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
			toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})
			require.Eventually(t, balanceCheck, 5*time.Second, time.Second)
		})

		t.Run(fmt.Sprintf("%s gas limit too high", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family(), testadapters.NewGasLimitExtraArg(big.NewInt(math.MaxUint32)))
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s payload larger than limit", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			// Construct a payload that exceeds the typical 32KB limit
			oversizedData := make([]byte, 33*1024) // 33 KB of data

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              oversizedData,
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s invalid extra args tag", laneTag), func(t *testing.T) {
			if fromImpl.Family() == chainsel.FamilyTon {
				t.Skip("TON expects a well-formatted BOC or BuildMessage will fail")
			}

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          toImpl.CCIPReceiver(),
				Data:              []byte("hello world"),
				ExtraArgs:         []byte{1, 2, 3, 4, 99, 99}, // invalid extraArgs prefix
			})
			require.NoError(t, err)

			_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s invalid/unconfigured chain selector", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: 1, // invalid/unconfigured chain selector
				Receiver:          toImpl.CCIPReceiver(),
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s invalid receiver", laneTag), func(t *testing.T) {
			if fromImpl.Family() == chainsel.FamilySolana {
				t.Skip("GetExtraArgs fails with invalid pubkey receivers, we'd need to construct a raw payload to test against the contract")
			}

			invalidReceivers := toImpl.InvalidCCIPReceivers()

			for _, invalidReceiver := range invalidReceivers {

				extraArgs, err := toImpl.GetExtraArgs(invalidReceiver, fromImpl.Family(), testadapters.NewGasLimitExtraArg(big.NewInt(math.MaxInt64)))
				require.NoError(t, err)

				msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
					DestChainSelector: toImpl.ChainSelector(),
					Receiver:          invalidReceiver,
					Data:              []byte("hello world"),
					ExtraArgs:         extraArgs,
				})
				require.NoError(t, err)

				_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
				require.Error(t, err)
			}
		})

		// WIP: message with not enough gas
		// then manual re-exec with higher limit
		t.Run(fmt.Sprintf("%s not enough gas; manual re-exec", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family(), testadapters.NewGasLimitExtraArg(big.NewInt(1)))
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)

			// I am not sure if we can grab seq when SendMessage fails.

			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
			toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})

			// Assuming the error is due to not enough gas, we can attempt a manual re-execution with a higher gas limit
			// Note: The actual implementation of manual re-execution would depend on the capabilities of the test adapters and the underlying chains
			// For this example, we'll just log the intention to re-execute with a higher gas limit
			t.Log("Initial execution failed due to insufficient gas, attempting manual re-execution with higher gas limit for message", seq, seqNr)

			// TODO set a require.Eventually() here to poll for the message to be re-executed with the higher gas limit and succeed?
		})

		// WIP: receiver fails
		// then manual re-exec with higher limit
		t.Run(fmt.Sprintf("%s receiver fails; manual re-exec", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()

			err := toImpl.SetReceiverRejectAll(t.Context(), true)
			require.NoError(t, err)
			t.Cleanup(func() {
				err := toImpl.SetReceiverRejectAll(t.Context(), false)
				require.NoError(t, err)
			})

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)

			// I am not sure if we can grab seq when SendMessage fails.

			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
			toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})

			err = toImpl.SetReceiverRejectAll(t.Context(), false)
			require.NoError(t, err)

			// Assuming the error is due to not enough gas, we can attempt a manual re-execution with a higher gas limit
			// Note: The actual implementation of manual re-execution would depend on the capabilities of the test adapters and the underlying chains
			// For this example, we'll just log the intention to re-execute with a higher gas limit
			t.Log("Initial execution failed due to logical error, attempting manual re-execution after fixing error", seq, seqNr)

			// TODO set a require.Eventually() here to poll for the message to be re-executed with the higher gas limit and succeed?
		})

		t.Run(fmt.Sprintf("%s allowlist enabled", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()

			err := fromImpl.SetAllowlist(t.Context(), toImpl.ChainSelector(), true)
			require.NoError(t, err)
			t.Cleanup(func() {
				err := fromImpl.SetAllowlist(t.Context(), toImpl.ChainSelector(), false)
				require.NoError(t, err)
			})

			t.Run("sender not allowed", func(t *testing.T) {
				err := fromImpl.UpdateSenderAllowlistStatus(t.Context(), toImpl.ChainSelector(), false)
				require.NoError(t, err)
				t.Cleanup(func() {
					err := fromImpl.UpdateSenderAllowlistStatus(t.Context(), toImpl.ChainSelector(), true)
					require.NoError(t, err)
				})

				extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
				require.NoError(t, err)

				msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
					DestChainSelector: toImpl.ChainSelector(),
					Receiver:          receiver,
					Data:              []byte("hello world"),
					ExtraArgs:         extraArgs,
				})
				require.NoError(t, err)

				_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
				require.Error(t, err)
			})

			t.Run("sender allowed", func(t *testing.T) {
				err := fromImpl.UpdateSenderAllowlistStatus(t.Context(), toImpl.ChainSelector(), true)
				require.NoError(t, err)
				t.Cleanup(func() {
					err := fromImpl.UpdateSenderAllowlistStatus(t.Context(), toImpl.ChainSelector(), false)
					require.NoError(t, err)
				})

				extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
				require.NoError(t, err)

				msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
					DestChainSelector: toImpl.ChainSelector(),
					Receiver:          receiver,
					Data:              []byte("hello world"),
					ExtraArgs:         extraArgs,
				})
				require.NoError(t, err)

				seq, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
				require.NoError(t, err)
				seqNr := ccipocr3.SeqNum(seq)
				seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
				toImpl.ValidateCommit(t, fromImpl.ChainSelector(), nil, seqNumRange)
				toImpl.ValidateExec(t, fromImpl.ChainSelector(), nil, []uint64{seq})
			})

		})

		t.Run(fmt.Sprintf("%s RMN - arbitrary message from a cursed source chain", laneTag), func(t *testing.T) {
			err := toImpl.RMNSetSrcChainCursed(t.Context(), fromImpl.ChainSelector(), true)
			require.NoError(t, err)
			t.Cleanup(func() {
				err := toImpl.RMNSetSrcChainCursed(t.Context(), fromImpl.ChainSelector(), false)
				require.NoError(t, err)
			})

			receiver := toImpl.CCIPReceiver()

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			_, err = fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s RMN - arbitrary message to a cursed chain", laneTag), func(t *testing.T) {
			err := toImpl.RMNSetDestChainCursed(t.Context(), toImpl.ChainSelector(), true)
			require.NoError(t, err)
			t.Cleanup(func() {
				err := toImpl.RMNSetDestChainCursed(t.Context(), toImpl.ChainSelector(), false)
				require.NoError(t, err)
			})

			receiver := toImpl.CCIPReceiver()

			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			_, err = fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.Error(t, err)
		})

		t.Run(fmt.Sprintf("%s OOO flag is required on non-EVMs", laneTag), func(t *testing.T) {
			if fromImpl.Family() != chainsel.FamilyEVM || toImpl.Family() != chainsel.FamilyEVM {
				t.Skip("EVM->EVM still supports OOO, depending on config")
			}

			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family(), testadapters.NewOutOfOrderExtraArg(false))
			require.NoError(t, err)

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          receiver,
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			_, err = fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
			require.Error(t, err)
		})
	}
}
