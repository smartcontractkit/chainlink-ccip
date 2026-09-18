package solutils

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// TypeVersionLabelPrefix marks the datastore label carrying what the program itself reports.
//
// A label rather than a field because the datastore key is (chain, type, version, qualifier):
// writing the build into any of those re-keys the row, and every lookup in the Solana changesets
// is written against the key as it stands. The label records the evidence without moving anything.
const TypeVersionLabelPrefix = "type_version:"

// TypeVersionLabel renders a TypeVersion as the datastore label for it, e.g.
// "type_version:ccip-router 1.6.2".
func TypeVersionLabel(tv TypeVersion) string {
	return TypeVersionLabelPrefix + tv.String()
}

// ParseTypeVersionLabel returns the type_version string carried by a label set, and whether one
// was present. Only the first is returned: a row carries at most one, because RecordOnChainVersion
// replaces any earlier label rather than appending to it.
func ParseTypeVersionLabel(labels []string) (string, bool) {
	for _, l := range labels {
		if after, ok := strings.CutPrefix(l, TypeVersionLabelPrefix); ok {
			return after, true
		}
	}
	return "", false
}

// RecordOnChainVersion reads what the deployed program reports and returns tv with that answer
// attached as a label, for recording alongside the address.
//
// Call it after a deploy and after an upgrade, so what is recorded is what is on the chain rather
// than what the changeset intended to put there. An audit of the four live environments found
// those two had diverged for every program that answers -- routers on 1.6.2, off-ramps and RMN
// remotes on 1.6.3, a fee quoter on 1.6.4, 22 mainnet token pools on 0.1.0-dev builds -- because
// nothing was ever reading the deployed version back.
//
// What it does NOT do, deliberately:
//
//   - It does not touch tv.Type. The datastore type is our taxonomy ("BurnMintTokenPool"); the
//     name type_version reports is the program's own ("burnmint-token-pool"). They are different
//     namespaces and conflating them would break every lookup keyed on the first.
//   - It does not touch tv.Version. That field records the generation ("a 1.6-era off-ramp"),
//     which is how all ~45 reader sites use it, and it is part of the key.
//
// A program that does not implement type_version is not an error: most recorded addresses are
// mints, lookup tables, MCMS accounts and PDAs, which have no version to report. Those return tv
// unchanged. Nor is an RPC failure worth failing a deployment that has already succeeded on
// chain -- it is logged and tv is returned unchanged, so the label is best-effort by design.
func RecordOnChainVersion(
	ctx context.Context,
	lggr logger.Logger,
	client *solRpc.Client,
	programID, feePayer solana.PublicKey,
	tv deployment.TypeAndVersion,
) deployment.TypeAndVersion {
	onChain, err := ReadTypeVersion(ctx, client, programID, feePayer)
	switch {
	case errors.Is(err, ErrNoTypeVersion):
		return tv
	case err != nil:
		lggr.Warnw("Could not read on-chain type_version; recording without it",
			"program", programID.String(), "type", string(tv.Type), "err", err)

		return tv
	}

	// Drop any label from an earlier deploy or upgrade of this same key, so the row carries what
	// is on the chain now rather than an accumulated history of what used to be.
	kept := make([]string, 0, len(tv.Labels)+1)
	for _, l := range tv.Labels.List() {
		if !strings.HasPrefix(l, TypeVersionLabelPrefix) {
			kept = append(kept, l)
		}
	}
	kept = append(kept, TypeVersionLabel(onChain))
	tv.Labels = deployment.NewLabelSet(kept...)

	if onChain.Version.Major() != tv.Version.Major() || onChain.Version.Minor() != tv.Version.Minor() {
		lggr.Warnw("Deployed program reports a different generation than the datastore records",
			"program", programID.String(), "type", string(tv.Type),
			"recorded", tv.Version.String(), "onChain", onChain.String())
	}

	return tv
}

// MustParseTypeVersionLabel is ParseTypeVersionLabel for tests and tooling that treat an absent
// label as a bug.
func MustParseTypeVersionLabel(labels []string) string {
	s, ok := ParseTypeVersionLabel(labels)
	if !ok {
		panic(fmt.Sprintf("no %s label in %v", TypeVersionLabelPrefix, labels))
	}

	return s
}
