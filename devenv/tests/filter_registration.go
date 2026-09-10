package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/clnode"
)

const (
	filterRegistrationTimeout = 2 * time.Minute
	filterRegistrationPoll    = time.Second
)

var tonCCIPFilterNames = []string{
	"CCIPMessageSent",
	"CommitReportAccepted",
	"ExecutionStateChanged",
}

// WaitForCCIPFilterRegistration currently waits for TON filters only. Extend it
// as other chain families need a pre-send registration gate.
func WaitForCCIPFilterRegistration(t *testing.T, nodes []*clnode.Output, selectors []uint64) {
	t.Helper()

	tonSelectors := make([]uint64, 0)
	for _, selector := range selectors {
		family, err := chainsel.GetSelectorFamily(selector)
		require.NoError(t, err)
		if family == chainsel.FamilyTon {
			tonSelectors = append(tonSelectors, selector)
		}
	}
	if len(tonSelectors) > 0 {
		waitForTONCCIPFilterRegistration(t, nodes, tonSelectors)
	}
}

func waitForTONCCIPFilterRegistration(t *testing.T, nodes []*clnode.Output, selectors []uint64) {
	t.Helper()
	expectedFilters := len(selectors)

	workerDBURLs := make([]string, 0, len(nodes))
	for i, node := range nodes {
		// Node 0 is the bootstrap node.
		if i == 0 || node == nil || node.PostgreSQL == nil || node.PostgreSQL.Url == "" {
			continue
		}
		workerDBURLs = append(workerDBURLs, node.PostgreSQL.Url)
	}
	require.NotEmpty(t, workerDBURLs, "TON CCIP filter gate requires at least one worker node database")

	t.Logf("waiting for TON CCIP filter registration for %d chains on %d worker nodes", expectedFilters, len(workerDBURLs))
	require.Eventually(t, func() bool {
		for _, dbURL := range workerDBURLs {
			registered, err := tonFiltersRegistered(t.Context(), dbURL, expectedFilters)
			if err != nil || !registered {
				return false
			}
		}
		return true
	}, filterRegistrationTimeout, filterRegistrationPoll,
		"TON CCIP filters were not registered for every chain on every worker")
}

func tonFiltersRegistered(ctx context.Context, dbURL string, expectedFilters int) (bool, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return false, fmt.Errorf("open node database: %w", err)
	}
	defer db.Close()

	for _, filterName := range tonCCIPFilterNames {
		var filterCount int
		err = db.QueryRowContext(ctx, `
			SELECT COUNT(*)
			FROM ton.log_poller_filters
			WHERE name = $1 AND NOT is_deleted
		`, filterName).Scan(&filterCount)
		if err != nil {
			return false, fmt.Errorf("query TON logpoller filter %s: %w", filterName, err)
		}
		if filterCount < expectedFilters {
			return false, nil
		}
	}
	return true, nil
}
