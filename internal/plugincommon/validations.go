package plugincommon

import (
	"errors"
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ValidateFChain validates that the FChain values are positive.
func ValidateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for chainSelector, f := range fChain {
		if f <= 0 {
			return fmt.Errorf("fChain for chain %d is not positive: %d", chainSelector, f)
		}
	}
	return nil
}

// ErrReportValidation is used when the validation being performed in order to validate a report has errored.
// This does not necessarily mean that the report is valid or invalid.
var ErrReportValidation = errors.New("report validation errors")

// ErrInvalidReport is used when a report is found to be invalid.
var ErrInvalidReport = errors.New("invalid report")

// NewErrValidatingReport is returned when the report could not be validated due to an error.
func NewErrValidatingReport(err error) error {
	return &errWrappedValidatingReport{err: err}
}

type errWrappedValidatingReport struct {
	err error
}

func (e *errWrappedValidatingReport) Error() string {
	return fmt.Sprintf("report validation error: %v", e.err)
}

func (e *errWrappedValidatingReport) Unwrap() error {
	return ErrReportValidation
}

// NewErrInvalidReport is returned when the report is specifically invalid due to a validation rule.
func NewErrInvalidReport(reason string) error {
	return &errWrappedInvalidReport{reason: reason}
}

type errWrappedInvalidReport struct {
	reason string
}

func (e *errWrappedInvalidReport) Error() string {
	return fmt.Sprintf("invalid report: %s", e.reason)
}

func (e *errWrappedInvalidReport) Unwrap() error {
	return ErrInvalidReport
}
