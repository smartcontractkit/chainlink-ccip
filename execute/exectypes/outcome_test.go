package exectypes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPluginState_Next(t *testing.T) {
	tests := []struct {
		name    string
		p       PluginState
		want    PluginState
		isPanic bool
	}{
		{
			name: "Zero value",
			p:    Unknown,
			want: GetCommitReports,
		},
		{
			name: "Initialized",
			p:    Initialized,
			want: GetCommitReports,
		},
		{
			name: "Phase 1 to 2",
			p:    GetCommitReports,
			want: GetMessages,
		},
		{
			name: "Phase 2 to 3",
			p:    GetMessages,
			want: Filter,
		},
		{
			name: "Phase 3 to 1",
			p:    Filter,
			want: GetCommitReports,
		},
		{
			name:    "panic",
			p:       PluginState("ElToroLoco"),
			isPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.isPanic {
				require.Panics(t, func() {
					tt.p.Next()
				})
				return
			}

			if got := tt.p.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
