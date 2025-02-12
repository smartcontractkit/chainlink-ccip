package main

import (
	"reflect"
	"testing"
)

func TestBuildNetworkConfigs(t *testing.T) {
	type args struct {
		gethChainsCount   int
		solanaChainsCount int
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			"1 chain",
			args{
				1,
				0,
			},
			Config{
				GethChains: []EVMChain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
				SolanaChains: []SolanaChain{},
			},
		},
		{
			"2 chains",
			args{2, 0},
			Config{
				GethChains: []EVMChain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     2337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
				SolanaChains: []SolanaChain{},
			},
		},
		{
			"4 chains",
			args{4, 0},
			Config{
				GethChains: []EVMChain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     2337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     90000001,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     90000002,
						FinalityDepth: defaultFinalityDepth,
					},
				},
				SolanaChains: []SolanaChain{},
			},
		},
		{
			"Geth and Solana chains",
			args{2, 2},
			Config{
				GethChains: []EVMChain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     2337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
				SolanaChains: []SolanaChain{
					{
						ChainId: 1001,
					},
					{
						ChainId: 1002,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildNetworkConfigs(tt.args.gethChainsCount, tt.args.solanaChainsCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildNetworkConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
