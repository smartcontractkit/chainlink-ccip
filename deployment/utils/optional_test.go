package utils

import (
	"testing"

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
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if opt.Value != tt.wantValue {
				t.Errorf("Value = %v, want %v", opt.Value, tt.wantValue)
			}
			if opt.Valid != tt.wantValid {
				t.Errorf("Valid = %v, want %v", opt.Valid, tt.wantValid)
			}
		})
	}
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
			if err != nil {
				t.Errorf("Unmarshal error = %v", err)
				return
			}
			if opt.Value != tt.wantValue {
				t.Errorf("Value = %v, want %v", opt.Value, tt.wantValue)
			}
			if opt.Valid != tt.wantValid {
				t.Errorf("Valid = %v, want %v", opt.Valid, tt.wantValid)
			}
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
			if err != nil {
				t.Errorf("Unmarshal error = %v", err)
				return
			}
			if opt.Value != tt.wantValue {
				t.Errorf("Value = %v, want %v", opt.Value, tt.wantValue)
			}
			if opt.Valid != tt.wantValid {
				t.Errorf("Valid = %v, want %v", opt.Valid, tt.wantValid)
			}
		})
	}
}
