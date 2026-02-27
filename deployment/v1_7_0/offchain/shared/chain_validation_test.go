package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateNOPChainSupport_AllChainsSupported_ReturnsNil(t *testing.T) {
	result := ValidateNOPChainSupport(
		"nop-1",
		[]uint64{1, 2, 3},
		[]uint64{1, 2, 3, 4, 5},
	)
	assert.Nil(t, result)
}

func TestValidateNOPChainSupport_EmptyRequiredChains_ReturnsNil(t *testing.T) {
	result := ValidateNOPChainSupport(
		"nop-1",
		[]uint64{},
		[]uint64{1, 2, 3},
	)
	assert.Nil(t, result)
}

func TestValidateNOPChainSupport_MissingChains_ReturnsResult(t *testing.T) {
	result := ValidateNOPChainSupport(
		"nop-1",
		[]uint64{1, 2, 3, 4},
		[]uint64{1, 3},
	)
	require.NotNil(t, result)
	assert.Equal(t, "nop-1", result.NOPAlias)
	assert.ElementsMatch(t, []uint64{2, 4}, result.MissingChains)
}

func TestValidateNOPChainSupport_AllChainsMissing_ReturnsAllInResult(t *testing.T) {
	result := ValidateNOPChainSupport(
		"nop-1",
		[]uint64{1, 2, 3},
		[]uint64{},
	)
	require.NotNil(t, result)
	assert.Equal(t, "nop-1", result.NOPAlias)
	assert.ElementsMatch(t, []uint64{1, 2, 3}, result.MissingChains)
}

func TestValidateNOPChainSupport_NilSupportedChains_ReturnsAllMissing(t *testing.T) {
	result := ValidateNOPChainSupport(
		"nop-1",
		[]uint64{1, 2},
		nil,
	)
	require.NotNil(t, result)
	assert.Equal(t, "nop-1", result.NOPAlias)
	assert.ElementsMatch(t, []uint64{1, 2}, result.MissingChains)
}

func TestFormatChainValidationError_EmptyResults_ReturnsNil(t *testing.T) {
	err := FormatChainValidationError([]ChainValidationResult{})
	assert.Nil(t, err)
}

func TestFormatChainValidationError_NilResults_ReturnsNil(t *testing.T) {
	err := FormatChainValidationError(nil)
	assert.Nil(t, err)
}

func TestFormatChainValidationError_SingleResult_FormatsCorrectly(t *testing.T) {
	results := []ChainValidationResult{
		{
			NOPAlias:      "nop-1",
			MissingChains: []uint64{16015286601757825753},
		},
	}
	err := FormatChainValidationError(results)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain support validation failed")
	assert.Contains(t, err.Error(), "nop-1")
	assert.Contains(t, err.Error(), "16015286601757825753")
	assert.Contains(t, err.Error(), "action: ensure configs are created for the missing chains on the node")
}

func TestFormatChainValidationError_MultipleResults_FormatsAllNOPs(t *testing.T) {
	results := []ChainValidationResult{
		{
			NOPAlias:      "nop-1",
			MissingChains: []uint64{1},
		},
		{
			NOPAlias:      "nop-2",
			MissingChains: []uint64{2, 3},
		},
	}
	err := FormatChainValidationError(results)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nop-1")
	assert.Contains(t, err.Error(), "nop-2")
}
