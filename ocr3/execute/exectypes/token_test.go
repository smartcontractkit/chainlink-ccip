package exectypes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MessageTokensData(t *testing.T) {
	tests := []struct {
		name         string
		msgTokenData MessageTokenData
		ready        bool
		errorMsg     string
	}{
		{
			name:         "empty MessageTokenData is always ready - message doesnt carry tokens",
			msgTokenData: MessageTokenData{},
			ready:        true,
			errorMsg:     "",
		},
		{
			name: "MessageTokenData is ready - all tokens are ready",
			msgTokenData: MessageTokenData{
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
		},
		{
			name: "MessageTokenData is not ready - one token is not ready",
			msgTokenData: MessageTokenData{
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
			ready:    false,
			errorMsg: "some error",
		},
		{
			name: "MessageTokenData is not ready - all tokens are not ready, first error is returned",
			msgTokenData: MessageTokenData{
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
			ready:    false,
			errorMsg: "error1\nerror2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.ready, tt.msgTokenData.IsReady())
			if tt.errorMsg == "" {
				assert.Nil(t, tt.msgTokenData.Error())
			} else {
				assert.Equal(t, tt.errorMsg, tt.msgTokenData.Error().Error())
			}
		})
	}
}

func Test_MessageTokensData_ToByteSlice(t *testing.T) {
	tests := []struct {
		name         string
		msgTokenData MessageTokenData
		expected     [][]byte
	}{
		{
			name:         "empty MessageTokenData is always ready - message doesnt carry tokens",
			msgTokenData: MessageTokenData{},
			expected:     [][]byte{},
		},
		{
			name: "MessageTokenData is ready - all tokens are ready",
			msgTokenData: MessageTokenData{
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

func Test_MessageTokenData_Append(t *testing.T) {
	t.Run("append token to the firs empty slot", func(t *testing.T) {
		msg := NewMessageTokenData()
		msg = msg.Append(0, NewSuccessTokenData([]byte{1}))
		assert.Equal(t, 1, len(msg.TokenData))
		assert.True(t, msg.IsReady())
	})

	t.Run("append error to the next slot", func(t *testing.T) {
		msg := NewMessageTokenData()
		msg = msg.Append(0, NewSuccessTokenData([]byte{1}))
		msg = msg.Append(1, NewErrorTokenData(fmt.Errorf("token not supported")))
		assert.Equal(t, 2, len(msg.TokenData))
		assert.False(t, msg.IsReady())
		assert.Equal(t, "token not supported", msg.Error().Error())
	})

	t.Run("append token to the very last slot", func(t *testing.T) {
		msg := NewMessageTokenData()
		msg = msg.Append(10, NewSuccessTokenData([]byte{10}))
		assert.Equal(t, 11, len(msg.TokenData))
		assert.False(t, msg.IsReady())
		assert.Nil(t, msg.Error())

		for i := 0; i < 10; i++ {
			assert.Nil(t, msg.TokenData[i].Data)
			msg = msg.Append(i, NewSuccessTokenData([]byte{byte(i)}))
		}
		assert.True(t, msg.IsReady())
	})
}
