package main

import (
	"os"
	"path"
	"reflect"
	"testing"
)

func TestBuildTemplateConfig(t *testing.T) {
	tmpDir := t.TempDir()
	defaultOverrideFile := DefaultOverrideFile(t, tmpDir)
	overrideFileWith2Chains := OverrideFileWith2Chains(t, tmpDir)

	type args struct {
		overridesFilePath string
		chainsCount       int
	}
	tests := []struct {
		name    string
		args    args
		want    ValuesTmplConfig
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				overridesFilePath: path.Join(tmpDir, defaultOverrideFile),
				chainsCount:       2,
			}, want: ValuesTmplConfig{
			Chains: []chain{
				{NetworkId: 1337, BlockTime: 0},
				{NetworkId: 2337, BlockTime: 0},
			},
		}},
		{
			name: "just one chain",
			args: args{
				overridesFilePath: path.Join(tmpDir, defaultOverrideFile),
				chainsCount:       1,
			}, want: ValuesTmplConfig{
			Chains: []chain{
				{NetworkId: 1337, BlockTime: 0},
			},
		}},
		{
			name: "additional chains",
			args: args{
				overridesFilePath: path.Join(tmpDir, overrideFileWith2Chains),
				chainsCount:       4,
			}, want: ValuesTmplConfig{
			Chains: []chain{
				{NetworkId: 1337, BlockTime: 20},
				{NetworkId: 2337, BlockTime: 14},
				{NetworkId: 90000001, BlockTime: 0},
				{NetworkId: 90000002, BlockTime: 0},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildTemplateConfig(tt.args.overridesFilePath, tt.args.chainsCount, "", "", "")
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildTemplateConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTemplateConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func DefaultOverrideFile(t *testing.T, tmpDir string) string {
	filename := "default.yaml"
	tmpFile := path.Join(tmpDir, filename)
	overridesContent := `
chains:
`
	err := os.WriteFile(tmpFile, []byte(overridesContent), 0644)
	if err != nil {
		t.Fatal("Failed to write to file", err)
	}
	return filename
}

func OverrideFileWith2Chains(t *testing.T, tmpDir string) string {
	filename := "2chains.yaml"
	tmpFile := path.Join(tmpDir, filename)
	overridesContent := `
chains:
  - networkId: 1337
    blockTime: 20
  - networkId: 2337
    blockTime: 14

`
	err := os.WriteFile(tmpFile, []byte(overridesContent), 0644)
	if err != nil {
		t.Fatal("Failed to write to file", err)
	}
	return filename
}
