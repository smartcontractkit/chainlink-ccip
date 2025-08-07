package evmimpls_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/evmimpls/gethwrappers"
	"github.com/smartcontractkit/chainlink-ccip/chainlink-modsec/libmodsec/modsectypes"
)

func Test_DestReader(t *testing.T) {
	// Setup
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	require.NoError(t, err)

	genesisAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: new(big.Int).SetInt64(1000000000000000000)},
	}
	sim := simulated.NewBackend(genesisAlloc, simulated.WithBlockGasLimit(30e6))

	addr, _, contract, err := gethwrappers.DeployCCIPMessageSentEmitter(auth, sim.Client())
	require.NoError(t, err)
	sim.Commit()

	abi, err := gethwrappers.CCIPMessageSentEmitterMetaData.GetAbi()
	require.NoError(t, err)

	getNonceMethod, ok := abi.Methods["getNonce"]
	require.True(t, ok)
	isExecutedMethod, ok := abi.Methods["isExecuted"]
	require.True(t, ok)

	destReader := evmimpls.NewEVMDestReader(
		sim.Client(),
		addr,
		addr,
		getNonceMethod,
		isExecutedMethod,
	)

	t.Run("GetNonce", func(t *testing.T) {
		// Set a nonce onchain
		sourceChainSelector := uint64(1)
		expectedNonce := uint64(42)
		_, err := contract.SetNonce(auth, sourceChainSelector, auth.From, expectedNonce)
		require.NoError(t, err)
		sim.Commit()

		// Get nonce via dest reader
		nonce, err := destReader.GetNonce(t.Context(), sourceChainSelector, auth.From.Bytes())
		require.NoError(t, err)
		require.Equal(t, expectedNonce, nonce)
	})

	t.Run("IsExecuted", func(t *testing.T) {
		// Set executed onchain
		sourceChainSelector := uint64(1)
		sequenceNumber := uint64(123)
		_, err := contract.SetExecuted(auth, sourceChainSelector, sequenceNumber, true)
		require.NoError(t, err)
		sim.Commit()

		// Check executed via dest reader
		executed, err := destReader.IsExecuted(t.Context(), modsectypes.Message{
			Header: modsectypes.Header{
				SourceChainSelector: sourceChainSelector,
				SequenceNumber:      sequenceNumber,
			},
		})
		require.NoError(t, err)
		require.True(t, executed)
	})
}
