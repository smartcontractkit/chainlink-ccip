// Package glamsterdam holds the version-independent primitives shared by the v1.6 and v2.0
// "update gas config for Glamsterdam" changesets: comparing an on-chain value against its
// expected Prague baseline, falling back to a derived value on mismatch, and formatting the
// resulting per-chain report.
package glamsterdam

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/exp/constraints"
)

// FieldSpec describes one on-chain gas-related field that needs to move from its Prague baseline
// to a Glamsterdam value, plus the fallback to compute instead when the current on-chain value
// doesn't match ExpectedPrague.
type FieldSpec[T comparable] struct {
	// Name identifies the field for logging/reporting, e.g. "FeeQuoter.DestChainConfig.DestGasOverhead".
	Name string
	// ExpectedPrague is the baseline value the source doc assumed when deriving GlamsterdamValue.
	ExpectedPrague T
	// GlamsterdamValue is the literal value to apply when the current on-chain value matches
	// ExpectedPrague.
	GlamsterdamValue T
	// Fallback computes the value to apply when the current on-chain value does not match
	// ExpectedPrague.
	Fallback func(current T) T
}

// ApplyRatio returns a Fallback function that scales a mismatched current value by the same
// ratio the source doc used to go from its Prague baseline to its Glamsterdam value, rounded to
// the nearest integer.
func ApplyRatio[T constraints.Integer](prague, glamsterdam T) func(current T) T {
	return func(current T) T {
		return T(math.Round(float64(current) * float64(glamsterdam) / float64(prague)))
	}
}

// FieldResult is the outcome of resolving one FieldSpec against a chain's current on-chain value.
type FieldResult[T comparable] struct {
	Spec         FieldSpec[T]
	Current      T
	Matched      bool
	AppliedValue T
}

// Resolve compares current against spec.ExpectedPrague. If it matches, the literal
// GlamsterdamValue is applied; otherwise spec.Fallback(current) is applied instead.
func Resolve[T comparable](spec FieldSpec[T], current T) FieldResult[T] {
	if current == spec.ExpectedPrague {
		return FieldResult[T]{Spec: spec, Current: current, Matched: true, AppliedValue: spec.GlamsterdamValue}
	}
	return FieldResult[T]{Spec: spec, Current: current, Matched: false, AppliedValue: spec.Fallback(current)}
}

// FieldResultString renders a FieldResult as a single human-readable report line for a given
// chain selector.
func FieldResultString[T comparable](chainSelector uint64, result FieldResult[T]) string {
	if result.Matched {
		return fmt.Sprintf(
			"chain %d: %s matched expected Prague value %v, applying Glamsterdam value %v",
			chainSelector, result.Spec.Name, result.Spec.ExpectedPrague, result.AppliedValue,
		)
	}
	return fmt.Sprintf(
		"chain %d: %s MISMATCH - current value %v does not match expected Prague value %v, "+
			"applying fallback value %v instead of literal Glamsterdam value %v",
		chainSelector, result.Spec.Name, result.Current, result.Spec.ExpectedPrague,
		result.AppliedValue, result.Spec.GlamsterdamValue,
	)
}

// Report accumulates the human-readable summary of a Glamsterdam gas-update run: which chains
// were skipped, which had no lane to the target, which fields matched or mismatched their
// expected Prague baseline, and which chains had a contract that couldn't be resolved. Intended
// for inclusion in the resulting MCMS proposal's Description and/or logs.
type Report struct {
	lines []string
}

// NewReport returns an empty Report.
func NewReport() *Report {
	return &Report{}
}

// AddSkipped records a chain that was unconditionally skipped via SkipChainSelectors.
func (r *Report) AddSkipped(chainSelector uint64) {
	r.lines = append(r.lines, fmt.Sprintf("chain %d: skipped (explicit SkipChainSelectors entry)", chainSelector))
}

// AddNoLane records a chain that was scanned but has no lane pointed at the target chain.
func (r *Report) AddNoLane(chainSelector uint64) {
	r.lines = append(r.lines, fmt.Sprintf("chain %d: no lane to target chain, skipped", chainSelector))
}

// AddUnresolvedContract records a chain where an expected contract could not be resolved from
// the datastore. This chain is excluded from further processing, but the run continues for
// every other chain.
func (r *Report) AddUnresolvedContract(chainSelector uint64, contractName string) {
	r.lines = append(r.lines, fmt.Sprintf(
		"chain %d: ERROR - could not resolve %s address, skipping this chain", chainSelector, contractName,
	))
}

// AddLine appends an arbitrary pre-formatted line to the report, e.g. the output of
// FieldResultString.
func (r *Report) AddLine(line string) {
	r.lines = append(r.lines, line)
}

// AddField records the outcome of resolving a single FieldSpec on a given chain.
func AddField[T comparable](r *Report, chainSelector uint64, result FieldResult[T]) {
	r.AddLine(FieldResultString(chainSelector, result))
}

// String joins all recorded lines into a single newline-separated report.
func (r *Report) String() string {
	return strings.Join(r.lines, "\n")
}
