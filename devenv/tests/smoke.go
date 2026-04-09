package tests

import (
	"errors"
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

			sendMsgRequireNoError(t, fromImpl, toImpl, msg)
		})

		t.Run(fmt.Sprintf("%s message to eoa", laneTag), func(t *testing.T) {
			if toImpl.Family() == chainsel.FamilyTon {
				t.Skip("This will hang")
			}

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

			sendMsgRequireNoError(t, fromImpl, toImpl, msg)
		})

		t.Run(fmt.Sprintf("%s token transfer", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			fromImplTokenExpansionConfig, err := fromImpl.GetTokenExpansionConfig()
			if errors.Is(err, errors.ErrUnsupported) {
				t.Skip("source chain does not support token transfers")
			}
			require.NoError(t, err)
			srcChainSel, srcTokenCfg := fromImpl.ChainSelector(), fromImplTokenExpansionConfig.DeployTokenInput
			toImplTokenExpansionConfig, err := toImpl.GetTokenExpansionConfig()
			if errors.Is(err, errors.ErrUnsupported) {
				t.Skip("destination chain does not support token transfers")
			}
			require.NoError(t, err)
			dstChainSel, dstTokenCfg := toImpl.ChainSelector(), toImplTokenExpansionConfig.DeployTokenInput

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

			sendMsgRequireNoError(t, fromImpl, toImpl, msg)
			require.Eventually(t, balanceCheck, 5*time.Second, time.Second)
		})

		t.Run(fmt.Sprintf("%s gas limit too high", laneTag), func(t *testing.T) {
			if fromImpl.Family() == chainsel.FamilySolana {
				t.Skip("TODO: evm adapter GetExtraArgs returns nil adapter")
			}
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

			sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
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

			sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
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

			sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
		})

		t.Run(fmt.Sprintf("%s empty extra args tag", laneTag), func(t *testing.T) {
			if toImpl.Family() == chainsel.FamilyTon {
				t.Skip("TODO: Debug why message is committed but not executed. Haven't been able to find the log.")
			}
			if fromImpl.Family() == chainsel.FamilyTon {
				t.Skip("TON expects a well-formatted BOC or BuildMessage will fail")
			}
			if toImpl.Family() == chainsel.FamilyEVM {
				t.Skip("Any->EVM still uses default extraArgs when empty.")
			}

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: toImpl.ChainSelector(),
				Receiver:          toImpl.CCIPReceiver(),
				Data:              []byte("hello world"),
				ExtraArgs:         []byte{}, // empty extraArgs
			})
			require.NoError(t, err)

			sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
		})

		t.Run(fmt.Sprintf("%s invalid/unconfigured chain selector", laneTag), func(t *testing.T) {
			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			const invalidUnconfiguredChainSelector = 1
			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: invalidUnconfiguredChainSelector,
				Receiver:          toImpl.CCIPReceiver(),
				Data:              []byte("hello world"),
				ExtraArgs:         extraArgs,
			})
			require.NoError(t, err)

			_, messageID, err := fromImpl.SendMessage(t.Context(), invalidUnconfiguredChainSelector, msg)
			if err == nil {
				t.Fatalf("expected error when sending message to invalid/unconfigured chain selector, but got success with messageId: %s", messageID)
			}
		})

		t.Run(fmt.Sprintf("%s invalid receiver", laneTag), func(t *testing.T) {
			if fromImpl.Family() == chainsel.FamilySolana {
				t.Skip("GetExtraArgs fails with invalid pubkey receivers, we'd need to construct a raw payload to test against the contract")
			}

			invalidReceivers := toImpl.InvalidAddresses()
			if len(invalidReceivers) == 0 {
				t.Skip("destination chain provided no invalid addresses to test against")
			}

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

				sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
			}
		})

		t.Run(fmt.Sprintf("%s OOO flag is required on non-EVMs", laneTag), func(t *testing.T) {
			if fromImpl.Family() == chainsel.FamilyEVM && toImpl.Family() == chainsel.FamilyEVM {
				t.Skip("EVM->EVM still supports OOO, depending on config")
			}
			if (fromImpl.Family() == chainsel.FamilySolana && toImpl.Family() == chainsel.FamilyEVM) ||
				(fromImpl.Family() == chainsel.FamilyEVM && toImpl.Family() == chainsel.FamilySolana) {
				t.Skip("TODO: Setup lane block OOO on Solana->EVM")
				// 1. evm adapter returns nil adapter
				// 2. solana setup lane seems not to be setting enforeceOOO on the contract side
			}

			if (fromImpl.Family() == chainsel.FamilyTon && toImpl.Family() == chainsel.FamilyEVM) ||
				(fromImpl.Family() == chainsel.FamilyEVM && toImpl.Family() == chainsel.FamilyTon) {
				t.Skip("Skipping OOO test: FeeQuoter v2.0 defaults OOO=true on TON<->EVM lanes")
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

			sendMsgRequireErrorOnSrcChain(t, fromImpl, toImpl, msg)
		})
	}
}

func sendMsgRequireNoError(t *testing.T, fromImpl, toImpl ccip.CCIP16ProductConfiguration, msg any) (uint64, string) {
	block := toImpl.CurrentBlock(t)
	seqNr, messageID, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
	t.Logf("sendMsgRequireNoError got messageID: %s", messageID)
	if err != nil {
		t.Fatalf("failed to send message (id: %s): %v", messageID, err)
	}
	seqNrUint := ccipocr3.SeqNum(seqNr)
	seqNumRange := ccipocr3.NewSeqNumRange(seqNrUint, seqNrUint)
	toImpl.ValidateCommit(t, fromImpl.ChainSelector(), &block, seqNumRange)
	toImpl.ValidateExecSucceeds(t, fromImpl.ChainSelector(), &block, []uint64{seqNr})
	return seqNr, messageID
}

func sendMsgRequireErrorOnSrcChain(t *testing.T, fromImpl, toImpl ccip.CCIP16ProductConfiguration, msg any) {
	t.Helper()
	_, messageID, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
	require.Error(t, err, "sendMsgRequireErrorOnSrcChain got messageID: %s", messageID)
}
