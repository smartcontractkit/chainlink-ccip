package devspace

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test splitYAML function
func TestSplitYAML(t *testing.T) {
	t.Parallel()
	input := `apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
---
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
`

	expected := [][]byte{
		[]byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config`),
		[]byte(`apiVersion: v1
kind: Secret
metadata:
  name: my-secret
`),
	}

	result := splitYAML([]byte(input))

	assert.Equal(t, len(expected), len(result), "Expected number of manifests to match")
	for i, exp := range expected {
		assert.Equal(t, string(exp), string(result[i]), "Manifest content should match")
	}
}

// Test removeEscapedLines function
func TestRemoveEscapedLines(t *testing.T) {
	t.Parallel()

	input := "\x1b[31mThis is an escaped line" +
		"\x1b[0m\n" +
		"apiVersion: v1\n" +
		"kind: ConfigMap\n" +
		"metadata:\n" +
		"  name: my-config\n" +
		"\x1b[32mAnother escaped line" +
		"\x1b[0m\n" +
		"data:\n" +
		"  key: value\n"

	expected := `apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
data:
  key: value
`

	reader := bytes.NewReader([]byte(input))
	cleaned := removeEscapedLines(reader)

	assert.Equal(t, expected, cleaned, "Escaped lines should be removed")
}
