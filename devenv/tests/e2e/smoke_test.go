package e2e

import (
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

func TestE2ESmoke(t *testing.T) {
	in, err := ccip.LoadOutput[ccip.Cfg]("../../env-out.toml")
	require.NoError(t, err)
	if in.ForkedEnvConfig != nil {
		t.Skip("Skipping E2E tests on forked environments, not supported yet")
	}
	chainIDs, wsURLs := make([]string, 0), make([]string, 0)
	for _, bc := range in.Blockchains {
		chainIDs = append(chainIDs, bc.ChainID)
		wsURLs = append(wsURLs, bc.Out.Nodes[0].ExternalWSUrl)
	}

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)
	selectorsToImpl := make(map[uint64]ccip.CCIP16ProductConfiguration)

	for _, bc := range in.Blockchains {
		i, err := ccip.NewCCIPImplFromNetwork(bc.Type, bc.ChainID)
		require.NoError(t, err)
		i.SetCLDF(e)
		family, err := chainsel.GetSelectorFamily(i.ChainSelector())
		require.NoError(t, err)
		networkInfo, err := chainsel.GetChainDetailsByChainIDAndFamily(bc.ChainID, family)
		require.NoError(t, err)
		selectorsToImpl[networkInfo.ChainSelector] = i
	}

	t.Cleanup(func() {
		_, err := framework.SaveContainerLogs(fmt.Sprintf("%s-%s", framework.DefaultCTFLogsDir, t.Name()))
		require.NoError(t, err)
	})

	type testcase struct {
		name         string
		fromSelector uint64
		toSelector   uint64
	}
	tcs := []testcase{}
	for i := range selectors {
		for j := range selectors {
			if i == j {
				continue
			}
			fromFamily, _ := chainsel.GetSelectorFamily(selectors[i])
			toFamily, _ := chainsel.GetSelectorFamily(selectors[j])
			tcs = append(tcs, testcase{
				name:         fmt.Sprintf("msg execution eoa receiver from %s to %s", fromFamily, toFamily),
				fromSelector: selectors[i],
				toSelector:   selectors[j],
			})
		}
	}

	for _, tc := range tcs {
		fromImpl := selectorsToImpl[tc.fromSelector]
		toImpl := selectorsToImpl[tc.toSelector]
		supportedTokenFamilies := map[string]bool{
			chainsel.FamilyEVM:    true,
			chainsel.FamilySolana: true,
		}
		_, fromSupported := supportedTokenFamilies[fromImpl.Family()]
		_, toSupported := supportedTokenFamilies[toImpl.Family()]
		if fromSupported && toSupported {
			tc.name += " with token transfer"
		} else {
			tc.name += " without token transfer"
		}
		// Capture the loop variable so each goroutine gets its own copy.
		t.Run(tc.name, func(t *testing.T) {
			if os.Getenv("PARALLEL_E2E_TESTS") == "true" {
				t.Parallel()
			}

			t.Logf("Testing CCIP message from chain %d to chain %d", tc.fromSelector, tc.toSelector)

			receiver := toImpl.CCIPReceiver()
			extraArgs, err := toImpl.GetExtraArgs(receiver, fromImpl.Family())
			require.NoError(t, err)

			// TODO: once non-EVM tooling supports token transfers, we'll be able
			// to remove the EVM <-> EVM filter and directly set these variables.
			var tokenAmounts []testadapters.TokenAmount = nil
			var balanceCheck func() bool = nil
			// technically something like solana <> solana isn't valid, but this
			// check is just to ensure we only run token transfer tests on supported
			// chain families for now.
			if fromSupported && toSupported {
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
				balanceCheck = func() bool {
					t.Helper()

					balance, err := toImpl.GetTokenBalance(t.Context(), dstTokenRef.Address, receiver)
					require.NoError(t, err)

					t.Log(fmt.Sprintf("Fetched receiver token balance on chain %d (%s)", tc.toSelector, toImpl.Family()),
						"token.qualifier="+dstTokenRef.Qualifier,
						"token.address="+dstTokenRef.Address,
						"token.type="+dstTokenRef.Type,
						"balance.target="+trgtAmnt.String(),
						"balance.actual="+balance.String(),
					)

					return balance.Cmp(trgtAmnt) == 0
				}

				tokenAmounts = []testadapters.TokenAmount{
					{Amount: sendAmnt, Token: srcTokenRef.Address},
				}
			}

			msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
				DestChainSelector: tc.toSelector,
				Receiver:          receiver,
				Data:              []byte("hello eoa"),
				FeeToken:          "",
				ExtraArgs:         extraArgs,
				TokenAmounts:      tokenAmounts,
			})
			require.NoError(t, err)

			seq, err := fromImpl.SendMessage(t.Context(), tc.toSelector, msg)
			require.NoError(t, err)
			seqNr := ccipocr3.SeqNum(seq)
			seqNumRange := ccipocr3.NewSeqNumRange(seqNr, seqNr)
			toImpl.ValidateCommit(t, tc.fromSelector, nil, seqNumRange)
			toImpl.ValidateExec(t, tc.fromSelector, nil, []uint64{seq})

			// TODO: once non-EVM tooling supports token transfers we can
			// remove this if statement and always run the balance check.
			if balanceCheck != nil {
				require.Eventually(t, balanceCheck, 5*time.Second, time.Second)
			}
		})
	}
}
