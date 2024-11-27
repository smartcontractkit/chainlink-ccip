package utils

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"
)

// LookupFunc defines the signature of the hostname lookup function, so we can mock it in tests.
type LookupFunc func(string) ([]string, error)

// CheckHostnameResolution checks if the specified hostname can be resolved by the DNS server.
// Does not use net.LookupHost because of https://github.com/golang/go/issues/67925 (dial udp 127.0.0.53:53: i/o timeout")
// It retries the resolution until the timeout (specified in the context) is reached.
func CheckHostnameResolution(hostname string, nsTimeout, interval time.Duration, lookup LookupFunc) (time.Duration, error) {
	// every attempt has a timeout of 2 seconds
	// TODO: make this timeout configurable
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if lookup == nil {
		lookup = func(string) ([]string, error) {
			if err := exec.CommandContext(ctx, "dig", hostname, "+short").Run(); err != nil {
				return nil, err
			}
			return nil, nil
		}
	}

	start := time.Now()
	for time.Since(start) < nsTimeout {
		_, err := lookup(hostname)
		if err == nil {
			return time.Since(start), nil
		}
		slog.Error("DNS lookup failed", slog.String("hostname", hostname), slog.Any("error", err))

		time.Sleep(interval)
	}
	return time.Since(start), fmt.Errorf("DNS lookup failed after %d seconds for %s", int(nsTimeout.Seconds()), hostname)
}

// AddToEtcHosts adds an entry to the /etc/hosts file.
// hostsFile is the path to the hosts file, if empty it defaults to /etc/hosts (this is useful for testing)
// Returns true if the entry already exists, false otherwise.
func AddToEtcHosts(entry string, hostsFilePath string) (bool, error) {
	if hostsFilePath == "" {
		hostsFilePath = "/etc/hosts"
	}

	b, err := os.ReadFile(hostsFilePath)
	if err != nil {
		return false, fmt.Errorf("error reading hosts file: %w", err)
	}
	hostsFileContent := string(b)

	// Check if the entry already exists
	if strings.Contains(hostsFileContent, entry) {
		return true, nil
	}

	// Try to open the file with write permissions
	sudo := ""
	file, err := os.OpenFile(hostsFilePath, os.O_RDWR, 0o66)
	if err != nil {
		sudo = "sudo"
	}
	defer file.Close()

	// Append the entry using sudo
	cmdString := fmt.Sprintf(`echo "%s" | %s tee -a %s >/dev/null`, entry, sudo, hostsFilePath)
	cmd := exec.Command("bash", "-c", cmdString)
	if err := cmd.Run(); err != nil {
		return false, fmt.Errorf("error appending entry: %w", err)
	}

	return false, nil
}
