package fees

import (
	"math"

	chainsel "github.com/smartcontractkit/chain-selectors"
)

func GetDefaultChainAgnosticTokenTransferFeeConfig(src uint64, dst uint64, overrides ...func(*TokenTransferFeeArgs)) TokenTransferFeeArgs {
	minFeeUSDCents := uint32(25)

	// NOTE: we validate that src != dst so only one of these if statements will execute
	if src == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 50
	}
	if dst == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 150
	}

	cfg := TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    minFeeUSDCents,
		MaxFeeUSDCents:    math.MaxUint32,
		DeciBps:           0,
		IsEnabled:         true,
	}

	for _, override := range overrides {
		override(&cfg)
	}

	return cfg
}
