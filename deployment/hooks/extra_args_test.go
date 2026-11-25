package hooks

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

// The nil embedded interface is never called.
type forkSrcWithoutFamily struct {
	testadapters.ForkCCIPSendTestAdapter
}

type forkSrcWithFamily struct {
	forkSrcWithoutFamily
}

func (forkSrcWithFamily) Family() string { return "adapter-family" }

func TestExtraArgsSource(t *testing.T) {
	t.Run("adapter with Family is used as is", func(t *testing.T) {
		src := forkSrcWithFamily{}
		assert.Equal(t, src, extraArgsSource(src, "hook-family"))
	})

	t.Run("adapter without Family falls back to the hook's family", func(t *testing.T) {
		got := extraArgsSource(forkSrcWithoutFamily{}, "hook-family")
		assert.Equal(t, familyOnlySource("hook-family"), got)
		assert.Equal(t, "hook-family", got.Family())
	})
}
