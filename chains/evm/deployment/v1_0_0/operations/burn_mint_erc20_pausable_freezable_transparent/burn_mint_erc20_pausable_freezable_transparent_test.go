package burn_mint_erc20_pausable_freezable_transparent

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestValidateProxyConstructorArgs(t *testing.T) {
	validLogic := common.HexToAddress("0x1111111111111111111111111111111111111111")
	validOwner := common.HexToAddress("0x2222222222222222222222222222222222222222")

	testCases := []struct {
		name    string
		args    ProxyConstructorArgs
		wantErr string
	}{
		{
			name: "valid",
			args: ProxyConstructorArgs{Logic: validLogic, InitialOwner: validOwner},
		},
		{
			name:    "zero logic address",
			args:    ProxyConstructorArgs{Logic: common.Address{}, InitialOwner: validOwner},
			wantErr: "logic (implementation) address must not be the zero address",
		},
		{
			name:    "zero initialOwner address",
			args:    ProxyConstructorArgs{Logic: validLogic, InitialOwner: common.Address{}},
			wantErr: "initialOwner (proxy admin / upgrade authority) must not be the zero address",
		},
		{
			name:    "both zero",
			args:    ProxyConstructorArgs{},
			wantErr: "logic (implementation) address must not be the zero address",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateProxyConstructorArgs(tc.args)
			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}

func TestValidateInitializeArgs(t *testing.T) {
	validAdmin := common.HexToAddress("0x3333333333333333333333333333333333333333")

	testCases := []struct {
		name    string
		args    InitializeArgs
		wantErr string
	}{
		{
			name: "valid with maxSupply and preMint",
			args: InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(1000), PreMint: big.NewInt(500)},
		},
		{
			name: "valid with preMint equal to maxSupply",
			args: InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(1000), PreMint: big.NewInt(1000)},
		},
		{
			name: "valid with nil maxSupply/preMint (unlimited supply, no pre-mint)",
			args: InitializeArgs{DefaultAdmin: validAdmin},
		},
		{
			name: "valid with maxSupply set and preMint nil",
			args: InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(1000)},
		},
		{
			name:    "zero defaultAdmin",
			args:    InitializeArgs{DefaultAdmin: common.Address{}, MaxSupply: big.NewInt(1000), PreMint: big.NewInt(500)},
			wantErr: "defaultAdmin must not be the zero address",
		},
		{
			name:    "preMint exceeds maxSupply",
			args:    InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(1000), PreMint: big.NewInt(1001)},
			wantErr: "preMint (1001) exceeds maxSupply (1000)",
		},
		{
			name:    "preMint set but maxSupply nil (unlimited supply) is rejected",
			args:    InitializeArgs{DefaultAdmin: validAdmin, PreMint: big.NewInt(500)},
			wantErr: "preMint requires a bounded maxSupply: preMint (500) cannot be minted against unlimited supply",
		},
		{
			name:    "maxSupply zero (unlimited) with preMint set is rejected",
			args:    InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(0), PreMint: big.NewInt(500)},
			wantErr: "preMint requires a bounded maxSupply: preMint (500) cannot be minted against unlimited supply",
		},
		{
			name: "preMint zero with maxSupply nil (unlimited supply, no pre-mint) is valid",
			args: InitializeArgs{DefaultAdmin: validAdmin, PreMint: big.NewInt(0)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateInitializeArgs(tc.args)
			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}

func TestRoleConstants(t *testing.T) {
	require.Equal(t, [32]byte(crypto.Keccak256Hash([]byte("PAUSER_ROLE"))), PauserRole)
	require.Equal(t, [32]byte(crypto.Keccak256Hash([]byte("FREEZER_ROLE"))), FreezerRole)
	require.NotEqual(t, PauserRole, FreezerRole)
}

func TestValidateRoleAssignment(t *testing.T) {
	testCases := []struct {
		name    string
		args    RoleAssignment
		wantErr string
	}{
		{
			name: "valid",
			args: RoleAssignment{Role: PauserRole, To: common.HexToAddress("0x5555555555555555555555555555555555555555")},
		},
		{
			name:    "zero recipient",
			args:    RoleAssignment{Role: PauserRole, To: common.Address{}},
			wantErr: "role recipient must not be the zero address",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateRoleAssignment(tc.args)
			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}

func TestValidateBeginDefaultAdminTransferArgs(t *testing.T) {
	testCases := []struct {
		name     string
		newAdmin common.Address
		wantErr  string
	}{
		{
			name:     "valid",
			newAdmin: common.HexToAddress("0x4444444444444444444444444444444444444444"),
		},
		{
			name:     "zero address",
			newAdmin: common.Address{},
			wantErr:  "newAdmin must not be the zero address",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateBeginDefaultAdminTransferArgs(tc.newAdmin)
			if tc.wantErr == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantErr)
			}
		})
	}
}
