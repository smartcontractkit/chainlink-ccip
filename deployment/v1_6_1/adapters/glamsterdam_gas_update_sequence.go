package adapters

import (
	"fmt"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
)

// GlamsterdamGasUpdateSequenceInput is the input to GlamsterdamGasUpdateSequence.
type GlamsterdamGasUpdateSequenceInput struct {
	Adapter                    GasUpdateAdapter
	TargetChainSelector        uint64
	CandidateChainSelectors    []uint64
	Report                     *glamsterdamutils.Report
}

// GlamsterdamGasUpdateSequenceOutput is the output of GlamsterdamGasUpdateSequence.
type GlamsterdamGasUpdateSequenceOutput struct {
	BatchOps []mcms_types.BatchOperation
}

// GlamsterdamGasUpdateSequence orchestrates the v1.6.1 Glamsterdam gas config update for a single
// chain family, using the provided adapter to read/write on-chain values.
func GlamsterdamGasUpdateSequence(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	input GlamsterdamGasUpdateSequenceInput,
) (GlamsterdamGasUpdateSequenceOutput, error) {
	var batchOps []mcms_types.BatchOperation
	adapter := input.Adapter
	report := input.Report
	target := input.TargetChainSelector

	for _, chainSel := range input.CandidateChainSelectors {
		// Check if this chain has a lane to the target
		hasLane, err := adapter.HasLaneToTarget(b, chains, ds, chainSel, target)
		if err != nil {
			report.AddReadError(chainSel, "check lane to target", err)
			continue
		}
		if !hasLane {
			report.AddNoLane(chainSel)
			continue
		}

		// Read current on-chain field values
		currentFields, err := adapter.ReadDestGasFields(b, chains, ds, chainSel, target)
		if err != nil {
			report.AddReadError(chainSel, "read dest gas fields", err)
			continue
		}

		// Resolve each field against baseline
		resolved := make(map[string]uint32)
		for _, fieldSpec := range getAllFieldSpecs() {
			current, ok := currentFields[fieldSpec.Name]
			if !ok {
				// Field not present in read output, skip
				continue
			}
			result := glamsterdamutils.Resolve(fieldSpec, current)
			glamsterdamutils.AddField(report, chainSel, result)
			resolved[fieldSpec.Name] = result.AppliedValue
		}

		// Write resolved values back on-chain
		if len(resolved) > 0 {
			writes, err := adapter.WriteDestGasFields(b, chains, ds, chainSel, target, resolved)
			if err != nil {
				report.AddReadError(chainSel, "write dest gas fields", err)
				continue
			}
			batchOps = append(batchOps, writes...)
		}

		// Read and check immutable sanity fields
		sanityFields, err := adapter.ReadImmutableSanityFields(b, chains, ds, chainSel)
		if err != nil {
			report.AddReadError(chainSel, "read immutable sanity fields", err)
			continue
		}
		checkImmutableFields(report, chainSel, sanityFields)

		// Process token-specific gas fields
		tokens, err := adapter.DiscoverCandidateTokens(b, chains, ds, chainSel)
		if err != nil {
			report.AddReadError(chainSel, "discover candidate tokens", err)
			continue
		}

		for _, token := range tokens {
			currentToken, isConfigured, err := adapter.ReadTokenGasField(b, chains, ds, chainSel, target, token)
			if err != nil {
				report.AddReadError(chainSel, fmt.Sprintf("read token gas field for token %x", token), err)
				continue
			}
			if !isConfigured {
				continue
			}

			// Resolve token field
			tokenFieldSpec := getTokenFieldSpec()
			result := glamsterdamutils.Resolve(tokenFieldSpec, currentToken)
			report.AddLine(fmt.Sprintf("chain %d: token %x - %s", chainSel, token, getTokenFieldReportSuffix(result)))

			// Write token field if value changed
			if result.AppliedValue != currentToken {
				write, err := adapter.WriteTokenGasField(b, chains, ds, chainSel, target, token, result.AppliedValue)
				if err != nil {
					report.AddReadError(chainSel, fmt.Sprintf("write token gas field for token %x", token), err)
					continue
				}
				batchOps = append(batchOps, write)
			}
		}
	}

	return GlamsterdamGasUpdateSequenceOutput{
		BatchOps: batchOps,
	}, nil
}

// getAllFieldSpecs returns all the field specs for v1.6.1.
func getAllFieldSpecs() []glamsterdamutils.FieldSpec[uint32] {
	return []glamsterdamutils.FieldSpec[uint32]{
		FeeQuoterDestGasOverhead,
		FeeQuoterDefaultTokenDestGasOverhead,
	}
}

// getTokenFieldSpec returns the token field spec for v1.6.1.
func getTokenFieldSpec() glamsterdamutils.FieldSpec[uint32] {
	return USDCTokenPoolDestGasOverhead
}

// getTokenFieldReportSuffix returns a report string suffix for token field results.
func getTokenFieldReportSuffix(result glamsterdamutils.FieldResult[uint32]) string {
	if result.Matched {
		return fmt.Sprintf("%s matched expected Prague value %v, applying Glamsterdam value %v",
			result.Spec.Name, result.Spec.ExpectedPrague, result.AppliedValue)
	}
	return fmt.Sprintf("%s MISMATCH - current value %v does not match expected Prague value %v, "+
		"applying fallback value %v instead of literal Glamsterdam value %v",
		result.Spec.Name, result.Current, result.Spec.ExpectedPrague,
		result.AppliedValue, result.Spec.GlamsterdamValue)
}

// checkImmutableFields validates immutable fields and adds report lines for mismatches.
func checkImmutableFields(report *glamsterdamutils.Report, chainSel uint64, sanityFields map[string]uint32) {
	expectedGasForCallExactCheck := uint32(OffRampExpectedGasForCallExactCheck)
	if actual, ok := sanityFields["OffRamp.GasForCallExactCheck"]; ok {
		if actual != expectedGasForCallExactCheck {
			report.AddLine(fmt.Sprintf(
				"chain %d: WARNING - OffRamp.GasForCallExactCheck is %d, expected %d (immutable, cannot be changed)",
				chainSel, actual, expectedGasForCallExactCheck,
			))
		}
	}
}
