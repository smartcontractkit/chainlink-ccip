package main

import (
	"reflect"
	"testing"
)

func TestBuildNetworkConfigs(t *testing.T) {
	type args struct {
		besuChainsCount   int
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
				besuChainsCount:   0,
				gethChainsCount:   1,
				solanaChainsCount: 0,
		  },
			Config{
				BesuChains: []EVMChain{},
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
			args{
				besuChainsCount:   0,
				gethChainsCount:   2,
				solanaChainsCount: 0,
		  },
			Config{
				BesuChains: []EVMChain{},
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
			args{
				besuChainsCount:   0,
				gethChainsCount:   4,
				solanaChainsCount: 0,
		  },
			Config{
				BesuChains: []EVMChain{},
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
			args{
				besuChainsCount:   0,
				gethChainsCount:   2,
				solanaChainsCount: 2,
		  },
			Config{
				BesuChains: []EVMChain{},
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
		{
			"Besu, Geth and Solana chains",
			args{
				besuChainsCount:   2,
				gethChainsCount:   2,
				solanaChainsCount: 2,
		  },
			Config{
				BesuChains: []EVMChain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     2337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
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
			if got := BuildNetworkConfigs(tt.args.besuChainsCount, tt.args.gethChainsCount, tt.args.solanaChainsCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildNetworkConfigs() = got %+v, want %+v", got, tt.want)
			}
		})
	}
}
