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
	"strconv"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_message_transmitter"
	token_messenger_minter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/cctp_token_messenger_minter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"
)

func TestCctpDevnet(t *testing.T) {
	devnetInfo, err := getDevnetInfo()
	require.NoError(t, err)

	ctx := tests.Context(t)
	client := rpc.New(devnetInfo.RPC)

	admin := solana.PrivateKey(devnetInfo.PrivateKeys.Admin)
	require.True(t, admin.IsValid())

	tokenMessengerMinter := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.TokenMessengerMinter)
	token_messenger_minter.SetProgramID(tokenMessengerMinter)
	messageTransmitter := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.MessageTransmitter)
	message_transmitter.SetProgramID(messageTransmitter)

	usdcAddress := solana.MustPublicKeyFromBase58("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU")

	type RemoteDomain struct {
		Domain    uint32
		Recipient solana.PublicKey
	}
	domains := struct {
		Sepolia,
		Sui,
		Solana RemoteDomain
	}{ // https://developers.circle.com/stablecoins/supported-domains
		Sepolia: RemoteDomain{0, solana.MustPublicKeyFromBase58(devnetInfo.CCTP.Sepolia.RecipientBase58)},
	}
	chosenDomain := domains.Sepolia

	usdcATA, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcAddress, admin.PublicKey())
	require.NoError(t, err)

	var messageSent message_transmitter.MessageSent

	t.Run("OnRamp", func(t *testing.T) {
		// t.Skip()

		t.Parallel()

		t.Run("Deposit for Burn", func(t *testing.T) {
			pdas, err := getDepositForBurnPDAs(tokenMessengerMinter, messageTransmitter, usdcAddress, chosenDomain.Domain)

			messageSentEventKeypair, err := solana.NewRandomPrivateKey()
			// messageSentEventKeypair, err := solana.PrivateKeyFromBase58("4aufxGiqtjJySZ4RrCzAU55jKniMgKbzEhnCoVUyNxi8wcSUxJ5c7GSR2BLXF8SjDJdywL7Hydsqus27iSoRBU6n")
			require.NoError(t, err)
			fmt.Println("Message Sent Event Keypair:", messageSentEventKeypair.PublicKey(), messageSentEventKeypair)
			pdas.Print()

			ix, err := token_messenger_minter.NewDepositForBurnInstruction(
				token_messenger_minter.DepositForBurnParams{
					Amount:            1e4, // 1 cent, as USDC has 6 decimals
					DestinationDomain: uint32(chosenDomain.Domain),
					MintRecipient:     chosenDomain.Recipient,
				},
				admin.PublicKey(), // owner
				admin.PublicKey(), // event rent payer
				pdas.authorityPda,
				usdcATA,
				pdas.messageTransmitterAccount,
				pdas.tokenMessengerAccount,
				pdas.remoteTokenMessengerKey,
				pdas.tokenMinterAccount,
				pdas.localToken,
				usdcAddress,
				messageSentEventKeypair.PublicKey(),
				messageTransmitter,
				tokenMessengerMinter,
				solana.TokenProgramID,
				solana.SystemProgramID,
				pdas.eventAuthority,
				pdas.program,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment, common.AddSigners(messageSentEventKeypair))
			require.NotNil(t, result)

			tx, err := result.Transaction.GetTransaction()
			require.NoError(t, err)

			fmt.Println("Transaction signature:", tx.Signatures)
			fmt.Println(result.Meta.LogMessages)

			returnedNonce, err := common.ExtractTypedReturnValue(ctx, result.Meta.LogMessages, tokenMessengerMinter.String(), binary.LittleEndian.Uint64)
			require.NoError(t, err)
			fmt.Println("Nonce:", returnedNonce)
			fmt.Println("Result return data:", result.Meta.ReturnData.Data.String())

			require.NoError(t, common.GetAccountDataBorshInto(ctx, client, messageSentEventKeypair.PublicKey(), config.DefaultCommitment, &messageSent))
			fmt.Println("Message Sent Event Account Data:", messageSent)
			fmt.Println("Message Sent Event Bytes (hex):", hex.EncodeToString(messageSent.Message))
		})

		var attestation []byte

		type AttestationResponse struct {
			Attestation string `json:"attestation"`
			Status      string `json:"status"`
		}

		t.Run("Await attestation", func(t *testing.T) {
			messageHash := hex.EncodeToString(eth.Keccak256(messageSent.Message))

			startTime := time.Now()
			for elapsed := time.Since(startTime); elapsed < time.Second*90; elapsed = time.Since(startTime) {
				time.Sleep(time.Second * 5)

				url := fmt.Sprintf("https://iris-api-sandbox.circle.com/v1/attestations/0x%s", messageHash)
				req, err := http.NewRequest("GET", url, nil)
				require.NoError(t, err)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Accept", "application/json")

				resp, err := http.DefaultClient.Do(req)
				require.NoError(t, err)
				defer resp.Body.Close()

				body, err := ioutil.ReadAll(resp.Body)
				fmt.Printf("StatusCode: %d - Body: %s\n", resp.StatusCode, body)

				if resp.StatusCode >= 200 && resp.StatusCode < 300 {
					fmt.Println("Success:", resp.Status)
					var attestationResponse AttestationResponse
					err = json.Unmarshal(body, &attestationResponse)
					fmt.Println("AttestationResponse:", attestationResponse)
					require.NoError(t, err)

					if attestationResponse.Status == "complete" {
						fmt.Println("Attestation completed")
						attestation, err = hex.DecodeString(attestationResponse.Attestation[2:]) // remove 0x prefix
						require.NoError(t, err)
						fmt.Println("Attestation:", attestation)
						break
					} else {
						fmt.Println("Attestation not completed yet, retrying...")
						continue
					}
				}

				if resp.StatusCode >= 400 && resp.StatusCode < 500 {
					fmt.Println("Client error:", resp.Status)
					if resp.StatusCode == 404 {
						fmt.Println("Attestation not found, retrying...")
						continue
					}
					if resp.StatusCode == 429 {
						fmt.Println("Rate limit exceeded, retrying...")
						retryAfter := resp.Header.Get("Retry-After")
						if retryAfter != "" {
							retryAfterInt, err := strconv.Atoi(retryAfter)
							if err != nil {
								fmt.Println("Error parsing Retry-After header:", err)
							} else {
								fmt.Printf("Retrying after %d seconds...\n", retryAfterInt)
								time.Sleep(time.Duration(retryAfterInt) * time.Second)
							}
						} else {
							fmt.Println("Retry-After header not found, retrying after default timestep...")
						}
						continue
					}
					t.Error(fmt.Sprintf("Unexpected status code: %s", resp.Status))
					break
				}

				if resp.StatusCode >= 500 && resp.StatusCode < 600 {
					fmt.Println("Server error:", resp.Status)
					break
				}

				t.Error(fmt.Sprintf("Unexpected status code: %s", resp.Status))
				break
			}
		})
	})

	t.Run("OffRamp: Receive Message", func(t *testing.T) {
		t.Parallel()

		// t.Skip()

		messageBytes, err := hex.DecodeString(devnetInfo.CCTP.Message.MessageBytesHex)
		fmt.Println("Message bytes:", messageBytes)
		require.NoError(t, err)
		attestationBytes, err := hex.DecodeString(devnetInfo.CCTP.Message.AttestationBytesHex)
		require.NoError(t, err)

		nonce := binary.BigEndian.Uint64(messageBytes[12:20][:])

		pdas, err := getReceiveMessagePdas(tokenMessengerMinter, messageTransmitter, usdcAddress, chosenDomain.Domain, nonce)

		metas := []*solana.AccountMeta{
			solana.Meta(pdas.tokenMessengerAccount),
			solana.Meta(pdas.remoteTokenMessengerKey),
			solana.Meta(pdas.tokenMinterAccount).WRITE(),
			solana.Meta(pdas.localToken).WRITE(),
			solana.Meta(pdas.tokenPair),
			solana.Meta(usdcATA).WRITE(), // userATA
			solana.Meta(pdas.custodyTokenAccount).WRITE(),
			solana.Meta(solana.TokenProgramID),
			solana.Meta(pdas.tokenMessengerEventAuthority),
			solana.Meta(tokenMessengerMinter),
		}

		raw := message_transmitter.NewReceiveMessageInstruction(
			message_transmitter.ReceiveMessageParams{
				Message:     messageBytes,
				Attestation: attestationBytes,
			},
			admin.PublicKey(), // payer
			admin.PublicKey(), // caller
			pdas.authorityPda,
			pdas.messageTransmitterAccount,
			pdas.usedNonces,
			tokenMessengerMinter, // receiver
			solana.SystemProgramID,
			pdas.eventAuthority,
			messageTransmitter,
		)
		raw.AccountMetaSlice = append(raw.AccountMetaSlice, metas...)
		ix, err := raw.ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{ix}, admin, config.DefaultCommitment)
		require.NotNil(t, result)
		tx, err := result.Transaction.GetTransaction()
		require.NoError(t, err)
		fmt.Println("Transaction signature:", tx.Signatures)
		fmt.Println(result.Meta.LogMessages)
	})
}

type DepositForBurnPDAs struct {
	messageTransmitterAccount,
	tokenMessengerAccount,
	tokenMinterAccount,
	localToken,
	remoteTokenMessengerKey,
	authorityPda,
	eventAuthority,
	program solana.PublicKey
}

func (d *DepositForBurnPDAs) Print() {
	fmt.Println("DepositForBurnPDAs:")
	fmt.Println("  - messageTransmitterAccount:", d.messageTransmitterAccount)
	fmt.Println("  - tokenMessengerAccount:", d.tokenMessengerAccount)
	fmt.Println("  - tokenMinterAccount:", d.tokenMinterAccount)
	fmt.Println("  - localToken:", d.localToken)
	fmt.Println("  - remoteTokenMessengerKey:", d.remoteTokenMessengerKey)
	fmt.Println("  - authorityPda:", d.authorityPda)
}

func getDepositForBurnPDAs(tokenMessengerMinter, messageTransmitter, usdcAddress solana.PublicKey, domain uint32) (DepositForBurnPDAs, error) {
	messageTransmitterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, messageTransmitter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}
	tokenMessengerAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}
	tokenMinterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcAddress.Bytes()}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	remoteTokenMessengerKey, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), numToSeed(domain)}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("sender_authority")}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, tokenMessengerMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	return DepositForBurnPDAs{
		messageTransmitterAccount: messageTransmitterAccount,
		tokenMessengerAccount:     tokenMessengerAccount,
		tokenMinterAccount:        tokenMinterAccount,
		localToken:                localToken,
		remoteTokenMessengerKey:   remoteTokenMessengerKey,
		authorityPda:              authorityPda,
		eventAuthority:            eventAuthority,
		program:                   tokenMessengerMinter,
	}, nil
}

type ReceiveMessagePDAs struct {
	messageTransmitterAccount,
	tokenMessengerAccount,
	tokenMinterAccount,
	localToken,
	remoteTokenMessengerKey,
	remoteTokenKey,
	tokenPair,
	custodyTokenAccount,
	authorityPda,
	tokenMessengerEventAuthority,
	usedNonces,
	eventAuthority,
	program solana.PublicKey
}

const MAX_NONCES uint64 = 6400

func getReceiveMessagePdas(tokenMessengerMinter, messageTransmitter, usdcAddress solana.PublicKey, domain uint32, nonce uint64) (ReceiveMessagePDAs, error) {
	tokenMessengerAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	messageTransmitterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, messageTransmitter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	tokenMinterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcAddress.Bytes()}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	remoteTokenMessengerKey, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), numToSeed(domain)}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	remoteTokenKey := solana.MustPublicKeyFromBase58("111111111111Q2C3h43sGe552tUtBG3FqBKz8gX") // usdc sepolia address

	tokenPair, _, err := solana.FindProgramAddress([][]byte{[]byte("token_pair"), numToSeed(domain), remoteTokenKey.Bytes()}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	custodyTokenAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("custody"), usdcAddress.Bytes()}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter_authority"), tokenMessengerMinter.Bytes()}, messageTransmitter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	tokenMessengerEventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, tokenMessengerMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, messageTransmitter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	usedNonces, _, err := solana.FindProgramAddress([][]byte{
		[]byte("used_nonces"),
		numToSeed(domain),
		[]byte(""), // used_nonces_seed_delimiter - this is empty when the domain is a valid one
		firstNonceSeed(nonce),
	}, messageTransmitter)

	return ReceiveMessagePDAs{
		messageTransmitterAccount:    messageTransmitterAccount,
		tokenMessengerAccount:        tokenMessengerAccount,
		tokenMinterAccount:           tokenMinterAccount,
		localToken:                   localToken,
		remoteTokenMessengerKey:      remoteTokenMessengerKey,
		remoteTokenKey:               remoteTokenKey,
		tokenPair:                    tokenPair,
		custodyTokenAccount:          custodyTokenAccount,
		authorityPda:                 authorityPda,
		tokenMessengerEventAuthority: tokenMessengerEventAuthority,
		usedNonces:                   usedNonces,
		eventAuthority:               eventAuthority,
		program:                      messageTransmitter,
	}, nil
}

func firstNonceSeed(nonce uint64) []byte {
	// It looks like nonces are stored in chunks of MAX_NONCES (6400), starting from 1.
	// So, the first chunk is 1-6400, the second is 6401-12800, etc.
	// The "first_nonce" is the first nonce in the chunk that contains the given nonce.
	firstNonce := ((nonce-1)/MAX_NONCES)*MAX_NONCES + 1
	fmt.Println("First nonce:", firstNonce, "from original nonce:", nonce)
	return numToSeed(firstNonce)
}

func numToSeed[T interface {
	uint64 | uint32 | uint16 | uint8
}](num T) []byte {
	return []byte(fmt.Sprintf("%d", num))
}
