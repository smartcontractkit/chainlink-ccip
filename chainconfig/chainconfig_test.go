package chainconfig

import (
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestChainConfig_Validate(t *testing.T) {
	type fields struct {
		GasPriceDeviationPPB    cciptypes.BigInt
		DAGasPriceDeviationPPB  cciptypes.BigInt
		OptimisticConfirmations uint32
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"valid",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(1)},
			},
			false,
		},
		{
			"invalid, gas price deviation is nil",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(1)},
			},
			true,
		},
		{
			"invalid, da gas price deviation is nil",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{},
			},
			true,
		},
		{
			"invalid, gas price deviation not set",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(0)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(1)},
			},
			true,
		},
		{
			"invalid, gas price deviation is negative",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(-1)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(1)},
			},
			true,
		},
		{
			"invalid, data-availability gas price deviation not set",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(0)},
			},
			true,
		},
		{
			"invalid, data-availability gas price deviation is negative",
			fields{
				GasPriceDeviationPPB:   cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB: cciptypes.BigInt{Int: big.NewInt(-1)},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := ChainConfig{
				GasPriceDeviationPPB:    tt.fields.GasPriceDeviationPPB,
				DAGasPriceDeviationPPB:  tt.fields.DAGasPriceDeviationPPB,
				OptimisticConfirmations: tt.fields.OptimisticConfirmations,
			}
			if err := cc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ChainConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
