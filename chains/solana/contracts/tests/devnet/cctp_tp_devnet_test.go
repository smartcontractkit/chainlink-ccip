//go:build devnet
// +build devnet

package contracts

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/base_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

func TestCctpTpDevnet(t *testing.T) {
	devnetInfo, err := getDevnetInfo()
	require.NoError(t, err)

	ctx := tests.Context(t)
	client := rpc.New(devnetInfo.RPC)

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())
	deployer := solana.PrivateKey(devnetInfo.PrivateKeys.Deployer)
	require.True(t, deployer.IsValid())

	offrampAddress, err := solana.PublicKeyFromBase58(devnetInfo.Offramp)
	require.NoError(t, err)

	offrampPDAs, err := getOfframpPDAs(offrampAddress)
	require.NoError(t, err)

	var referenceAddresses ccip_offramp.ReferenceAddresses
	t.Run("Read Reference Addresses", func(t *testing.T) {
		require.NoError(t, common.GetAccountDataBorshInto(ctx, client, offrampPDAs.referenceAddresses, rpc.CommitmentConfirmed, &referenceAddresses))
		fmt.Printf("Reference Addresses: %+v\n", referenceAddresses)
	})

	ccip_router.SetProgramID(referenceAddresses.Router)
	fee_quoter.SetProgramID(referenceAddresses.FeeQuoter)

	cctpTpProgram := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.TokenPool)
	cctpMtProgram := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.MessageTransmitter)
	cctpTmmProgram := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.TokenMessengerMinter)
	cctp_token_pool.SetProgramID(cctpTpProgram)

	usdcMint := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.UsdcMint)
	usdcDecimals := uint8(6)

	domains := map[uint64]uint32{
		devnetInfo.ChainSelectors.Sepolia: 0,
	}

	usdcEvm := "1c7D4B196Cb0C7B01d743Fbc6116a902379C7238"
	usdcEvmPartialBytes, err := hex.DecodeString(usdcEvm)
	require.NoError(t, err)
	usdcEvmBytes := [32]byte{}
	copy(usdcEvmBytes[32-len(usdcEvmPartialBytes):], usdcEvmPartialBytes)

	receiverAddress := devnetInfo.CCTP.Sepolia.ReceiverAddress
	receiverAddrBytes, err := hex.DecodeString(receiverAddress)
	require.NoError(t, err)
	fullReceiverAddress := [32]byte{}
	copy(fullReceiverAddress[32-len(receiverAddrBytes):], receiverAddrBytes)

	remotePoolBytes, err := hex.DecodeString(devnetInfo.CCTP.Sepolia.TokenPool)
	require.NoError(t, err)

	chainSelector := devnetInfo.ChainSelectors.Sepolia
	domain := domains[chainSelector]

	domainDestCallerAddress, err := hex.DecodeString(devnetInfo.CCTP.Sepolia.AllowedCaller)
	require.NoError(t, err)
	domainDestCallerBytes := [32]byte{}
	copy(domainDestCallerBytes[32-len(domainDestCallerAddress):], domainDestCallerAddress)
	domainDestCaller := solana.PublicKey(domainDestCallerBytes)

	cctpPool := getCctpTokenPoolPDAs(t, cctpTpProgram, chainSelector, usdcMint)
	messageTransmitter := getMessageTransmitterPDAs(t, cctpMtProgram, cctpTmmProgram)
	tokenMessengerMinter := getTokenMessengerMinterPDAs(t, cctpTmmProgram, domain, usdcMint)

	var tpLookupTableAddr solana.PublicKey
	var tpLookupTable map[solana.PublicKey]solana.PublicKeySlice

	t.Run("TypeVersion", func(t *testing.T) {
		t.Skip()

		ix, err := cctp_token_pool.NewTypeVersionInstruction().ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		require.NotNil(t, result)

		output, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, cctpPool.program.String(), func(b []byte) string {
			require.Len(t, b, int(binary.LittleEndian.Uint32(b[:4]))+4) // the first 4 bytes just encodes the length
			return string(b[4:])
		})
		require.NoError(t, err)
		// regex from https://semver.org/#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
		semverRegex := "(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(?:-((?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+([0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?"
		require.Regexp(t, fmt.Sprintf("^%s %s$", "cctp-token-pool", semverRegex), output)
		fmt.Printf("Type Version: %s\n", output)
	})

	routerConfig, _, err := state.FindConfigPDA(referenceAddresses.Router)
	require.NoError(t, err)
	tokenAdminRegistry, _, err := state.FindTokenAdminRegistryPDA(usdcMint, referenceAddresses.Router)
	require.NoError(t, err)

	type ProgramData struct {
		DataType uint32
		Address  solana.PublicKey
	}

	// get program data account
	data, err := client.GetAccountInfoWithOpts(ctx, cctpPool.program, &rpc.GetAccountInfoOpts{
		Commitment: config.DefaultCommitment,
	})
	require.NoError(t, err)
	// Decode program data
	var programData ProgramData
	require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

	t.Run("Delete chain config", func(t *testing.T) {
		t.Skip()

		ix, err := cctp_token_pool.NewDeleteChainConfigInstruction(
			chainSelector,
			usdcMint,
			cctpPool.state,
			cctpPool.chainConfig,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)

		res := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		require.NotNil(t, res)

		for _, log := range res.Meta.LogMessages {
			fmt.Println(log)
		}
	})

	t.Run("Initialize TokenPool", func(t *testing.T) {
		t.Skip()

		// poolInitI, err := cctp_token_pool.NewInitializeInstruction(
		// 	referenceAddresses.Router,
		// 	referenceAddresses.RmnRemote,
		// 	cctpPool.state,
		// 	usdcMint,
		// 	admin.PublicKey(),
		// 	solana.SystemProgramID,
		// 	cctpPool.program,
		// 	programData.Address,
		// ).ValidateAndBuild()
		// require.NoError(t, err)

		// poolEditI, err := cctp_token_pool.NewSetRmnRemoteInstruction(
		// 	referenceAddresses.RmnRemote,
		// 	cctpPool.state,
		// 	usdcMint,
		// 	admin.PublicKey(),
		// ).ValidateAndBuild()
		// require.NoError(t, err)

		// // set pool config
		ixConfigure, err := cctp_token_pool.NewInitChainRemoteConfigInstruction(
			chainSelector,
			usdcMint,
			cctp_token_pool.RemoteConfig{
				TokenAddress: cctp_token_pool.RemoteAddress{
					Address: usdcMint.Bytes(),
				},
				Decimals:      usdcDecimals,
				PoolAddresses: []cctp_token_pool.RemoteAddress{},
			},
			cctpPool.state,
			cctpPool.chainConfig,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		ixEditChainRemoteConfig, err := cctp_token_pool.NewEditChainRemoteConfigInstruction(
			chainSelector,
			usdcMint,
			cctp_token_pool.RemoteConfig{
				PoolAddresses: []base_token_pool.RemoteAddress{},
				TokenAddress: base_token_pool.RemoteAddress{
					Address: usdcEvmBytes[:],
				},
				Decimals: usdcDecimals,
			},
			cctpPool.state,
			cctpPool.chainConfig,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		ixCctpConfigure, err := cctp_token_pool.NewEditChainRemoteConfigCctpInstruction(
			chainSelector,
			usdcMint,
			cctp_token_pool.CctpChain{
				DomainId:          domain,
				DestinationCaller: domainDestCaller,
			},
			cctpPool.state,
			cctpPool.chainConfig,
			admin.PublicKey(),
		).ValidateAndBuild()
		require.NoError(t, err)

		ixAppend, err := cctp_token_pool.NewAppendRemotePoolAddressesInstruction(
			chainSelector,
			usdcMint,
			[]cctp_token_pool.RemoteAddress{{Address: remotePoolBytes}},
			cctpPool.state,
			cctpPool.chainConfig,
			admin.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		// create pool token account
		// createP, poolTokenAccount, err := tokens.CreateAssociatedTokenAccount(solana.TokenProgramID, usdcMint, cctpPool.signer, admin.PublicKey())
		// require.NoError(t, err)
		// require.Equal(t, poolTokenAccount, cctpPool.tokenAccount)

		// submit tx with all instructions
		res := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ixConfigure, ixEditChainRemoteConfig, ixAppend, ixCctpConfigure}, admin, config.DefaultCommitment)
		require.NotNil(t, res)

		// 	// validate state
		// 	var configAccount cctp_token_pool.State
		// 	require.NoError(t, common.GetAccountDataBorshInto(ctx, client, cctpPool.state, config.DefaultCommitment, &configAccount))
		// 	require.Equal(t, cctpPool.tokenAccount, configAccount.Config.PoolTokenAccount)

		// 	// validate events
		// 	var eventConfigured tokens.EventChainConfigured
		// 	require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainConfigured", &eventConfigured, config.PrintEvents))
		// 	require.Equal(t, chainSelector, eventConfigured.ChainSelector)
		// 	require.Equal(t, 0, len(eventConfigured.PoolAddresses))
		// 	require.Equal(t, 0, len(eventConfigured.PreviousPoolAddresses))
		// 	require.Equal(t, cctp_token_pool.RemoteAddress{Address: usdcMint.Bytes()}, eventConfigured.Token)
		// 	require.Equal(t, 0, len(eventConfigured.PreviousToken.Address))
		// 	require.Equal(t, usdcMint, eventConfigured.Mint)

		var eventAppended tokens.EventRemotePoolsAppended
		require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemotePoolsAppended", &eventAppended, config.PrintEvents))
		require.Equal(t, chainSelector, eventAppended.ChainSelector)
		// require.Equal(t, []cctp_token_pool.RemoteAddress{{Address: remotePoolAddress.Bytes()}}, eventAppended.PoolAddresses)
		// require.Equal(t, 0, len(eventAppended.PreviousPoolAddresses))
		require.Equal(t, usdcMint, eventAppended.Mint)

		var eventCctpEdit tokens.EventRemoteChainCctpConfigEdited
		require.NoError(t, common.ParseEvent(res.Meta.LogMessages, "RemoteChainCctpConfigChanged", &eventCctpEdit, config.PrintEvents))
		require.Equal(t, domain, eventCctpEdit.Config.DomainId)
		require.Equal(t, domainDestCaller, eventCctpEdit.Config.DestinationCaller)
	})

	t.Run("TokenAdminRegistry", func(t *testing.T) {
		t.Skip()

		t.Run("Propose token admin", func(t *testing.T) {
			ix, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
				admin.PublicKey(),
				routerConfig,
				tokenAdminRegistry,
				usdcMint,
				deployer.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, deployer, config.DefaultCommitment)
		})

		t.Run("Accept token admin", func(t *testing.T) {
			ix, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
				routerConfig,
				tokenAdminRegistry,
				usdcMint,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		})
	})

	fqPerChainPerToken, _, err := state.FindFqPerChainPerTokenConfigPDA(chainSelector, usdcMint, referenceAddresses.FeeQuoter)
	require.NoError(t, err)
	fqConfigPDA, _, err := state.FindFqConfigPDA(referenceAddresses.FeeQuoter)
	require.NoError(t, err)

	t.Run("FeeQuoter PerChainPerToken", func(t *testing.T) {
		t.Skip()

		ix, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(
			chainSelector,
			usdcMint,
			fee_quoter.TokenTransferFeeConfig{
				MinFeeUsdcents:    0,
				MaxFeeUsdcents:    1, // TODO, placeholder value
				DeciBps:           0,
				DestGasOverhead:   200000,
				DestBytesOverhead: 68, // 64 bytes for the message + 4 bytes for the length (vec prefix)
				IsEnabled:         true,
			},
			fqConfigPDA,
			fqPerChainPerToken,
			deployer.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		require.NoError(t, err)

		testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, deployer, config.DefaultCommitment)
	})

	t.Run("Router interaction (setup + onramp)", func(t *testing.T) {
		// t.Skip()

		fqUsdcBillingTokenConfig, _, err := state.FindFqBillingTokenConfigPDA(usdcMint, referenceAddresses.FeeQuoter)
		require.NoError(t, err)
		routerSigner, _, err := state.FindExternalTokenPoolsSignerPDA(cctpPool.program, referenceAddresses.Router)
		require.NoError(t, err)

		t.Run("Update pool router/rmn config", func(t *testing.T) {
			t.Skip()

			rmnIx, err := cctp_token_pool.NewSetRmnInstruction(
				referenceAddresses.RmnRemote,
				cctpPool.state,
				usdcMint,
				admin.PublicKey(),
				cctpPool.program,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			routerIx, err := cctp_token_pool.NewSetRouterInstruction(
				referenceAddresses.Router,
				cctpPool.state,
				usdcMint,
				admin.PublicKey(),
				cctpPool.program,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{rmnIx, routerIx}, admin, config.DefaultCommitment)
		})

		routerNoncesPDA, err := state.FindNoncePDA(chainSelector, admin.PublicKey(), referenceAddresses.Router)
		fmt.Printf("Router Nonces PDA: %s\n", routerNoncesPDA)
		require.NoError(t, err)

		var nonces ccip_router.Nonce
		err = common.GetAccountDataBorshInto(ctx, client, routerNoncesPDA, rpc.CommitmentConfirmed, &nonces)
		if err != nil {
			fmt.Println("WARNING: Nonce account error, initializing it !!! - This is normal for the first message of a sender to a chain")
			fmt.Println(err)
			nonces = ccip_router.Nonce{
				Version:      0,
				OrderedNonce: 0,
				TotalNonce:   0,
			}
		}
		fmt.Printf("Nonces: %v\n", nonces)

		messageSentEvent, _, err := solana.FindProgramAddress(
			[][]byte{
				[]byte("ccip_cctp_message_sent_event"),
				admin.PublicKey().Bytes(), // original sender
				common.Uint64ToLE(chainSelector),
				common.Uint64ToLE(nonces.TotalNonce + 1), // next counter, as it will be incremented for the new msg
			},
			cctpTpProgram,
		)
		fmt.Printf("Message Sent Event PDA: %s\n", messageSentEvent)
		require.NoError(t, err)

		type AttestationResponse struct {
			Attestation string `json:"attestation"`
			Status      string `json:"status"`
		}

		t.Run("Reclaim some accounts", func(t *testing.T) {
			resp, err := client.GetAccountInfoWithOpts(
				ctx,
				messageSentEvent,
				&rpc.GetAccountInfoOpts{
					Commitment: config.DefaultCommitment,
					DataSlice:  nil,
				},
			)
			if err != nil {
				t.Skip("Error fetching message sent event account, skipping message reclaiming. This is expected in most cases.")
			}
			require.NoError(t, err)
			bytes := resp.Value.Data.GetBinary()
			messageBytes := bytes[44:]
			messageHash := hex.EncodeToString(eth.Keccak256(messageBytes))

			url := fmt.Sprintf("https://iris-api-sandbox.circle.com/v1/attestations/0x%s", messageHash)
			req, err := http.NewRequest("GET", url, nil)
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)
			fmt.Printf("StatusCode: %d - Body: %s\n", res.StatusCode, body)

			var attestation []byte

			if res.StatusCode >= 200 && res.StatusCode < 300 {
				fmt.Println("Success:", res.Status)
				var attestationResponse AttestationResponse
				err = json.Unmarshal(body, &attestationResponse)
				fmt.Println("AttestationResponse:", attestationResponse)
				require.NoError(t, err)

				if attestationResponse.Status == "complete" {
					fmt.Println("Attestation completed")
					attestation, err = hex.DecodeString(attestationResponse.Attestation[2:]) // remove 0x prefix
					require.NoError(t, err)
					fmt.Println("Attestation:", attestation)
				} else {
					fmt.Println("Attestation not completed yet, retrying...")
					t.Fail()
				}
			}

			fmt.Printf("Message Hash: %s\n", messageHash)
			fmt.Printf("Attestation: %s\n", hex.EncodeToString(attestation))

			ix, err := cctp_token_pool.NewReclaimEventAccountInstruction(
				usdcMint,
				admin.PublicKey(),
				chainSelector,
				nonces.TotalNonce+1,
				attestation,
				cctpPool.state,
				cctpPool.signer,
				messageSentEvent,
				messageTransmitter.messageTransmitter,
				messageTransmitter.program,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()

			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		})

		t.Run("Create Lookup Table", func(t *testing.T) {
			t.Skip()

			tpLookupTableAddr, err = common.CreateLookupTable(ctx, client, admin)
			require.NoError(t, err)

			// These are the "real" static-only entries
			entries := solana.PublicKeySlice{
				tpLookupTableAddr,
				tokenAdminRegistry,
				cctpPool.program,
				cctpPool.state,
				cctpPool.tokenAccount, // 4 - writable
				cctpPool.signer,       // 5 - writable (to pay for event account)
				solana.TokenProgramID,
				usdcMint, // 7 - writable
				fqUsdcBillingTokenConfig,
				routerSigner,
				// -- CCTP custom entries --
				messageTransmitter.messageTransmitter, // 10 - writable
				tokenMessengerMinter.program,
				solana.SystemProgramID,
				messageTransmitter.program,
				tokenMessengerMinter.tokenMessenger,
				tokenMessengerMinter.tokenMinter,
				tokenMessengerMinter.localToken, // 16 - writable
				tokenMessengerMinter.eventAuthority,
			}

			tpLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
				tpLookupTableAddr: entries,
			}

			fmt.Printf("Lookup Table: %v\n", tpLookupTable)

			require.NoError(t, common.ExtendLookupTable(ctx, client, tpLookupTableAddr, admin, entries))
			common.AwaitSlotChange(ctx, client)
		})

		writableIndexes := []byte{4, 5, 7, 10, 16}

		t.Run("Upgrade TokenAdminRegistry", func(t *testing.T) {
			t.Skip()

			ix, err := ccip_router.NewUpgradeTokenAdminRegistryFromV1Instruction(
				routerConfig,
				tokenAdminRegistry,
				usdcMint,
				admin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		})

		t.Run("SetPool", func(t *testing.T) {
			t.Skip()

			autoDerivationIx, err := ccip_router.NewSetPoolSupportsAutoDerivationInstruction(
				usdcMint,
				true,
				routerConfig,
				tokenAdminRegistry,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			poolIx, err := ccip_router.NewSetPoolInstruction(
				writableIndexes,
				routerConfig,
				tokenAdminRegistry,
				usdcMint,
				tpLookupTableAddr,
				admin.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{autoDerivationIx, poolIx}, admin, config.DefaultCommitment)
		})

		t.Run("CCIP Send", func(t *testing.T) {
			routerBillingSignerPDA, _, err := state.FindFeeBillingSignerPDA(referenceAddresses.Router)
			require.NoError(t, err)

			adminUsdcATA, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcMint, admin.PublicKey())
			require.NoError(t, err)

			amount := uint64(1 * 1e4) // 0.01 USDC

			approveIx, err := tokens.TokenApproveChecked(amount, usdcDecimals, solana.TokenProgramID, adminUsdcATA, usdcMint, routerBillingSignerPDA, admin.PublicKey(), []solana.PublicKey{})

			message := ccip_router.SVM2AnyMessage{
				Receiver: fullReceiverAddress[:],
				Data:     []byte{},
				TokenAmounts: []ccip_router.SVMTokenAmount{
					{
						Token:  usdcMint,
						Amount: amount,
					},
				},
				FeeToken: solana.PublicKey{},
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.GenericExtraArgsV2{
					GasLimit:                 bin.Uint128{Lo: 0, Hi: 0},
					AllowOutOfOrderExecution: true,
				}, ccip.GenericExtraArgsV2Tag),
			}

			derivedAccounts, derivedLookupTables, tokenIndexes := testutils.DeriveSendAccounts(ctx, t, admin, message, chainSelector, client, referenceAddresses.Router)

			fmt.Printf("Derived Accounts: %d\n", len(derivedAccounts))
			for i, acc := range derivedAccounts {
				fmt.Printf("  %d: %s\n", i, acc.PublicKey)
			}

			builder := ccip_router.NewCcipSendInstructionBuilder().
				SetDestChainSelector(chainSelector).
				SetMessage(message).
				SetTokenIndexes(tokenIndexes)
			builder.AccountMetaSlice = derivedAccounts
			ix, err := builder.ValidateAndBuild()
			require.NoError(t, err)

			lookupTables := make(map[solana.PublicKey]solana.PublicKeySlice)
			for _, table := range derivedLookupTables {
				entries, lutErr := common.GetAddressLookupTable(ctx, client, table)
				require.NoError(t, lutErr)
				lookupTables[table] = entries
			}

			result := testutils.SendAndConfirmWithLookupTables(ctx, t, client, []solana.Instruction{approveIx, ix}, admin, config.DefaultCommitment, lookupTables)
			require.NotNil(t, result)
			tx, err := result.Transaction.GetTransaction()
			fmt.Printf("Transaction Signature: %v\n", tx.Signatures)
			fmt.Printf("Result: \n    %s\n", strings.Join(result.Meta.LogMessages, "\n    "))

			var ccipSentEvent ccip.EventCCIPMessageSent
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipSentEvent, config.PrintEvents))

			var cctpSentEvent ccip.EventCcipCctpMessageSent
			require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "CcipCctpMessageSentEvent", &cctpSentEvent, config.PrintEvents))
		})
	})
}
