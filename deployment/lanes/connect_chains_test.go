package lanes

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	deploy "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// fakeLaneVersionResolver is a LaneVersionResolver whose on-chain reads are stubbed so the
// downgrade guard can be exercised without a chain.
type fakeLaneVersionResolver struct {
	supported bool
	// versions maps remote chain selector to the version currently on chain. A missing entry
	// means no lane is configured to that remote (nil version).
	versions map[uint64]*semver.Version
	err      error
}

func (r *fakeLaneVersionResolver) IsSupportedChain(cldf.Environment, uint64) bool {
	return r.supported
}

func (r *fakeLaneVersionResolver) DeriveLaneVersionsForChain(cldf.Environment, uint64) (map[uint64]*semver.Version, []*semver.Version, error) {
	return r.versions, nil, r.err
}

func (r *fakeLaneVersionResolver) LaneVersionForRemoteChain(_ cldf.Environment, _, remoteChain uint64) (*semver.Version, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.versions[remoteChain], nil
}

// fakeLaneVersionResolverProvider is a LaneVersionResolverProvider backed by a per-selector map.
type fakeLaneVersionResolverProvider struct {
	resolvers map[uint64]deploy.LaneVersionResolver
}

func (p *fakeLaneVersionResolverProvider) GetLaneVersionResolver(sel uint64) (deploy.LaneVersionResolver, bool) {
	r, ok := p.resolvers[sel]
	return r, ok
}

func TestValidateNoDowngrade(t *testing.T) {
	t.Parallel()

	const (
		chainA = uint64(5009297550715157269)
		chainB = uint64(4949039107694359620)
	)

	// errRead is a sentinel so the wrapped error can be matched with errors.Is.
	errRead := errors.New("rpc unavailable")

	lane := LaneConfig{
		ChainA:  ChainDefinition{Selector: chainA},
		ChainB:  ChainDefinition{Selector: chainB},
		Version: semver.MustParse("1.6.0"),
	}

	// providerWith returns a provider whose resolver reports the given on-chain version for the
	// lane's remote chain, for both legs of the lane.
	providerWith := func(current *semver.Version) *fakeLaneVersionResolverProvider {
		resolver := &fakeLaneVersionResolver{
			supported: true,
			versions:  map[uint64]*semver.Version{chainA: current, chainB: current},
		}
		return &fakeLaneVersionResolverProvider{
			resolvers: map[uint64]deploy.LaneVersionResolver{chainA: resolver, chainB: resolver},
		}
	}

	tests := []struct {
		name           string
		provider       LaneVersionResolverProvider
		allowDowngrade bool
		wantErr        error
	}{
		{
			name:     "Success - nil provider skips the guard",
			provider: nil,
		},
		{
			name:           "Success - allowDowngrade bypasses the guard",
			provider:       providerWith(semver.MustParse("2.0.0")),
			allowDowngrade: true,
		},
		{
			name:     "Success - upgrade is permitted",
			provider: providerWith(semver.MustParse("1.5.0")),
		},
		{
			name:     "Success - same version is permitted",
			provider: providerWith(semver.MustParse("1.6.0")),
		},
		{
			name:     "Success - no on-chain lane yet is permitted",
			provider: providerWith(nil),
		},
		{
			name: "Success - unregistered chain is skipped",
			provider: &fakeLaneVersionResolverProvider{
				resolvers: map[uint64]deploy.LaneVersionResolver{},
			},
		},
		{
			name: "Success - unsupported chain is skipped",
			provider: &fakeLaneVersionResolverProvider{
				resolvers: map[uint64]deploy.LaneVersionResolver{
					chainA: &fakeLaneVersionResolver{supported: false},
					chainB: &fakeLaneVersionResolver{supported: false},
				},
			},
		},
		{
			name:    "Failure - downgrade is refused",
			provider: providerWith(semver.MustParse("2.0.0")),
			wantErr: ErrLaneDowngrade,
		},
		{
			name: "Failure - resolver read error is surfaced",
			provider: &fakeLaneVersionResolverProvider{
				resolvers: map[uint64]deploy.LaneVersionResolver{
					chainA: &fakeLaneVersionResolver{supported: true, err: errRead},
					chainB: &fakeLaneVersionResolver{supported: true, err: errRead},
				},
			},
			wantErr: errRead,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := validateNoDowngrade(cldf.Environment{}, tc.provider, lane, tc.allowDowngrade)
			if tc.wantErr == nil {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.ErrorIs(t, err, tc.wantErr)
		})
	}
}

func TestValidateChainDefinition(t *testing.T) {
	t.Parallel()

	t.Run("valid input passes", func(t *testing.T) {
		t.Parallel()
		require.NoError(t, validateChainDefinition(ChainDefinition{
			Selector: 1,
		}))
	})

	programmaticFields := []struct {
		name string
		def  ChainDefinition
	}{
		{"OnRamp", ChainDefinition{OnRamp: []byte{0x01}}},
		{"OffRamp", ChainDefinition{OffRamp: []byte{0x01}}},
		{"Router", ChainDefinition{Router: []byte{0x01}}},
		{"FeeQuoter", ChainDefinition{FeeQuoter: []byte{0x01}}},
		{"FeeQuoterDestChainConfig", ChainDefinition{FeeQuoterDestChainConfig: FeeQuoterDestChainConfig{IsEnabled: true}}},
		{"FeeQuoterVersion", ChainDefinition{FeeQuoterVersion: semver.MustParse("1.6.0")}},
	}

	for _, tc := range programmaticFields {
		t.Run(tc.name+" is rejected", func(t *testing.T) {
			t.Parallel()
			err := validateChainDefinition(tc.def)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.name)
			assert.Contains(t, err.Error(), "must not be set by the caller")
		})
	}
}
