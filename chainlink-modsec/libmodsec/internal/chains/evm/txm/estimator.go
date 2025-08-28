package txm

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

// This wrapper is required because of the way Gas Estimator components expect configs.
// Instead of passing down a struct with values, we need to imlpement an interface with
// the required methods.
type GasEstimator struct {
	EIP1559DynamicFeesF bool
	BumpPercentF        uint16
	BumpThresholdF      uint64
	BumpTxDepthF        uint32
	BumpMinF            *assets.Wei
	FeeCapDefaultF      *assets.Wei
	LimitDefaultF       uint64
	LimitMaxF           uint64
	LimitMultiplierF    float32
	LimitTransferF      uint64
	PriceDefaultF       *assets.Wei
	TipCapDefaultF      *assets.Wei
	TipCapMinF          *assets.Wei
	PriceMaxF           *assets.Wei
	PriceMinF           *assets.Wei
	ModeF               string
	EstimateLimitF      bool
	SenderAddressF      types.EIP55Address
}

func (g GasEstimator) PriceMaxKey(common.Address) *assets.Wei {
	return g.PriceMaxF
}

func (g GasEstimator) EIP1559DynamicFees() bool {
	return g.EIP1559DynamicFeesF
}

func (b GasEstimator) BumpPercent() uint16 {
	return b.BumpPercentF
}

func (b GasEstimator) BumpThreshold() uint64 {
	return b.BumpThresholdF
}

func (b GasEstimator) BumpTxDepth() uint32 {
	return b.BumpTxDepthF
}
func (b GasEstimator) BumpMin() *assets.Wei {
	return b.BumpMinF
}

func (b GasEstimator) TipCapMin() *assets.Wei {
	return b.TipCapMinF
}

func (b GasEstimator) PriceMax() *assets.Wei {
	return b.PriceMaxF
}

func (b GasEstimator) PriceMin() *assets.Wei {
	return b.PriceMinF
}

func (b GasEstimator) Mode() string {
	return b.ModeF
}

func (b GasEstimator) PriceDefault() *assets.Wei {
	return b.PriceDefaultF
}

func (b GasEstimator) TipCapDefault() *assets.Wei {
	return b.TipCapDefaultF
}

func (b GasEstimator) FeeCapDefault() *assets.Wei {
	return b.FeeCapDefaultF
}

func (b GasEstimator) LimitDefault() uint64 {
	return b.LimitDefaultF
}

func (b GasEstimator) LimitMax() uint64 {
	return b.LimitMaxF
}

func (b GasEstimator) LimitMultiplier() float32 {
	return b.LimitMultiplierF
}

func (b GasEstimator) LimitTransfer() uint64 {
	return b.LimitTransferF
}

func (b GasEstimator) EstimateLimit() bool {
	return b.EstimateLimitF
}

func (b GasEstimator) SenderAddress() *types.EIP55Address {
	return &b.SenderAddressF
}

// -------------------------------
func (g GasEstimator) DAOracle() config.DAOracle {
	return &DAOracle{}
}

type DAOracle struct {
	OracleTypeF             *toml.DAOracleType
	OracleAddressF          *types.EIP55Address
	CustomGasPriceCalldataF *string
}

func (o DAOracle) OracleType() *toml.DAOracleType {
	return o.OracleTypeF
}

func (o DAOracle) OracleAddress() *types.EIP55Address {
	return o.OracleAddressF
}

func (o DAOracle) CustomGasPriceCalldata() *string {
	return o.CustomGasPriceCalldataF
}

// -------------------------------
func (g GasEstimator) LimitJobType() config.LimitJobType {
	return nil
}

// -------------------------------
func (g GasEstimator) FeeHistory() config.FeeHistory {
	return &FeeHistory{
		CacheTimeoutF: 5 * time.Second,
	}
}

type FeeHistory struct {
	CacheTimeoutF time.Duration
}

func (b FeeHistory) CacheTimeout() time.Duration {
	return b.CacheTimeoutF
}

// -------------------------------
func (g GasEstimator) BlockHistory() config.BlockHistory {
	return &BlockHistory{
		BlockHistorySizeF:      4,
		TransactionPercentileF: 55,
	}
}

type BlockHistory struct {
	BatchSizeF                 uint32
	BlockHistorySizeF          uint16
	BlockDelayF                uint16
	CheckInclusionBlocksF      uint16
	CheckInclusionPercentileF  uint16
	EIP1559FeeCapBufferBlocksF uint16
	TransactionPercentileF     uint16
	SenderAddressF             types.EIP55Address
}

func (b BlockHistory) BatchSize() uint32 {
	return b.BatchSizeF
}

func (b BlockHistory) BlockHistorySize() uint16 {
	return b.BlockHistorySizeF
}

func (b BlockHistory) BlockDelay() uint16 {
	return b.BlockDelayF
}

func (b BlockHistory) CheckInclusionBlocks() uint16 {
	return b.CheckInclusionBlocksF
}

func (b BlockHistory) CheckInclusionPercentile() uint16 {
	return b.CheckInclusionPercentileF
}

func (b BlockHistory) EIP1559FeeCapBufferBlocks() uint16 {
	return b.EIP1559FeeCapBufferBlocksF
}

func (b BlockHistory) TransactionPercentile() uint16 {
	return b.TransactionPercentileF
}

func (b BlockHistory) SenderAddress() types.EIP55Address {
	return b.SenderAddressF
}

func GetGasConfig() GasEstimator {
	// Configs for Ethereum Sepolia
	gasConfig := GasEstimator{
		EIP1559DynamicFeesF: true,
		BumpPercentF:        20,
		LimitDefaultF:       30000,
		LimitMultiplierF:    1,
		PriceMaxF:           assets.GWei(700),
		ModeF:               "FeeHistory",
		SenderAddressF:      types.MustEIP55Address("0x0000000000000000000000000000000000000000"),
	}

	return gasConfig
}

func InitGasEstimator(lggr logger.Logger, client client.Client, chainID *big.Int, gasConfig config.GasEstimator) (gas.EvmFeeEstimator, error) {
	estimator, err := gas.NewEstimator(lggr, client, "", chainID, gasConfig, nil)
	if err != nil {
		lggr.Fatal("Couldn't create estimator!", err)
		return nil, err
	}
	err = estimator.Start(context.TODO())
	if err != nil {
		lggr.Fatal("Couldn't start estimator!", err)
		return nil, err
	}
	// Committing a synchronization sin for the purposes of this test to make sure the estimator has up-to-date prices.
	time.Sleep(3 * time.Second)
	return estimator, nil
}
