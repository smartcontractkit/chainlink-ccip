package ccip

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

/*
This code should be generalized and moved to devenv library after we finish CCIPv1.7 environment!
*/

type TimeTracker struct {
	logger    zerolog.Logger
	start     time.Time
	last      time.Time
	intervals []interval
}

type interval struct {
	tag   string
	delta time.Duration
}

// NewTimeTracker is a simple utility function that tracks execution time.
func NewTimeTracker(l zerolog.Logger) *TimeTracker { //nolint:gocritic
	now := time.Now()
	return &TimeTracker{
		start:     now,
		last:      now,
		logger:    l,
		intervals: make([]interval, 0),
	}
}

func (t *TimeTracker) Record(tag string) {
	now := time.Now()
	delta := now.Sub(t.last)
	t.intervals = append(t.intervals, interval{
		tag:   tag,
		delta: delta,
	})
	t.last = now
}

func (t *TimeTracker) Print() {
	total := time.Since(t.start)
	t.logger.Debug().Msg("Time tracking results:")
	for _, i := range t.intervals {
		t.logger.Debug().
			Str("Tag", i.tag).
			Str("Duration", i.delta.String()).
			Send()
	}

	t.logger.Debug().
		Str("Duration", total.String()).
		Msg("Total environment boot up time")
}

func PrintCLDFAddresses(in *Cfg) error {
	for _, addr := range in.CLDF.Addresses {
		var refs []datastore.AddressRef
		if err := json.Unmarshal([]byte(addr), &refs); err != nil {
			return fmt.Errorf("failed to unmarshal addresses: %w", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "Selector\tType\tAddress\tVersion\tQualifier")
		fmt.Fprintln(w, "--------\t----\t-------\t-------\t---------")

		for _, ref := range refs {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", ref.ChainSelector, ref.Type, ref.Address, ref.Version, ref.Qualifier)
		}
	}
	return nil
}
