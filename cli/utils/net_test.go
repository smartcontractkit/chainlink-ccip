package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckHostnameResolution(t *testing.T) {
	t.Parallel()

	nsTimeout := 1 * time.Second
	retryInterval := 100 * time.Millisecond
	tests := []struct {
		name           string
		mockLookupHost func(string) ([]string, error)
		expectedError  string
	}{
		{
			name: "Success",
			mockLookupHost: func(host string) ([]string, error) {
				return []string{"127.0.0.1"}, nil
			},
			expectedError: "",
		},
		{
			name: "Failure",
			mockLookupHost: func(host string) ([]string, error) {
				return nil, fmt.Errorf("host not found")
			},
			expectedError: "DNS lookup failed after 1 seconds for example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			startTime := time.Now()
			_, err := CheckHostnameResolution("example.com", nsTimeout, retryInterval, tt.mockLookupHost)
			duration := time.Since(startTime)

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
				if tt.name == "Failure" {
					assert.GreaterOrEqual(t, duration, nsTimeout)
				}
			}
		})
	}
}

func TestAddToEtcHosts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		entry          string
		initialContent string
		expectedExists bool
		expectedError  string
	}{
		{
			name:           "Entry already exists",
			entry:          "127.0.0.1 example.com",
			initialContent: "127.0.0.1 example.com\n",
			expectedExists: true,
			expectedError:  "",
		},
		{
			name:           "Entry does not exist",
			entry:          "127.0.0.1 example.com",
			initialContent: "",
			expectedExists: false,
			expectedError:  "",
		},
		{
			name:           "Error reading hosts file",
			entry:          "127.0.0.1 example.com",
			initialContent: "",
			expectedExists: false,
			expectedError:  "no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmpHostsFilePath := filepath.Join(t.TempDir(), "hosts")
			if tt.name != "Error reading hosts file" {
				err := os.WriteFile(tmpHostsFilePath, []byte(tt.initialContent), 0o600)
				require.NoError(t, err)
			}
			exists, err := AddToEtcHosts(tt.entry, tmpHostsFilePath)

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			}
			assert.Equal(t, tt.expectedExists, exists)

			// Verify the entry was added if it didn't exist
			if !tt.expectedExists && tt.expectedError == "" {
				content, err := os.ReadFile(tmpHostsFilePath)
				assert.NoError(t, err)
				assert.Contains(t, string(content), tt.entry)
			}
		})
	}
}
