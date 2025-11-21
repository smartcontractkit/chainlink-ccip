package utils

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
)

// UpgradeableLoaderState mirrors the Rust enum in the Solana SDK.
type UpgradeableLoaderState struct {
	Type          uint32
	Program       *Program
	ProgramData   *ProgramData
	Uninitialized bool
}

// Program holds the address of the ProgramData account.
type Program struct {
	ProgramData solana.PublicKey
}

// ProgramData holds the optional UpgradeAuthority.
type ProgramData struct {
	Slot            uint64
	AuthorityOption uint32 // 0 = none, 1 = present
	Authority       solana.PublicKey
}

func decodeUpgradeableLoaderState(data []byte) (*UpgradeableLoaderState, error) {
	if len(data) < 4 {
		return nil, errors.New("data too short")
	}
	state := &UpgradeableLoaderState{}
	state.Type = binary.LittleEndian.Uint32(data[:4])

	switch state.Type {
	case 2: // Program
		if len(data) < 36 {
			return nil, errors.New("program data too short")
		}
		state.Program = &Program{
			ProgramData: solana.PublicKeyFromBytes(data[4:36]),
		}
	case 3: // ProgramData
		slot := binary.LittleEndian.Uint64(data[4:12])
		opt := data[12]
		var auth *solana.PublicKey
		if opt == 1 {
			if len(data) < 45 {
				return nil, errors.New("missing authority pubkey")
			}
			pk := solana.PublicKeyFromBytes(data[13:45])
			auth = &pk
		}
		state.ProgramData = &ProgramData{
			Slot:            slot,
			AuthorityOption: uint32(opt),
		}
		if state.ProgramData.AuthorityOption == 1 {
			state.ProgramData.Authority = *auth
		}
	default:
		// other variants (Uninitialized, Buffer) are not needed here
	}
	return state, nil
}

func getUpgradeableLoaderState(client *solRpc.Client, progPubkey solana.PublicKey) (*UpgradeableLoaderState, error) {
	resp, err := client.GetAccountInfo(context.Background(), progPubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch program account: %w", err)
	}
	if resp.Value == nil {
		return nil, errors.New("program account does not exist")
	}

	state, err := decodeUpgradeableLoaderState(resp.Value.Data.GetBinary())
	if err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}
	return state, nil
}

func GetUpgradeAuthority(client *solRpc.Client, progDataPubkey solana.PublicKey) (solana.PublicKey, error) {
	state, err := getUpgradeableLoaderState(client, progDataPubkey)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get upgrade authority for program data %s: %w", progDataPubkey.String(), err)
	}

	if state.ProgramData == nil {
		return solana.PublicKey{}, errors.New("unexpected state: not programdata")
	}

	if state.ProgramData.AuthorityOption == 0 {
		// No authority – the program is immutable
		return solana.PublicKey{}, nil
	}
	return state.ProgramData.Authority, nil
}
