package cctp

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

//////////////////////////
// CCTP testing helpers //
//////////////////////////

type MessageTransmitterPDAs struct {
	Program,
	AuthorityPda,
	MessageTransmitter,
	EventAuthority solana.PublicKey
}

func GetMessageTransmitterPDAs(t *testing.T) MessageTransmitterPDAs {
	messageTransmitterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, config.CctpMessageTransmitter)
	require.NoError(t, err)
	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, config.CctpMessageTransmitter)
	require.NoError(t, err)
	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter_authority"), config.CctpTokenMessengerMinter.Bytes()}, config.CctpMessageTransmitter)
	require.NoError(t, err)

	return MessageTransmitterPDAs{
		Program:            config.CctpMessageTransmitter,
		MessageTransmitter: messageTransmitterPDA,
		AuthorityPda:       authorityPda,
		EventAuthority:     eventAuthority,
	}
}

type TokenMessengerMinterPDAs struct {
	Program,
	AuthorityPda,
	TokenMessenger,
	RemoteTokenMessenger,
	TokenMinter,
	LocalToken,
	TokenPair,
	CustodyTokenAccount,
	EventAuthority solana.PublicKey
}

func GetTokenMessengerMinterPDAs(t *testing.T, domain uint32, usdcMint solana.PublicKey, remoteTokenAddr []byte) TokenMessengerMinterPDAs {
	remoteTokenAddrBytes := make([]byte, 32)
	copy(remoteTokenAddrBytes[32-len(remoteTokenAddr):], remoteTokenAddr)
	remoteTokenPubkey := solana.PublicKeyFromBytes(remoteTokenAddrBytes)

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("sender_authority")}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	tokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	remoteTokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), []byte(common.NumToStr(domain))}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	tokenMinterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	custodyTokenAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("custody"), usdcMint.Bytes()}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)
	tokenPair, _, err := solana.FindProgramAddress([][]byte{[]byte("token_pair"), []byte(common.NumToStr(domain)), remoteTokenPubkey[:]}, config.CctpTokenMessengerMinter) // faking that solana is again the remote domain
	require.NoError(t, err)
	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcMint.Bytes()}, config.CctpTokenMessengerMinter)
	require.NoError(t, err)

	return TokenMessengerMinterPDAs{
		Program:              config.CctpTokenMessengerMinter,
		AuthorityPda:         authorityPda,
		TokenMessenger:       tokenMessengerPDA,
		RemoteTokenMessenger: remoteTokenMessengerPDA,
		TokenMinter:          tokenMinterPDA,
		EventAuthority:       eventAuthority,
		LocalToken:           localToken,
		TokenPair:            tokenPair,
		CustodyTokenAccount:  custodyTokenAccount,
	}
}

type TokenPoolPDAs struct {
	Program,
	GlobalConfig,
	State,
	Signer,
	TokenAccount,
	SvmChainConfig solana.PublicKey
}

func GetTokenPoolPDAs(t *testing.T, usdcMint solana.PublicKey) TokenPoolPDAs {
	t.Helper()

	statePDA, err := tokens.TokenPoolConfigAddress(usdcMint, config.CctpTokenPoolProgram)
	require.NoError(t, err)
	globalConfigPDA, err := tokens.TokenPoolGlobalConfigPDA(config.CctpTokenPoolProgram)
	require.NoError(t, err)
	signerPDA, err := tokens.TokenPoolSignerAddress(usdcMint, config.CctpTokenPoolProgram)
	require.NoError(t, err)
	poolTokenAccount, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcMint, signerPDA)
	require.NoError(t, err)
	chainConfigPDA, _, err := tokens.TokenPoolChainConfigPDA(config.SvmChainSelector, usdcMint, config.CctpTokenPoolProgram)
	require.NoError(t, err)

	return TokenPoolPDAs{
		Program:        config.CctpTokenPoolProgram,
		GlobalConfig:   globalConfigPDA,
		State:          statePDA,
		Signer:         signerPDA,
		TokenAccount:   poolTokenAccount,
		SvmChainConfig: chainConfigPDA,
	}
}

func GetNonce(messageBytes []byte) uint64 {
	nonceBytes := messageBytes[12:20] // extract nonce from message, which is the 12th to 20th byte of the message
	nonce := binary.BigEndian.Uint64(nonceBytes)
	return nonce
}

func SetNonce(messageBytes *[]byte, nonce uint64) {
	binary.BigEndian.PutUint64((*messageBytes)[12:20], nonce) // set nonce in message, which is the 12th to 20th byte of the message
}

func GetSrcDomain(cctpMessage []byte) uint32 {
	srcDomainBytes := cctpMessage[4:8] // extract source domain from message, which is the 4th to 8th byte of the message
	srcDomain := binary.BigEndian.Uint32(srcDomainBytes)
	return srcDomain
}

func SetSrcDomain(cctpMessage *[]byte, srcDomain uint32) {
	binary.BigEndian.PutUint32((*cctpMessage)[4:8], srcDomain) // set source domain in message, which is the 4th to 8th byte of the message
}

func GetUsedNoncesPDA(t *testing.T, cctpMessage []byte) solana.PublicKey {
	t.Helper()
	nonce := GetNonce(cctpMessage)
	fmt.Println("Nonce for receiving message:", nonce)
	firstNonce := (nonce-1)/6400*6400 + 1 // round down to the first nonce that is a multiple of 6400
	fmt.Println("First nonce:", firstNonce)
	fmt.Println("Domain:", GetSrcDomain(cctpMessage))
	firstNoncePda, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("used_nonces"),
			[]byte(common.NumToStr(GetSrcDomain(cctpMessage))),
			[]byte(common.NumToStr(firstNonce)),
		},
		config.CctpMessageTransmitter,
	)
	require.NoError(t, err)
	return firstNoncePda
}
