package addressbook

import (
	"errors"
	"sync"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestDON_GetOnRampAddress(t *testing.T) {
	sourceChain := ccipocr3.ChainSelector(1)
	destChain := ccipocr3.ChainSelector(2)
	addr := ccipocr3.UnknownAddress{0x1, 0x2, 0x3}

	tests := []struct {
		name        string
		state       ContractAddresses
		want        []byte
		wantErr     error
		sourceChain ccipocr3.ChainSelector
	}{
		{
			name:    "not initialized",
			state:   ContractAddresses{},
			wantErr: ErrAddressBookNotInitialized,
		},
		{
			name: "contract not registered",
			state: ContractAddresses{
				"other": {sourceChain: addr},
			},
			wantErr: ErrContractNotRegistered,
		},
		{
			name: "address not found",
			state: ContractAddresses{
				consts.ContractNameOnRamp: {},
			},
			wantErr: ErrContractAddressNotFound,
		},
		{
			name: "address empty",
			state: ContractAddresses{
				consts.ContractNameOnRamp: {sourceChain: ccipocr3.UnknownAddress{}},
			},
			sourceChain: sourceChain,
			wantErr:     ErrContractAddressEmpty,
		},
		{
			name: "success",
			state: ContractAddresses{
				consts.ContractNameOnRamp: {sourceChain: addr},
			},
			want:        addr,
			wantErr:     nil,
			sourceChain: sourceChain,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			don := NewDon(destChain)
			err := don.SetState(tt.state)
			if err != nil {
				t.Fatal(err)
			}

			got, err := don.GetOnRampAddress(tt.sourceChain)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
			if tt.wantErr == nil && string(got) != string(tt.want) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestDON_GetOffRampAddress(t *testing.T) {
	destChain := ccipocr3.ChainSelector(2)
	addr := ccipocr3.UnknownAddress{0x4, 0x5, 0x6}

	tests := []struct {
		name    string
		state   ContractAddresses
		want    []byte
		wantErr error
	}{
		{
			name:    "not initialized",
			state:   ContractAddresses{},
			wantErr: ErrAddressBookNotInitialized,
		},
		{
			name: "contract not registered",
			state: ContractAddresses{
				"other": {destChain: addr},
			},
			wantErr: ErrContractNotRegistered,
		},
		{
			name: "address not found",
			state: ContractAddresses{
				consts.ContractNameOffRamp: {},
			},
			wantErr: ErrContractAddressNotFound,
		},
		{
			name: "address empty",
			state: ContractAddresses{
				consts.ContractNameOffRamp: {destChain: ccipocr3.UnknownAddress{}},
			},
			wantErr: ErrContractAddressEmpty,
		},
		{
			name: "success",
			state: ContractAddresses{
				consts.ContractNameOffRamp: {destChain: addr},
			},
			want:    addr,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			don := NewDon(destChain)
			err := don.SetState(tt.state)
			if err != nil {
				t.Fatal(err)
			}

			got, err := don.GetOffRampAddress()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v", tt.wantErr, err)
			}
			if tt.wantErr == nil && string(got) != string(tt.want) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestDON_SetState_Concurrent(t *testing.T) {
	destChain := ccipocr3.ChainSelector(2)
	don := NewDon(destChain)

	var wg sync.WaitGroup
	numRoutines := 10

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			state := ContractAddresses{
				consts.ContractNameOnRamp: {
					ccipocr3.ChainSelector(i): ccipocr3.UnknownAddress{byte(i)},
				},
			}
			_ = don.SetState(state)
		}(i)
	}
	wg.Wait()

	// After concurrent writes, last state should be valid for at least one value
	_, err := don.GetOnRampAddress(ccipocr3.ChainSelector(numRoutines - 1))
	if err != nil && !errors.Is(err, ErrContractAddressNotFound) {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestDON_GetOnRampAddress_Concurrent(t *testing.T) {
	sourceChain := ccipocr3.ChainSelector(1)
	destChain := ccipocr3.ChainSelector(2)
	addr := ccipocr3.UnknownAddress{0x1, 0x2, 0x3}
	state := ContractAddresses{
		consts.ContractNameOnRamp: {sourceChain: addr},
	}
	don := NewDon(destChain)
	_ = don.SetState(state)

	var wg sync.WaitGroup
	numRoutines := 20
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			got, err := don.GetOnRampAddress(sourceChain)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if string(got) != string(addr) {
				t.Errorf("expected %v, got %v", addr, got)
			}
		}()
	}
	wg.Wait()
}
