package ccip_evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

const (
	// DefaultAnvilKey that's the first Anvil and Geth private key, serves as a fallback for local testing if not overridden
	DefaultAnvilKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
)

func getNetworkPrivateKey() string {
	pk := os.Getenv("PRIVATE_KEY")
	if pk == "" {
		return DefaultAnvilKey
	}
	return pk
}

type GasSettings struct {
	FeeCapMultiplier int64 `toml:"fee_cap_multiplier"`
	TipCapMultiplier int64 `toml:"tip_cap_multiplier"`
}

const (
	DefaultNativeTransferGasLimit = 21_000
)

// ETHClient creates a basic Ethereum client using PRIVATE_KEY env var and tip/cap gas settings.
// used for common operations like funding where creating CLDF environment makes no sense.
func ETHClient(ctx context.Context, wsURL string, gasSettings *GasSettings) (*ethclient.Client, *bind.TransactOpts, string, error) {
	l := zerolog.Ctx(ctx)
	client, err := ethclient.Dial(wsURL)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not connect to eth client: %w", err)
	}
	privateKey, err := crypto.HexToECDSA(getNetworkPrivateKey())
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not parse private key: %w", err)
	}
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey).String()
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not get chain ID: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not create transactor: %w", err)
	}
	fc, tc, err := MultiplyEIP1559GasPrices(client, gasSettings.FeeCapMultiplier, gasSettings.TipCapMultiplier)
	if err != nil {
		return nil, nil, "", fmt.Errorf("could not get bumped gas price: %w", err)
	}
	auth.GasFeeCap = fc
	auth.GasTipCap = tc
	l.Info().
		Uint64("ChainID", chainID.Uint64()).
		Str("GasFeeCap", fc.String()).
		Str("GasTipCap", tc.String()).
		Msg("Default gas prices set")
	return client, auth, address, nil
}

// MultiplyEIP1559GasPrices returns bumped EIP1159 gas prices increased by multiplier.
func MultiplyEIP1559GasPrices(client *ethclient.Client, fcMult, tcMult int64) (*big.Int, *big.Int, error) {
	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, nil, err
	}
	tipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, err
	}

	return new(big.Int).Mul(feeCap, big.NewInt(fcMult)), new(big.Int).Mul(tipCap, big.NewInt(tcMult)), nil
}

// FundNodeEIP1559 funds CL node using RPC URL, recipient address and amount of funds to send (ETH).
// Uses EIP-1559 transaction type.
func FundNodeEIP1559(ctx context.Context, c *ethclient.Client, pkey, recipientAddress string, amountOfFundsInETH float64) error {
	l := zerolog.Ctx(ctx)
	amount := new(big.Float).Mul(big.NewFloat(amountOfFundsInETH), big.NewFloat(1e18))
	amountWei, _ := amount.Int(nil)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return err
	}
	privateKeyStr := strings.TrimPrefix(pkey, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	l.Info().
		Str("ChainID", chainID.String()).
		Str("From", fromAddress.String()).
		Str("Addr", recipientAddress).
		Str("Wei", amountWei.String()).
		Msg("Funding Node")

	nonce, err := c.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}
	feeCap, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	tipCap, err := c.SuggestGasTipCap(context.Background())
	if err != nil {
		return err
	}
	recipient := common.HexToAddress(recipientAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &recipient,
		Value:     amountWei,
		Gas:       DefaultNativeTransferGasLimit,
		GasFeeCap: feeCap,
		GasTipCap: tipCap,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return err
	}
	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	if _, err := bind.WaitMined(context.Background(), c, signedTx); err != nil {
		return err
	}
	l.Info().Str("Wei", amountWei.String()).Msg("Funded with ETH")
	return nil
}
