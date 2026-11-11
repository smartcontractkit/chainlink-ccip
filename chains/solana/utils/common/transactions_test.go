package common

import (
	"testing"

	"github.com/gagliardetto/solana-go/rpc"
)

func TestCommitmentReached(t *testing.T) {
	tests := []struct {
		name   string
		status rpc.ConfirmationStatusType
		wanted rpc.ConfirmationStatusType
		want   bool
	}{
		{"nothing reported yet", "", rpc.ConfirmationStatusConfirmed, false},
		{"processed does not reach confirmed", rpc.ConfirmationStatusProcessed, rpc.ConfirmationStatusConfirmed, false},
		{"confirmed reaches confirmed", rpc.ConfirmationStatusConfirmed, rpc.ConfirmationStatusConfirmed, true},
		{"finalized reaches confirmed, even if confirmed was never reported", rpc.ConfirmationStatusFinalized, rpc.ConfirmationStatusConfirmed, true},
		{"confirmed does not reach finalized", rpc.ConfirmationStatusConfirmed, rpc.ConfirmationStatusFinalized, false},
		{"finalized reaches finalized", rpc.ConfirmationStatusFinalized, rpc.ConfirmationStatusFinalized, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commitmentReached(tt.status, tt.wanted); got != tt.want {
				t.Fatalf("commitmentReached(%q, %q) = %v, want %v", tt.status, tt.wanted, got, tt.want)
			}
		})
	}
}
