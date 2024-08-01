package chainconfig

import (
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestChainConfig_Validate(t *testing.T) {
	type fields struct {
		GasPriceDeviationPPB    cciptypes.BigInt
		DAGasPriceDeviationPPB  cciptypes.BigInt
		FinalityDepth           int64
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
				GasPriceDeviationPPB:    cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB:  cciptypes.BigInt{Int: big.NewInt(1)},
				FinalityDepth:           1,
				OptimisticConfirmations: 1,
			},
			false,
		},
		{
			"valid, finality tags enabled",
			fields{
				GasPriceDeviationPPB:    cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB:  cciptypes.BigInt{Int: big.NewInt(1)},
				FinalityDepth:           -1,
				OptimisticConfirmations: 1,
			},
			false,
		},
		{
			"invalid, gas price deviation not set",
			fields{
				GasPriceDeviationPPB:    cciptypes.BigInt{Int: big.NewInt(0)},
				DAGasPriceDeviationPPB:  cciptypes.BigInt{Int: big.NewInt(1)},
				FinalityDepth:           1,
				OptimisticConfirmations: 1,
			},
			true,
		},
		{
			"invalid, finality depth not set",
			fields{
				GasPriceDeviationPPB:    cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB:  cciptypes.BigInt{Int: big.NewInt(1)},
				FinalityDepth:           0,
				OptimisticConfirmations: 1,
			},
			true,
		},
		{
			"invalid, optimistic confirmations not set",
			fields{
				GasPriceDeviationPPB:    cciptypes.BigInt{Int: big.NewInt(1)},
				DAGasPriceDeviationPPB:  cciptypes.BigInt{Int: big.NewInt(1)},
				FinalityDepth:           1,
				OptimisticConfirmations: 0,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := ChainConfig{
				GasPriceDeviationPPB:    tt.fields.GasPriceDeviationPPB,
				DAGasPriceDeviationPPB:  tt.fields.DAGasPriceDeviationPPB,
				FinalityDepth:           tt.fields.FinalityDepth,
				OptimisticConfirmations: tt.fields.OptimisticConfirmations,
			}
			if err := cc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ChainConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
