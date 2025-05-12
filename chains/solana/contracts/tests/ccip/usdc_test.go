package contracts

import (
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	message_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_message_transmitter"
	token_messenger_minter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/cctp_token_message_minter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
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

	tokenMessageMinter := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.TokenMessageMinter)
	token_messenger_minter.SetProgramID(tokenMessageMinter)
	messageTransmitter := solana.MustPublicKeyFromBase58(devnetInfo.CCTP.MessageTransmitter)
	message_transmitter.SetProgramID(messageTransmitter)

	usdcAddress := solana.MustPublicKeyFromBase58("4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU")
	domains := struct {
		Sepolia,
		Sui,
		Solana uint32
	}{ // https://developers.circle.com/stablecoins/supported-domains
		Sepolia: 0,
		Solana:  5,
		Sui:     8,
	}
	chosenDomain := domains.Sui

	usdcATA, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcAddress, admin.PublicKey())
	require.NoError(t, err)

	t.Run("OnRamp: Deposit for Burn", func(t *testing.T) {
		t.Skip()

		pdas, err := getDepositForBurnPDAs(tokenMessageMinter, messageTransmitter, usdcAddress, chosenDomain)

		// messageSentEventKeypair, err := solana.NewRandomPrivateKey()
		messageSentEventKeypair, err := solana.PrivateKeyFromBase58("4aufxGiqtjJySZ4RrCzAU55jKniMgKbzEhnCoVUyNxi8wcSUxJ5c7GSR2BLXF8SjDJdywL7Hydsqus27iSoRBU6n")
		require.NoError(t, err)
		fmt.Println("Message Sent Event Keypair:", messageSentEventKeypair.PublicKey(), messageSentEventKeypair)
		pdas.Print()

		ix, err := token_messenger_minter.NewDepositForBurnInstruction(
			token_messenger_minter.DepositForBurnParams{
				Amount:            1e4, // 1 cent, as USDC has 6 decimals
				DestinationDomain: uint32(chosenDomain),
				MintRecipient:     admin.PublicKey(),
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
			tokenMessageMinter,
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
	})

	t.Run("OffRamp: Receive Message", func(t *testing.T) {})
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

func getDepositForBurnPDAs(tokenMessageMinter, messageTransmitter, usdcAddress solana.PublicKey, domain uint32) (DepositForBurnPDAs, error) {
	messageTransmitterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, messageTransmitter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}
	tokenMessengerAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, tokenMessageMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}
	tokenMinterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, tokenMessageMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcAddress.Bytes()}, tokenMessageMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	remoteTokenMessengerKey, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), numToSeed(domain)}, tokenMessageMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("sender_authority")}, tokenMessageMinter)
	if err != nil {
		return DepositForBurnPDAs{}, err
	}

	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, tokenMessageMinter)
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
		program:                   tokenMessageMinter,
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

func getReceiveMessagePdas(tokenMessageMinter, messageTransmitter, usdcAddress solana.PublicKey, domain uint32, nonce uint64) (ReceiveMessagePDAs, error) {
	tokenMessengerAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	messageTransmitterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, messageTransmitter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	tokenMinterAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcAddress.Bytes()}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	remoteTokenMessengerKey, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), numToSeed(domain)}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	// const remoteTokenKey = new PublicKey(hexToBytes(remoteUsdcAddressHex));
	remoteTokenKey := solana.PublicKey{} // TODO

	tokenPair, _, err := solana.FindProgramAddress([][]byte{[]byte("token_pair"), numToSeed(domain), remoteTokenKey.Bytes()}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	custodyTokenAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("custody"), usdcAddress.Bytes()}, tokenMessageMinter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter_authority"), tokenMessageMinter.Bytes()}, messageTransmitter)
	if err != nil {
		return ReceiveMessagePDAs{}, err
	}

	tokenMessengerEventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, tokenMessageMinter)
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
	return numToSeed(firstNonce)
}

func numToSeed[T interface {
	uint64 | uint32 | uint16 | uint8
}](num T) []byte {
	return []byte(fmt.Sprintf("%d", num))
}
