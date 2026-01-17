package addressbook

import (
	"bytes"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestBook_GetContractAddress(t *testing.T) {
	testCases := []struct {
		name     string
		state    ContractAddresses
		getName  ContractName
		getChain ccipocr3.ChainSelector
		expAddr  ccipocr3.UnknownAddress
		expErr   error
	}{
		{
			name:     "empty state",
			state:    nil,
			getName:  "abc",
			getChain: 1,
			expAddr:  nil,
			expErr:   ErrAddressBookNotInitialized,
		},
		{
			name: "contract not registered",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("addr1")},
			},
			getName:  "c2",
			getChain: 1,
			expAddr:  nil,
			expErr:   ErrContractNotRegistered,
		},
		{
			name: "contract address not found for that chain",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("addr1")},
			},
			getName:  "c1",
			getChain: 2,
			expAddr:  nil,
			expErr:   ErrContractAddressNotFound,
		},
		{
			name: "contract found but the address is empty",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("")},
			},
			getName:  "c1",
			getChain: 1,
			expAddr:  nil,
			expErr:   ErrContractAddressEmpty,
		},
		{
			name: "contract found",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("addr1")},
			},
			getName:  "c1",
			getChain: 1,
			expAddr:  []byte("addr1"),
			expErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ab := NewBook()
			require.NoError(t, ab.InsertOrUpdate(tc.state))

			res, err := ab.GetContractAddress(tc.getName, tc.getChain)
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expAddr, res)
		})
	}
}

func TestBook_InsertOrUpdate(t *testing.T) {
	testCases := []struct {
		name          string
		state         ContractAddresses
		appendedState ContractAddresses
		expState      ContractAddresses
	}{
		{
			name:  "initial state update",
			state: make(ContractAddresses),
			appendedState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("addr1")},
			},
			expState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"c1": {1: []byte("addr1")},
			},
		},
		{
			name: "add a new contract and a new chain under existing contract",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {1: []byte("addr1")},
			},
			appendedState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {2: []byte("addr2")}, // <-- another chain on an existing contract
				"contractB": {1: []byte("addr1")}, // <-- another contract same address as c1
			},
			expState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {
					1: []byte("addr1"),
					2: []byte("addr2"),
				},
				"contractB": {1: []byte("addr1")},
			},
		},
		{
			name: "overwrite an existing contract address",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {
					1: []byte("addr1"),
					2: []byte("addr2"),
				},
				"contractB": {1: []byte("addr1")},
			},
			appendedState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {1: []byte("addr1-new")}, // <-- set new address on contractA.2
			},
			expState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {
					1: []byte("addr1-new"),
					2: []byte("addr2"),
				},
				"contractB": {1: []byte("addr1")},
			},
		},
		{
			name: "setting a nil state should be a no-op",
			state: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {
					1: []byte("addr1"),
					2: []byte("addr2"),
				},
				"contractB": {1: []byte("addr1")},
			},
			appendedState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": nil, // <-- set contractA to nil should be a no-op
			},
			expState: map[ContractName]map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
				"contractA": {
					1: []byte("addr1"),
					2: []byte("addr2"),
				},
				"contractB": {1: []byte("addr1")},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ab := NewBook()

			require.NoError(t, ab.InsertOrUpdate(tc.state)) // set initial state
			require.Equal(t, tc.state, ab.mem)              // assert initial state was set correctly

			require.NoError(t, ab.InsertOrUpdate(tc.appendedState))
			require.Equal(t, tc.expState, ab.mem)
		})
	}
}

func TestBook_Async(t *testing.T) {
	ab := NewBook()

	// spawn a goroutine that keeps adding to the state every 1ms
	go func() {
		for _, contractName := range []string{"a", "b", "c"} {
			for _, chainSel := range []ccipocr3.ChainSelector{1, 2, 3} {
				for _, contractAddr := range []ccipocr3.UnknownAddress{{1}, {2}, {3}} {
					appendedState := ContractAddresses{
						ContractName(contractName): map[ccipocr3.ChainSelector]ccipocr3.UnknownAddress{
							chainSel: contractAddr,
						},
					}
					err := ab.InsertOrUpdate(appendedState)
					require.NoError(t, err)
					time.Sleep(time.Millisecond)
				}
			}
		}
	}()

	// keep reading every 1ms until we see the final state for a target contract
	require.Eventually(t, func() bool {
		addr, err := ab.GetContractAddress("b", 3)
		if err != nil {
			return false
		}
		if !bytes.Equal(addr, []byte{3}) {
			return false
		}

		return true
	}, tests.WaitTimeout(t), time.Millisecond)
}
