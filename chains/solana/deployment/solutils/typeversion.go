package solutils

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
)

// ErrNoTypeVersion is returned when a program does not implement type_version. Callers should
// treat it as "unknown", not as a failure: SPL mints, address lookup tables, MCMS accounts, PDAs
// and every program built before the instruction existed all answer this way, and they make up the
// large majority of the addresses recorded in the datastore.
var ErrNoTypeVersion = errors.New("program does not implement type_version")

// typeVersionDiscriminator is the Anchor discriminator for the type_version instruction.
var typeVersionDiscriminator = []byte{129, 251, 8, 243, 122, 229, 252, 164}

// clockSysvar is the single account type_version declares: [0] = [] clock.
var clockSysvar = solana.MustPublicKeyFromBase58("SysvarC1ock11111111111111111111111111111111")

// TypeVersion is what a CCIP Solana program reports about itself, e.g.
// Name "ccip-router", Version 1.6.2.
type TypeVersion struct {
	Name    string
	Version semver.Version
}

func (tv TypeVersion) String() string {
	return tv.Name + " " + tv.Version.String()
}

// ReadTypeVersion asks a program what it actually is.
//
// The answer is not stored in an account: type_version writes it to transaction return data, so it
// has to be simulated rather than read. Simulation runs with sigVerify disabled, so no signature is
// needed, but the fee payer must be a funded, system-owned account on the cluster being queried --
// an absent or non-system payer makes the transaction fail before the program executes, which looks
// identical to a program that has no type_version.
//
// Returns ErrNoTypeVersion when the program does not implement the instruction.
func ReadTypeVersion(
	ctx context.Context, client *solRpc.Client, programID, feePayer solana.PublicKey,
) (TypeVersion, error) {
	if feePayer.IsZero() {
		return TypeVersion{}, errors.New("fee payer is required: simulation fails before the program runs without one")
	}

	ix := solana.NewInstruction(programID, solana.AccountMetaSlice{
		{PublicKey: clockSysvar},
	}, typeVersionDiscriminator)

	tx, err := solana.NewTransaction([]solana.Instruction{ix}, solana.Hash{}, solana.TransactionPayer(feePayer))
	if err != nil {
		return TypeVersion{}, fmt.Errorf("building type_version transaction for %s: %w", programID, err)
	}
	// An unsigned transaction is fine under sigVerify=false, but the signature slot must exist for
	// the wire format to be valid.
	tx.Signatures = []solana.Signature{{}}

	raw, err := tx.MarshalBinary()
	if err != nil {
		return TypeVersion{}, fmt.Errorf("encoding type_version transaction for %s: %w", programID, err)
	}

	// solana-go v1.13's typed SimulateTransaction result does not expose returnData, so the call is
	// issued through the generic path and decoded here. Revisit if the typed result gains the field.
	var out struct {
		Value struct {
			Err        any `json:"err"`
			ReturnData *struct {
				Data []string `json:"data"`
			} `json:"returnData"`
		} `json:"value"`
	}
	err = client.RPCCallForInto(ctx, &out, "simulateTransaction", []any{
		base64.StdEncoding.EncodeToString(raw),
		map[string]any{
			"sigVerify":              false,
			"replaceRecentBlockhash": true,
			"encoding":               "base64",
		},
	})
	if err != nil {
		return TypeVersion{}, fmt.Errorf("simulating type_version for %s: %w", programID, err)
	}

	if out.Value.ReturnData == nil || len(out.Value.ReturnData.Data) == 0 || out.Value.ReturnData.Data[0] == "" {
		if e := fmt.Sprint(out.Value.Err); out.Value.Err != nil {
			// Distinguish "this program has no type_version" from "the simulation never ran", so a
			// bad fee payer is not silently recorded as an absent instruction.
			if strings.Contains(e, "AccountNotFound") || strings.Contains(e, "InvalidAccountForFee") {
				return TypeVersion{}, fmt.Errorf("fee payer %s is not usable on this cluster: %s", feePayer, e)
			}
		}
		return TypeVersion{}, ErrNoTypeVersion
	}

	data, err := base64.StdEncoding.DecodeString(out.Value.ReturnData.Data[0])
	if err != nil {
		return TypeVersion{}, fmt.Errorf("decoding type_version return data for %s: %w", programID, err)
	}

	return parseTypeVersion(data, programID)
}

// parseTypeVersion decodes the borsh string type_version returns -- a 4-byte little-endian length
// followed by the bytes -- and splits it into name and version, e.g. "ccip-router 1.6.2".
func parseTypeVersion(data []byte, programID solana.PublicKey) (TypeVersion, error) {
	if len(data) < 4 {
		return TypeVersion{}, fmt.Errorf("type_version return data for %s is %d bytes, too short for a borsh string", programID, len(data))
	}
	n := binary.LittleEndian.Uint32(data[:4])
	if int(n) > len(data)-4 {
		return TypeVersion{}, fmt.Errorf("type_version return data for %s declares %d bytes but carries %d", programID, n, len(data)-4)
	}

	name, rawVersion, ok := strings.Cut(string(data[4:4+n]), " ")
	if !ok {
		return TypeVersion{}, fmt.Errorf("type_version for %s returned %q, want \"<name> <version>\"", programID, string(data[4:4+n]))
	}

	// Pre-release versions are expected and valid here: deployed programs report things like
	// 0.1.0-dev and 1.6.1-candidate, and those must round-trip rather than be rejected or flattened.
	version, err := semver.NewVersion(rawVersion)
	if err != nil {
		return TypeVersion{}, fmt.Errorf("type_version for %s returned version %q: %w", programID, rawVersion, err)
	}

	return TypeVersion{Name: name, Version: *version}, nil
}
