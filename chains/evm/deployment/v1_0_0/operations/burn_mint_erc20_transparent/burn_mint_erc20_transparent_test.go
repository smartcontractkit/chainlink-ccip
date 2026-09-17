package burn_mint_erc20_transparent

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
			name: "preMint set but maxSupply nil (unlimited supply) does not trigger the exceeds-check",
			args: InitializeArgs{DefaultAdmin: validAdmin, PreMint: big.NewInt(500)},
		},
		{
			name: "maxSupply zero (unlimited) with preMint set does not trigger the exceeds-check",
			args: InitializeArgs{DefaultAdmin: validAdmin, MaxSupply: big.NewInt(0), PreMint: big.NewInt(500)},
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
