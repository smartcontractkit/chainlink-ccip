package main

import (
	"reflect"
	"testing"
)

func TestBuildNetworkConfigs(t *testing.T) {
	type args struct {
		chainsCount int
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			"1 chain",
			args{1},
			Config{
				Chains: []chain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
			},
		},
		{
			"2 chains",
			args{2},
			Config{
				Chains: []chain{
					{
						NetworkId:     1337,
						FinalityDepth: defaultFinalityDepth,
					},
					{
						NetworkId:     2337,
						FinalityDepth: defaultFinalityDepth,
					},
				},
			},
		},
		{
			"4 chains",
			args{4},
			Config{
				Chains: []chain{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildNetworkConfigs(tt.args.chainsCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildNetworkConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
