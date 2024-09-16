package exectypes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageTokensData(t *testing.T) {
	tests := []struct {
		name         string
		msgTokenData MessageTokensData
		ready        bool
		error        error
	}{
		{
			name:         "empty MessageTokenData is always ready - message doesnt carry tokens",
			msgTokenData: MessageTokensData{},
			ready:        true,
			error:        nil,
		},
		{
			name: "MessageTokenData is ready - all tokens are ready",
			msgTokenData: MessageTokensData{
				TokenData: []TokenData{
					{
						Ready: true,
						Data:  []byte{123},
					},
					{
						Ready: true,
						Data:  []byte{234},
					},
				},
			},
			ready: true,
			error: nil,
		},
		{
			name: "MessageTokenData is not ready - one token is not ready",
			msgTokenData: MessageTokensData{
				TokenData: []TokenData{
					{
						Ready: true,
						Data:  []byte{123},
					},
					{
						Ready: false,
						Data:  nil,
						Error: fmt.Errorf("some error"),
					},
				},
			},
			ready: false,
			error: fmt.Errorf("some error"),
		},
		{
			name: "MessageTokenData is not ready - all tokens are not ready, first error is returned",
			msgTokenData: MessageTokensData{
				TokenData: []TokenData{
					{
						Ready: false,
						Data:  nil,
						Error: fmt.Errorf("error1"),
					},
					{
						Ready: false,
						Data:  nil,
						Error: fmt.Errorf("error2"),
					},
				},
			},
			ready: false,
			error: fmt.Errorf("error1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.ready, tt.msgTokenData.IsReady())
			assert.Equal(t, tt.error, tt.msgTokenData.Error())
		})
	}
}

func Test_MessageTokensData_ToByteSlice(t *testing.T) {
	tests := []struct {
		name         string
		msgTokenData MessageTokensData
		expected     [][]byte
	}{
		{
			name:         "empty MessageTokenData is always ready - message doesnt carry tokens",
			msgTokenData: MessageTokensData{},
			expected:     [][]byte{},
		},
		{
			name: "MessageTokenData is ready - all tokens are ready",
			msgTokenData: MessageTokensData{
				TokenData: []TokenData{
					{
						Ready: true,
						Data:  []byte{123},
					},
					{
						Ready: true,
						Data:  []byte{234},
					},
				},
			},
			expected: [][]byte{
				{123},
				{234},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.msgTokenData.ToByteSlice())
		})
	}
}
