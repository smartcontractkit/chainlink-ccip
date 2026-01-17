package merkleroot

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	reader2 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func TestValidateRootBlessings(t *testing.T) {
	ctx := t.Context()
	chainA := cciptypes.ChainSelector(1)
	chainB := cciptypes.ChainSelector(2)
	blessed := []cciptypes.MerkleRootChain{{ChainSel: chainA}}
	unblessed := []cciptypes.MerkleRootChain{{ChainSel: chainB}}
	sourceChainConfig := map[cciptypes.ChainSelector]reader2.StaticSourceChainConfig{
		chainA: {IsRMNVerificationDisabled: false, IsEnabled: true},
		chainB: {IsRMNVerificationDisabled: true, IsEnabled: true},
	}

	// Helper function to copy and override config
	copyConfigWithOverride := func(chain cciptypes.ChainSelector,
		override reader2.StaticSourceChainConfig) map[cciptypes.ChainSelector]reader2.StaticSourceChainConfig {
		config := make(map[cciptypes.ChainSelector]reader2.StaticSourceChainConfig)
		for k, v := range sourceChainConfig {
			config[k] = v
		}
		config[chain] = override
		return config
	}

	reader := readermock.NewMockCCIPReader(t)
	reader.EXPECT().GetOffRampSourceChainsConfig(ctx, mock.Anything).Return(sourceChainConfig, nil).Maybe()

	err := ValidateRootBlessings(ctx, reader, blessed, unblessed)
	assert.NoError(t, err)

	t.Run("duplicate chain", func(t *testing.T) {
		dup := []cciptypes.MerkleRootChain{{ChainSel: chainA}}
		err := ValidateRootBlessings(ctx, reader, blessed, dup)
		assert.ErrorContains(t, err, "duplicate chain")
	})

	t.Run("missing chain in config", func(t *testing.T) {
		mockReader2 := readermock.NewMockCCIPReader(t)
		// Only include chainA, missing chainB
		configMissingChain := map[cciptypes.ChainSelector]reader2.StaticSourceChainConfig{
			chainA: sourceChainConfig[chainA],
		}
		mockReader2.EXPECT().GetOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB}).Return(
			configMissingChain, nil).Maybe()

		err := ValidateRootBlessings(ctx, mockReader2, blessed, unblessed)
		assert.ErrorContains(t, err, "not in the offRampSourceChainsConfig")
	})

	t.Run("RMN-disabled but blessed", func(t *testing.T) {
		mockReader3 := readermock.NewMockCCIPReader(t)
		configWithRMNDisabled := copyConfigWithOverride(chainA, reader2.StaticSourceChainConfig{
			IsRMNVerificationDisabled: true,
			IsEnabled:                 true,
		})
		mockReader3.EXPECT().GetOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB}).Return(
			configWithRMNDisabled, nil).Maybe()
		err := ValidateRootBlessings(ctx, mockReader3, blessed, unblessed)
		assert.ErrorContains(t, err, "RMN-disabled but root is blessed")
	})

	t.Run("disabled but blessed", func(t *testing.T) {
		mockReader4 := readermock.NewMockCCIPReader(t)
		configWithDisabled := copyConfigWithOverride(chainA, reader2.StaticSourceChainConfig{
			IsRMNVerificationDisabled: false,
			IsEnabled:                 false,
		})
		mockReader4.EXPECT().GetOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB}).Return(
			configWithDisabled, nil).Maybe()
		err := ValidateRootBlessings(ctx, mockReader4, blessed, unblessed)
		assert.ErrorContains(t, err, "disabled but root is blessed")
	})

	t.Run("RMN-enabled but unblessed", func(t *testing.T) {
		mockReader5 := readermock.NewMockCCIPReader(t)
		configWithRMNEnabled := copyConfigWithOverride(chainB, reader2.StaticSourceChainConfig{
			IsRMNVerificationDisabled: false,
			IsEnabled:                 true,
		})
		mockReader5.EXPECT().GetOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB}).Return(
			configWithRMNEnabled, nil).Maybe()
		err := ValidateRootBlessings(ctx, mockReader5, blessed, unblessed)
		assert.ErrorContains(t, err, "RMN-enabled but root is unblessed")
	})

	t.Run("disabled but root is reported", func(t *testing.T) {
		mockReader6 := readermock.NewMockCCIPReader(t)
		configWithDisabled := copyConfigWithOverride(chainB, reader2.StaticSourceChainConfig{
			IsRMNVerificationDisabled: true,
			IsEnabled:                 false,
		})
		mockReader6.EXPECT().GetOffRampSourceChainsConfig(ctx, []cciptypes.ChainSelector{chainA, chainB}).Return(
			configWithDisabled, nil).Maybe()
		err := ValidateRootBlessings(ctx, mockReader6, blessed, unblessed)
		assert.ErrorContains(t, err, "disabled but root is reported")
	})
}
