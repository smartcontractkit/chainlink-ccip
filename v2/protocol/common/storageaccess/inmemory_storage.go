package storageaccess

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// StorageEntry represents an internal storage entry for organizing CCV data with timestamp ordering.
// This encapsulates the metadata needed for timestamp-based querying
// while keeping the storage implementation details clean and maintainable.
type StorageEntry struct {
	CreatedAt      int64          // Microsecond timestamp when stored
	InsertionOrder int64          // Sequential order for deterministic sorting
	CCVData        common.CCVData // The actual CCV data
}

// Less enables sorting by (CreatedAt, InsertionOrder)
func (e StorageEntry) Less(other StorageEntry) bool {
	if e.CreatedAt != other.CreatedAt {
		return e.CreatedAt < other.CreatedAt
	}
	return e.InsertionOrder < other.InsertionOrder
}

// TimeProvider is a function type for providing current time (for testing)
type TimeProvider func() int64

// DefaultTimeProvider returns current time in microseconds since epoch
func DefaultTimeProvider() int64 {
	return time.Now().UnixMicro()
}

// InMemoryOffchainStorage implements both OffchainStorageWriter and OffchainStorageReader
// for testing and development. Uses a list of StorageEntry objects to store CCV data
// with timestamp ordering. Data is kept sorted by (CreatedAt, InsertionOrder) for efficient
// timestamp-based queries with deterministic ordering.
type InMemoryOffchainStorage struct {
	mu               sync.RWMutex
	storage          []StorageEntry // All entries sorted by (CreatedAt, InsertionOrder)
	insertionCounter int64          // Counter for deterministic ordering
	timeProvider     TimeProvider   // For providing timestamps (testable)
	lggr             logger.Logger

	// Channel to notify when data is stored (for testing)
	storedCh chan struct{}
}

// NewInMemoryOffchainStorage creates a new in-memory offchain storage
func NewInMemoryOffchainStorage(lggr logger.Logger) *InMemoryOffchainStorage {
	return NewInMemoryOffchainStorageWithTimeProvider(lggr, DefaultTimeProvider)
}

// NewInMemoryOffchainStorageWithTimeProvider creates a new in-memory offchain storage with custom time provider
func NewInMemoryOffchainStorageWithTimeProvider(lggr logger.Logger, timeProvider TimeProvider) *InMemoryOffchainStorage {
	return &InMemoryOffchainStorage{
		storage:          make([]StorageEntry, 0),
		insertionCounter: 0,
		timeProvider:     timeProvider,
		lggr:             lggr,
		storedCh:         make(chan struct{}, 100),
	}
}

// CreateReaderOnly creates a read-only view of the storage that only implements OffchainStorageReader
func CreateReaderOnly(storage *InMemoryOffchainStorage) common.OffchainStorageReader {
	return &ReaderOnlyView{storage: storage}
}

// CreateWriterOnly creates a write-only view of the storage that only implements OffchainStorageWriter
func CreateWriterOnly(storage *InMemoryOffchainStorage) common.OffchainStorageWriter {
	return &WriterOnlyView{storage: storage}
}

// WaitForStore waits for data to be stored or context to be cancelled
func (s *InMemoryOffchainStorage) WaitForStore(ctx context.Context) error {
	select {
	case <-s.storedCh:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StoreCCVData stores multiple CCV data entries in the offchain storage
func (s *InMemoryOffchainStorage) StoreCCVData(ctx context.Context, ccvDataList []common.CCVData) error {
	if len(ccvDataList) == 0 {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	createdAt := s.timeProvider() // Storage determines when data was stored
	s.lggr.Debugw("Storing CCV data",
		"count", len(ccvDataList),
		"createdAt", createdAt,
	)

	for _, ccvData := range ccvDataList {
		// Create storage entry with timestamp and ordering
		entry := StorageEntry{
			CreatedAt:      createdAt,
			InsertionOrder: s.insertionCounter,
			CCVData:        ccvData,
		}
		s.storage = append(s.storage, entry)
		s.insertionCounter++
	}

	// Keep storage sorted by (CreatedAt, InsertionOrder)
	sort.Slice(s.storage, func(i, j int) bool {
		return s.storage[i].Less(s.storage[j])
	})

	s.lggr.Debugw("Stored CCV data",
		"count", len(ccvDataList),
		"totalStored", len(s.storage),
	)

	// Notify that data was stored
	select {
	case s.storedCh <- struct{}{}:
	default:
		// Channel full, ignore
	}

	return nil
}

// GetCCVDataByTimestamp queries CCV data by timestamp with offset-based pagination
func (s *InMemoryOffchainStorage) GetCCVDataByTimestamp(
	ctx context.Context,
	destChainSelectors []cciptypes.ChainSelector,
	startTimestamp int64,
	sourceChainSelectors []cciptypes.ChainSelector,
	limit int,
	offset int,
) (*common.TimestampQueryResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.lggr.Debugw("Querying CCV data by timestamp",
		"destChains", destChainSelectors,
		"startTimestamp", startTimestamp,
		"sourceChains", sourceChainSelectors,
		"limit", limit,
		"offset", offset,
		"totalEntries", len(s.storage),
	)

	// Filter data by timestamp and other criteria
	var filteredEntries []StorageEntry

	for _, entry := range s.storage {
		// Skip entries before start_timestamp
		if entry.CreatedAt < startTimestamp {
			continue
		}

		// Filter by destination chain
		found := false
		for _, destChain := range destChainSelectors {
			if entry.CCVData.DestChainSelector == destChain {
				found = true
				break
			}
		}
		if !found {
			continue
		}

		// Filter by source chain if specified
		if len(sourceChainSelectors) > 0 {
			found = false
			for _, sourceChain := range sourceChainSelectors {
				if entry.CCVData.SourceChainSelector == sourceChain {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		filteredEntries = append(filteredEntries, entry)
	}

	// Apply offset and limit
	totalFiltered := len(filteredEntries)
	var paginatedEntries []StorageEntry
	if offset < totalFiltered {
		end := offset + limit
		if end > totalFiltered {
			end = totalFiltered
		}
		paginatedEntries = filteredEntries[offset:end]
	}

	// Organize results by destination chain
	resultData := make(map[cciptypes.ChainSelector][]common.CCVData)
	for _, entry := range paginatedEntries {
		destChain := entry.CCVData.DestChainSelector
		resultData[destChain] = append(resultData[destChain], entry.CCVData)
	}

	// Determine pagination metadata
	hasMore := (offset + limit) < totalFiltered
	var nextTimestamp *int64

	if len(paginatedEntries) > 0 && !hasMore {
		// If this is the last page, find the next timestamp for future queries
		lastEntryTimestamp := paginatedEntries[len(paginatedEntries)-1].CreatedAt
		for _, entry := range s.storage {
			if entry.CreatedAt > lastEntryTimestamp {
				nextTimestamp = &entry.CreatedAt
				break
			}
		}
	} else if len(paginatedEntries) > 0 && hasMore {
		// More data exists at current timestamp, keep same timestamp for next query
		nextTimestamp = &startTimestamp
	}

	response := &common.TimestampQueryResponse{
		Data:          resultData,
		NextTimestamp: nextTimestamp,
		HasMore:       hasMore,
		TotalCount:    len(paginatedEntries),
	}

	return response, nil
}

// GetAllCCVData retrieves all CCV data for a verifier (for testing/debugging)
func (s *InMemoryOffchainStorage) GetAllCCVData(verifierAddress []byte) ([]common.CCVData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []common.CCVData
	for _, entry := range s.storage {
		if string(entry.CCVData.SourceVerifierAddress) == string(verifierAddress) {
			result = append(result, entry.CCVData)
		}
	}

	return result, nil
}

// GetCCVDataByMessageID retrieves CCV data by message ID (for testing/debugging)
func (s *InMemoryOffchainStorage) GetCCVDataByMessageID(messageID cciptypes.Bytes32) (*common.CCVData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, entry := range s.storage {
		if string(entry.CCVData.MessageID[:]) == string(messageID[:]) {
			return &entry.CCVData, nil
		}
	}

	return nil, fmt.Errorf("CCV data not found for message ID: %x", messageID)
}

// GetStats returns storage statistics (for testing/debugging)
func (s *InMemoryOffchainStorage) GetStats() map[string]interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stats := make(map[string]interface{})
	totalEntries := len(s.storage)
	verifierCounts := make(map[string]int)

	for _, entry := range s.storage {
		verifierKey := string(entry.CCVData.SourceVerifierAddress)
		verifierCounts[verifierKey]++
	}

	stats["totalEntries"] = totalEntries
	stats["verifierCounts"] = verifierCounts
	stats["verifierCount"] = len(verifierCounts)

	return stats
}

// Clear removes all stored data (for testing)
func (s *InMemoryOffchainStorage) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage = make([]StorageEntry, 0)
	s.insertionCounter = 0
	s.lggr.Debugw("Cleared all stored data")
}

// GetTotalCount returns the total number of CCV data entries stored (for testing)
func (s *InMemoryOffchainStorage) GetTotalCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.storage)
}

// Helper methods for testing and debugging (not part of the interface)
func (s *InMemoryOffchainStorage) ListDestChains() []cciptypes.ChainSelector {
	s.mu.RLock()
	defer s.mu.RUnlock()

	destChains := make(map[cciptypes.ChainSelector]bool)
	for _, entry := range s.storage {
		destChains[entry.CCVData.DestChainSelector] = true
	}

	var result []cciptypes.ChainSelector
	for destChain := range destChains {
		result = append(result, destChain)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func (s *InMemoryOffchainStorage) ListSourceChains(destChainSelector cciptypes.ChainSelector) []cciptypes.ChainSelector {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sourceChains := make(map[cciptypes.ChainSelector]bool)
	for _, entry := range s.storage {
		if entry.CCVData.DestChainSelector == destChainSelector {
			sourceChains[entry.CCVData.SourceChainSelector] = true
		}
	}

	var result []cciptypes.ChainSelector
	for sourceChain := range sourceChains {
		result = append(result, sourceChain)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func (s *InMemoryOffchainStorage) ListVerifierAddresses(destChainSelector, sourceChainSelector cciptypes.ChainSelector) [][]byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	verifierAddresses := make(map[string][]byte)
	for _, entry := range s.storage {
		ccvData := entry.CCVData
		if ccvData.DestChainSelector == destChainSelector && ccvData.SourceChainSelector == sourceChainSelector {
			key := string(ccvData.SourceVerifierAddress)
			verifierAddresses[key] = ccvData.SourceVerifierAddress
		}
	}

	var result [][]byte
	for _, addr := range verifierAddresses {
		result = append(result, addr)
	}

	sort.Slice(result, func(i, j int) bool {
		return string(result[i]) < string(result[j])
	})

	return result
}

// ReaderOnlyView provides a read-only view of the storage that only implements OffchainStorageReader
type ReaderOnlyView struct {
	storage *InMemoryOffchainStorage
}

func (r *ReaderOnlyView) GetCCVDataByTimestamp(
	ctx context.Context,
	destChainSelectors []cciptypes.ChainSelector,
	startTimestamp int64,
	sourceChainSelectors []cciptypes.ChainSelector,
	limit int,
	offset int,
) (*common.TimestampQueryResponse, error) {
	return r.storage.GetCCVDataByTimestamp(ctx, destChainSelectors, startTimestamp, sourceChainSelectors, limit, offset)
}

// Utility methods for testing (not part of the interface)
func (r *ReaderOnlyView) GetTotalCount() int {
	return r.storage.GetTotalCount()
}

// WriterOnlyView provides a write-only view of the storage that only implements OffchainStorageWriter
type WriterOnlyView struct {
	storage *InMemoryOffchainStorage
}

func (w *WriterOnlyView) StoreCCVData(ctx context.Context, ccvDataList []common.CCVData) error {
	return w.storage.StoreCCVData(ctx, ccvDataList)
}
