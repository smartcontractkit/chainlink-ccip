package utils

import (
	"testing"
)

func TestIsCustomImage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		devspaceNamespace string
		devspaceImage     string
		expected          bool
	}{
		{
			name:              "Valid crib-local image",
			devspaceNamespace: "crib-local",
			devspaceImage:     "localhost:5001/chainlink-node-devspace:latest",
			expected:          true,
		},
		{
			name:              "Invalid crib-local image",
			devspaceNamespace: "crib-local",
			devspaceImage:     "localhost:5001/chainlink-node:latest",
			expected:          false,
		},
		{
			name:              "Valid AWS ECR image",
			devspaceNamespace: "crib-test",
			devspaceImage:     "323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-node-devspace:latest",
			expected:          true,
		},
		{
			name:              "Invalid AWS ECR image",
			devspaceNamespace: "crib-anothertest",
			devspaceImage:     "323150190480.dkr.ecr.us-west-2.amazonaws.com/chainlink-node:latest",
			expected:          false,
		},
		{
			name:              "Non-matching namespace and image",
			devspaceNamespace: "crib-other",
			devspaceImage:     "localhost:5001/chainlink-node-devspace:latest",
			expected:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := IsCustomImage(tt.devspaceNamespace, tt.devspaceImage)
			if result != tt.expected {
				t.Errorf("IsCustomImage(%s, %s) = %v; want %v", tt.devspaceNamespace, tt.devspaceImage, result, tt.expected)
			}
		})
	}
}
