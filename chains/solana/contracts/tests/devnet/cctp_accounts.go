//go:build devnet
// +build devnet

package contracts

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/stretchr/testify/require"
)

type CctpTokenPoolPDAs struct {
	program,
	state,
	signer,
	tokenAccount,
	chainConfig solana.PublicKey
}

func getCctpTokenPoolPDAs(t *testing.T, program solana.PublicKey, chainSelector uint64, usdcMint solana.PublicKey) CctpTokenPoolPDAs {
	t.Helper()

	statePDA, err := tokens.TokenPoolConfigAddress(usdcMint, program)
	require.NoError(t, err)
	signerPDA, err := tokens.TokenPoolSignerAddress(usdcMint, program)
	require.NoError(t, err)
	poolTokenAccount, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, usdcMint, signerPDA)
	require.NoError(t, err)
	chainConfigPDA, _, err := tokens.TokenPoolChainConfigPDA(chainSelector, usdcMint, program)
	require.NoError(t, err)

	return CctpTokenPoolPDAs{
		program:      program,
		state:        statePDA,
		signer:       signerPDA,
		tokenAccount: poolTokenAccount,
		chainConfig:  chainConfigPDA,
	}
}

type MessageTransmitterPDAs struct {
	program,
	authorityPda,
	messageTransmitter,
	eventAuthority solana.PublicKey
}

func getMessageTransmitterPDAs(t *testing.T, program solana.PublicKey, tokenMessageMinterProgram solana.PublicKey) MessageTransmitterPDAs {
	t.Helper()

	messageTransmitterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter")}, program)
	require.NoError(t, err)
	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, program)
	require.NoError(t, err)
	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("message_transmitter_authority"), tokenMessageMinterProgram.Bytes()}, program)
	require.NoError(t, err)

	return MessageTransmitterPDAs{
		program:            program,
		messageTransmitter: messageTransmitterPDA,
		authorityPda:       authorityPda,
		eventAuthority:     eventAuthority,
	}
}

type TokenMessengerMinterPDAs struct {
	program,
	authorityPda,
	tokenMessenger,
	remoteTokenMessenger,
	tokenMinter,
	localToken,
	tokenPair,
	custodyTokenAccount,
	eventAuthority solana.PublicKey
}

func getTokenMessengerMinterPDAs(t *testing.T, program solana.PublicKey, domain uint32, usdcMint solana.PublicKey) TokenMessengerMinterPDAs {
	t.Helper()

	authorityPda, _, err := solana.FindProgramAddress([][]byte{[]byte("sender_authority")}, program)
	require.NoError(t, err)
	tokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_messenger")}, program)
	require.NoError(t, err)
	remoteTokenMessengerPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("remote_token_messenger"), []byte(common.NumToStr(domain))}, program)
	require.NoError(t, err)
	tokenMinterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("token_minter")}, program)
	require.NoError(t, err)
	eventAuthority, _, err := solana.FindProgramAddress([][]byte{[]byte("__event_authority")}, program)
	require.NoError(t, err)
	custodyTokenAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("custody"), usdcMint.Bytes()}, program)
	require.NoError(t, err)
	tokenPair, _, err := solana.FindProgramAddress([][]byte{[]byte("token_pair"), []byte(common.NumToStr(domain)), usdcMint[:]}, program) // faking that solana is again the remote domain
	require.NoError(t, err)
	localToken, _, err := solana.FindProgramAddress([][]byte{[]byte("local_token"), usdcMint.Bytes()}, program)

	return TokenMessengerMinterPDAs{
		program:              program,
		authorityPda:         authorityPda,
		tokenMessenger:       tokenMessengerPDA,
		remoteTokenMessenger: remoteTokenMessengerPDA,
		tokenMinter:          tokenMinterPDA,
		eventAuthority:       eventAuthority,
		localToken:           localToken,
		tokenPair:            tokenPair,
		custodyTokenAccount:  custodyTokenAccount,
	}
}
