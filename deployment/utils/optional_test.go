package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestOptionalUnmarshalYAML(t *testing.T) {
	tests := []struct {
		name      string
		yaml      string
		wantValue int
		wantValid bool
		wantErr   bool
	}{
		{name: "verbose: value and valid false", yaml: "{value: 50, valid: false}", wantValue: 50, wantValid: false},
		{name: "verbose: value and valid true", yaml: "{value: 100, valid: true}", wantValue: 100, wantValid: true},
		{name: "semi-verbose: valid only", yaml: "valid: false", wantValue: 0, wantValid: false},
		{name: "semi-verbose: value only", yaml: "value: 32", wantValue: 32, wantValid: true},
		{name: "empty value (null)", yaml: "", wantValue: 0, wantValid: false},
		{name: "shorthand scalar", yaml: "4", wantValue: 4, wantValid: true},
		{name: "empty mapping", yaml: "{}", wantValue: 0, wantValid: false},
		{name: "explicit null", yaml: "~", wantValue: 0, wantValid: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opt Optional[int]
			err := yaml.Unmarshal([]byte(tt.yaml), &opt)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.wantValue, opt.Value)
			require.Equal(t, tt.wantValid, opt.Valid)
		})
	}
}

func TestOptionalUnmarshalYAMLResetsExistingValue(t *testing.T) {
	opt := Optional[int]{Value: 42, Valid: true}
	err := yaml.Unmarshal([]byte("~"), &opt)
	require.NoError(t, err)
	require.Equal(t, 0, opt.Value)
	require.Equal(t, false, opt.Valid)
}

func TestOptionalUnmarshalYAMLWithString(t *testing.T) {
	tests := []struct {
		name      string
		yaml      string
		wantValue string
		wantValid bool
	}{
		{name: "string semi-verbose", yaml: "value: world", wantValue: "world", wantValid: true},
		{name: "string shorthand", yaml: `"hello"`, wantValue: "hello", wantValid: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opt Optional[string]
			err := yaml.Unmarshal([]byte(tt.yaml), &opt)
			require.NoError(t, err)
			require.Equal(t, tt.wantValue, opt.Value)
			require.Equal(t, tt.wantValid, opt.Valid)
		})
	}
}

func TestOptionalUnmarshalYAMLWithBool(t *testing.T) {
	tests := []struct {
		name      string
		yaml      string
		wantValue bool
		wantValid bool
	}{
		{name: "bool semi-verbose true", yaml: "value: true", wantValue: true, wantValid: true},
		{name: "bool shorthand false", yaml: "false", wantValue: false, wantValid: true},
		{name: "bool shorthand true", yaml: "true", wantValue: true, wantValid: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opt Optional[bool]
			err := yaml.Unmarshal([]byte(tt.yaml), &opt)
			require.NoError(t, err)
			require.Equal(t, tt.wantValue, opt.Value)
			require.Equal(t, tt.wantValid, opt.Valid)
		})
	}
}
