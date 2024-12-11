package peergroup

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	// pgmocks "github.com/smartcontractkit/chainlink-ccip/mocks/commit/merkleroot/rmn"
	"github.com/smartcontractkit/libocr/networking"

	pgfactorymocks "github.com/smartcontractkit/chainlink-ccip/mocks/libocr_networking"
)

type mockPeerGroup struct {
	mock.Mock
}

func (m *mockPeerGroup) NewStream(
	remotePeerID string,
	newStreamArgs networking.NewStreamArgs) (networking.Stream, error) {
	args := m.Called(remotePeerID, newStreamArgs)
	return args.Get(0).(networking.Stream), args.Error(1)
}

func (m *mockPeerGroup) Close() error {
	return nil
}

func Test_writePrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		prefix    ocr2types.ConfigDigestPrefix
		hash      cciptypes.Bytes32
		wantFirst byte
		wantNext  byte
	}{
		{
			name:      "zero hash",
			prefix:    ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo,
			hash:      cciptypes.Bytes32{},
			wantFirst: byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo >> 8),
			wantNext:  byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo),
		},
		{
			name:      "non-zero hash",
			prefix:    ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo,
			hash:      cciptypes.Bytes32{0xFF, 0xFF, 0x3},
			wantFirst: byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo >> 8),
			wantNext:  byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := writePrefix(tt.prefix, tt.hash)

			assert.Equal(t, tt.wantFirst, result[0], "first byte should match prefix")
			assert.Equal(t, tt.wantNext, result[1], "second byte should match prefix")

			// Rest of hash should match input except first two bytes
			for i := 2; i < len(tt.hash); i++ {
				assert.Equal(t, tt.hash[i], result[i], "remaining bytes should be unchanged")
			}
		})
	}
}

func TestCreator_Create(t *testing.T) {
	t.Parallel()

	oracle1 := peerIDFromStr(t, "1")
	oracle2 := peerIDFromStr(t, "2")
	rmn1 := peerIDFromStr(t, "3")
	rmn2 := peerIDFromStr(t, "4")

	bootstrappers := []commontypes.BootstrapperLocator{{PeerID: "bootstrap1"}}
	mockPG := &mockPeerGroup{}

	tests := []struct {
		name          string
		opts          CreateOpts
		expectedPeers []string
		setupMocks    func(*pgfactorymocks.MockPeerGroupFactory)
		wantErr       bool
	}{
		{
			name: "success with both oracle and RMN peers",
			opts: CreateOpts{
				CommitConfigDigest:  cciptypes.Bytes32{0x1},
				RMNHomeConfigDigest: cciptypes.Bytes32{0x2},
				OraclePeerIDs: []ragep2ptypes.PeerID{
					oracle1,
					oracle2,
				},
				RMNNodes: []rmntypes.HomeNodeInfo{
					{ID: 1, PeerID: rmn1},
					{ID: 2, PeerID: rmn2},
				},
			},
			expectedPeers: []string{"oracle1", "oracle2", "rmn1", "rmn2"},
			setupMocks: func(m *pgfactorymocks.MockPeerGroupFactory) {
				m.EXPECT().NewPeerGroup(
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(mockPG, nil)
			},
		},
		{
			name: "success with only oracle peers",
			opts: CreateOpts{
				CommitConfigDigest:  cciptypes.Bytes32{0x1},
				RMNHomeConfigDigest: cciptypes.Bytes32{0x2},
				OraclePeerIDs: []ragep2ptypes.PeerID{
					oracle1,
				},
			},
			expectedPeers: []string{"oracle1"},
			setupMocks: func(m *pgfactorymocks.MockPeerGroupFactory) {
				m.EXPECT().NewPeerGroup(
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(mockPG, nil)
			},
		},
		{
			name: "success with empty peer lists",
			opts: CreateOpts{
				CommitConfigDigest:  cciptypes.Bytes32{0x1},
				RMNHomeConfigDigest: cciptypes.Bytes32{0x2},
			},
			expectedPeers: []string{},
			setupMocks: func(m *pgfactorymocks.MockPeerGroupFactory) {
				m.EXPECT().NewPeerGroup(
					mock.MatchedBy(func(digest ocr2types.ConfigDigest) bool {
						return digest[0] == byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo>>8) &&
							digest[1] == byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo)
					}),
					[]string{},
					bootstrappers,
				).Return(mockPG, nil)
			},
		},
		{
			name: "factory returns error",
			opts: CreateOpts{
				CommitConfigDigest:  cciptypes.Bytes32{0x1},
				RMNHomeConfigDigest: cciptypes.Bytes32{0x2},
			},
			setupMocks: func(m *pgfactorymocks.MockPeerGroupFactory) {
				m.EXPECT().NewPeerGroup(
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil, assert.AnError)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			factory := pgfactorymocks.NewMockPeerGroupFactory(t)
			tt.setupMocks(factory)

			creator := NewCreator(logger.Nop(), factory, bootstrappers)

			result, err := creator.Create(tt.opts)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.NotNil(t, result.PeerGroup)
			assert.NotEmpty(t, result.ConfigDigest)

			// Config digest should have correct prefix
			assert.Equal(t,
				byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo>>8),
				result.ConfigDigest[0],
				"config digest should have correct prefix first byte",
			)
			assert.Equal(t,
				byte(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo),
				result.ConfigDigest[1],
				"config digest should have correct prefix second byte",
			)
		})
	}
}

func TestNewCreator(t *testing.T) {
	t.Parallel()

	lggr := logger.Nop()
	factory := pgfactorymocks.NewMockPeerGroupFactory(t)
	bootstrappers := []commontypes.BootstrapperLocator{{PeerID: "bootstrap1"}}

	creator := NewCreator(lggr, factory, bootstrappers)

	assert.NotNil(t, creator)
	assert.Equal(t, factory, creator.factory)
	assert.Equal(t, bootstrappers, creator.bootstrappers)
	assert.NotNil(t, creator.lggr)
}

func peerIDFromStr(t *testing.T, s string) ragep2ptypes.PeerID {
	t.Helper()
	// Pad with zeros to make it 32 bytes
	padded := fmt.Sprintf("%064s", s)
	bytes, err := hex.DecodeString(padded)
	require.NoError(t, err)
	var result [32]byte
	copy(result[:], bytes)
	return ragep2ptypes.PeerID(result)
}
