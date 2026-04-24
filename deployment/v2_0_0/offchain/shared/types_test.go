package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProductionEnvironment(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want bool
	}{
		{name: "prod mainnet", env: "prod_mainnet", want: true},
		{name: "prod testnet", env: "prod_testnet", want: true},
		{name: "legacy mainnet", env: "mainnet", want: true},
		{name: "non prod", env: "test", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, IsProductionEnvironment(tt.env))
		})
	}
}
