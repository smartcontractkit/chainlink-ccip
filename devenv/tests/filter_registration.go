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
	// The plugins register their source-chain filters after their first OCR round, which
	// can be several minutes after the nodes boot.
	filterRegistrationTimeout = 5 * time.Minute
	filterRegistrationPoll    = time.Second
)

var tonCCIPFilterNames = []string{
	"CCIPMessageSent",
	"CommitReportAccepted",
	"ExecutionStateChanged",
}

// sourceFilterQueries counts, per chain family, the logpoller filters a node has
// registered for messages sent from one chain. $1 is the chain ID.
var sourceFilterQueries = map[string]string{
	chainsel.FamilyEVM: `
		SELECT COUNT(*)
		FROM evm.log_poller_filters
		WHERE evm_chain_id::text = $1 AND name LIKE 'OnRamp.%'`,
	chainsel.FamilySolana: `
		SELECT COUNT(*)
		FROM solana.log_poller_filters
		WHERE chain_id = $1 AND event_name = 'CCIPMessageSent' AND NOT is_deleted`,
}

// WaitForCCIPFilterRegistration blocks until every worker node listens for messages
// sent from every chain. A message sent before that can be missed for good: the
// commit plugin registers its source-chain filters only after its first OCR round,
// and it does not look far back when it does.
func WaitForCCIPFilterRegistration(t *testing.T, nodes []*clnode.Output, selectors []uint64) {
	t.Helper()

	tonSelectors := make([]uint64, 0)
	sourceChains := make(map[string][]string) // family -> chain IDs
	for _, selector := range selectors {
		family, err := chainsel.GetSelectorFamily(selector)
		require.NoError(t, err)
		if family == chainsel.FamilyTon {
			tonSelectors = append(tonSelectors, selector)
			continue
		}
		if _, ok := sourceFilterQueries[family]; !ok {
			continue
		}
		chainID, err := chainsel.GetChainIDFromSelector(selector)
		require.NoError(t, err)
		sourceChains[family] = append(sourceChains[family], chainID)
	}
	if len(tonSelectors) > 0 {
		waitForTONCCIPFilterRegistration(t, nodes, tonSelectors)
	}
	if len(sourceChains) > 0 {
		waitForSourceFilterRegistration(t, nodes, sourceChains)
	}
}

func waitForSourceFilterRegistration(t *testing.T, nodes []*clnode.Output, sourceChains map[string][]string) {
	t.Helper()

	workerDBURLs := workerNodeDBURLs(nodes)
	require.NotEmpty(t, workerDBURLs, "CCIP filter gate requires at least one worker node database")

	t.Logf("waiting for CCIP source filter registration for %v on %d worker nodes", sourceChains, len(workerDBURLs))
	require.Eventually(t, func() bool {
		for _, dbURL := range workerDBURLs {
			registered, err := sourceFiltersRegistered(t.Context(), dbURL, sourceChains)
			if err != nil || !registered {
				return false
			}
		}
		return true
	}, filterRegistrationTimeout, filterRegistrationPoll,
		"CCIP source filters were not registered for every chain on every worker")
}

func sourceFiltersRegistered(ctx context.Context, dbURL string, sourceChains map[string][]string) (bool, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return false, fmt.Errorf("open node database: %w", err)
	}
	defer db.Close()

	for family, chainIDs := range sourceChains {
		for _, chainID := range chainIDs {
			var filterCount int
			if err = db.QueryRowContext(ctx, sourceFilterQueries[family], chainID).Scan(&filterCount); err != nil {
				return false, fmt.Errorf("query %s logpoller filters of chain %s: %w", family, chainID, err)
			}
			if filterCount == 0 {
				return false, nil
			}
		}
	}
	return true, nil
}

// workerNodeDBURLs returns the database URL of every node except the bootstrap one.
func workerNodeDBURLs(nodes []*clnode.Output) []string {
	urls := make([]string, 0, len(nodes))
	for i, node := range nodes {
		// Node 0 is the bootstrap node.
		if i == 0 || node == nil || node.PostgreSQL == nil || node.PostgreSQL.Url == "" {
			continue
		}
		urls = append(urls, node.PostgreSQL.Url)
	}
	return urls
}

func waitForTONCCIPFilterRegistration(t *testing.T, nodes []*clnode.Output, selectors []uint64) {
	t.Helper()
	expectedFilters := len(selectors)

	workerDBURLs := workerNodeDBURLs(nodes)
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
