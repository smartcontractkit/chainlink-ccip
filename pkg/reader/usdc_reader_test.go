package reader

import (
	"errors"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	reader "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_USDCMessageReader_New(t *testing.T) {
	tests := []struct {
		name         string
		tokensConfig map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig
		readers      map[cciptypes.ChainSelector]*reader.MockContractReaderFacade
		errorMessage string
	}{
		{
			name:         "empty tokens and readers works",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{},
			readers:      map[cciptypes.ChainSelector]*reader.MockContractReaderFacade{},
		},
		{
			name: "missing readers",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(1): {
					SourcePoolAddress:            "0xA",
					SourceMessageTransmitterAddr: "0xB",
				},
			},
			readers:      map[cciptypes.ChainSelector]*reader.MockContractReaderFacade{},
			errorMessage: "no contract reader found for chain 1",
		},
		{
			name: "binding errors",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(1): {
					SourcePoolAddress:            "0xA",
					SourceMessageTransmitterAddr: "0xB",
				},
			},
			readers: func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade {
				readers := make(map[cciptypes.ChainSelector]*reader.MockContractReaderFacade)
				m := reader.NewMockContractReaderFacade(t)
				m.EXPECT().Bind(mock.Anything, mock.Anything).Return(errors.New("error"))
				readers[cciptypes.ChainSelector(1)] = m
				return readers
			}(),
			errorMessage: "unable to bind MessageTransmitter for chain 1",
		},
		{
			name: "happy path",
			tokensConfig: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				cciptypes.ChainSelector(1): {
					SourcePoolAddress:            "0xA",
					SourceMessageTransmitterAddr: "0xB",
				},
			},
			readers: func() map[cciptypes.ChainSelector]*reader.MockContractReaderFacade {
				readers := make(map[cciptypes.ChainSelector]*reader.MockContractReaderFacade)
				m := reader.NewMockContractReaderFacade(t)
				m.EXPECT().Bind(mock.Anything, []types.BoundContract{
					{
						Address: "0xB",
						Name:    consts.ContractNameCCTPMessageTransmitter,
					},
				}).Return(nil)

				readers[cciptypes.ChainSelector(1)] = m
				return readers
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readers := make(map[cciptypes.ChainSelector]types.ContractReader)
			for k, v := range tt.readers {
				readers[k] = v
			}

			r, err := NewUSDCMessageReader(tt.tokensConfig, readers)
			if tt.errorMessage != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorMessage)
			} else {
				require.NoError(t, err)
				require.NotNil(t, r)
			}
		})
	}
}

func Test_USDCMessageReader_MessageHashes(t *testing.T) {

}

func Test_MessageSentEvent_unpackID(t *testing.T) {
	nonEmptyEvent := eventID{}
	for i := 0; i < 32; i++ {
		nonEmptyEvent[i] = byte(i)
	}

	fullPayload := make([]byte, 0, 64)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)

	tests := []struct {
		name    string
		data    []byte
		want    eventID
		wantErr bool
	}{
		{
			name:    "event too short",
			data:    make([]byte, 31),
			wantErr: true,
		},
		{
			name: "event with proper length but empty",
			data: make([]byte, 32),
			want: eventID{},
		},
		{
			name: "event with proper length and data",
			data: fullPayload,
			want: nonEmptyEvent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := MessageSentEvent{Arg0: tt.data}
			got, err := event.unpackID()

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})

	}
}

func Test_SourceTokenDataPayload_ToBytes(t *testing.T) {
	tests := []struct {
		nonce        uint64
		sourceDomain uint32
	}{
		{
			nonce:        0,
			sourceDomain: 0,
		},
		{
			nonce:        2137,
			sourceDomain: 4,
		},
		{
			nonce:        2,
			sourceDomain: 2137,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("nonce=%d,sourceDomain=%d", test.nonce, test.sourceDomain), func(t *testing.T) {
			payload1 := NewSourceTokenDataPayload(test.nonce, test.sourceDomain)
			bytes := payload1.ToBytes()

			payload2, err := NewSourceTokenDataPayloadFromBytes(bytes)
			require.NoError(t, err)
			require.Equal(t, *payload1, *payload2)
		})
	}
}

func Test_SourceTokenDataPayload_FromBytes(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *SourceTokenDataPayload
		wantErr bool
	}{
		{
			name:    "too short data",
			data:    []byte{0x01, 0x02, 0x03},
			wantErr: true,
		},
		{
			name: "empty data but with proper length",
			data: make([]byte, 12),
			want: NewSourceTokenDataPayload(0, 0),
		},
		{
			name: "data with nonce and source domain",
			data: []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 2},
			want: NewSourceTokenDataPayload(1, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSourceTokenDataPayloadFromBytes(tt.data)

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})
	}
}
