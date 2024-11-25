package costlymessages

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	costcalculator "github.com/smartcontractkit/chainlink-ccip/pkg/costcalculator"
)

var (
	b1 = mustMakeBytes("0x01")
	b2 = mustMakeBytes("0x02")
	b3 = mustMakeBytes("0x03")
)

func mustMakeBytes(s string) ccipocr3.Bytes32 {
	b, err := ccipocr3.NewBytes32FromString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func TestCCIPCostlyMessageObserver_Observe(t *testing.T) {
	tests := []struct {
		name         string
		messageIDs   []ccipocr3.Bytes32
		messageFees  map[ccipocr3.Bytes32]plugintypes.USD18
		messageCosts map[ccipocr3.Bytes32]plugintypes.USD18
		want         []ccipocr3.Bytes32
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:         "empty",
			messageIDs:   []ccipocr3.Bytes32{},
			messageFees:  map[ccipocr3.Bytes32]plugintypes.USD18{},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			want:         []ccipocr3.Bytes32{},
			wantErr:      assert.NoError,
		},
		{
			name:         "missing fees",
			messageIDs:   []ccipocr3.Bytes32{b1, b2, b2},
			messageFees:  map[ccipocr3.Bytes32]plugintypes.USD18{},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			want:         nil,
			wantErr:      assert.Error,
		},
		{
			name:       "missing costs",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b2},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			want:         []ccipocr3.Bytes32{},
			wantErr:      assert.Error,
		},
		{
			name:       "happy path",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b3},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(5),
				b2: plugintypes.NewUSD18(25),
				b3: plugintypes.NewUSD18(15),
			},
			want:    []ccipocr3.Bytes32{b2},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			feeCalculator := costcalculator.NewStaticMessageFeeUSD18Calculator(tt.messageFees)
			execCostCalculator := costcalculator.NewStaticMessageExecCostUSD18Calculator(tt.messageCosts)
			observer := &observer{
				lggr:               logger.Test(t),
				enabled:            true,
				feeCalculator:      feeCalculator,
				execCostCalculator: execCostCalculator,
			}
			messages := make([]ccipocr3.Message, 0)
			for _, id := range tt.messageIDs {
				messages = append(messages, ccipocr3.Message{Header: ccipocr3.RampMessageHeader{MessageID: id}})
			}

			got, err := observer.Observe(ctx, messages, nil)
			if tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
