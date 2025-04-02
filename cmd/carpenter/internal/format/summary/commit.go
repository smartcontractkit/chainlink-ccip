package summary

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-viper/mapstructure/v2"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

func init() {
	reportRegex = regexp.MustCompile("^built (\\d+) reports$")
	divider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#75A3A3"))

	section = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#3366FF"))
	number = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#339966"))
	highlight = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF9933"))
}

var divider lipgloss.Style
var section lipgloss.Style
var highlight lipgloss.Style
var number lipgloss.Style

var reportRegex *regexp.Regexp

const padding = "    "

var bullet = fmt.Sprintf("\n%s* ", padding)

type commitSummary struct {
	logs      []*parse.Data
	seqNumber int
}

func commitObservationSummary(logs []*parse.Data) string {
	for _, log := range logs {
		if log.GetMessage() == merkleroot.SendingObservation {
			if raw, ok := log.RawLoggerFields["observation"].(map[string]interface{}); ok {
				var obs merkleroot.Observation
				err := mapstructure.Decode(raw, &obs)
				if err != nil {
					// ignore errors -- it complains about non-string map keys:
					//   * RMNEnabledChains
					//   * FChain
				}
				var parts []string
				if len(obs.OnRampMaxSeqNums) > 0 {
					parts = append(parts, fmt.Sprintf("OnRampMaxSeqNums: %s",
						commitFormatSeqNumChains(obs.OnRampMaxSeqNums)))
				}
				if len(obs.OffRampNextSeqNums) > 0 {
					parts = append(parts, fmt.Sprintf("OffRampNextSeqNums: %s",
						commitFormatSeqNumChains(obs.OffRampNextSeqNums)))
				}

				var buf strings.Builder
				buf.WriteString(padding)
				buf.WriteString(section.Render("Observation"))
				if len(parts) == 0 {
					buf.WriteString(": no seqNum data")
				} else {
					buf.WriteString(bullet)
					buf.WriteString(strings.Join(parts, bullet))
				}
				return buf.String()
			}
		}
	}

	return ""
}

func commitOutcomeSummary(logs []*parse.Data) string {
	for _, log := range logs {
		if log.GetMessage() == merkleroot.SendingOutcome {
			var buf strings.Builder
			var outcomeTypeName string
			var parts []string

			if raw, ok := log.RawLoggerFields["outcome"].(map[string]interface{}); ok {
				var otc merkleroot.Outcome
				err := mapstructure.Decode(raw, &otc)
				if err != nil {
					// Ignore RMNRemoteCfg errors due to slice/value type mismatches.
				}

				switch otc.OutcomeType {
				case merkleroot.ReportIntervalsSelected:
					outcomeTypeName = "ReportIntervalsSelected"
				case merkleroot.ReportGenerated:
					outcomeTypeName = "ReportGenerated"
				case merkleroot.ReportEmpty:
					outcomeTypeName = "ReportEmpty"
				case merkleroot.ReportInFlight:
					outcomeTypeName = "ReportInFlight"
				case merkleroot.ReportTransmitted:
					outcomeTypeName = "ReportTransmitted"
				case merkleroot.ReportTransmissionFailed:
					outcomeTypeName = "ReportTransmissionFailed"
				default:
					outcomeTypeName = "unknown"
				}

				if len(otc.RootsToReport) > 0 {
					parts = append(parts, fmt.Sprintf("RootsToReport: %d", len(otc.RootsToReport)))
				}
				if len(otc.RMNReportSignatures) != 0 {
					parts = append(parts, fmt.Sprintf("RMNReportSignatures: %d", len(otc.RMNReportSignatures)))
				}
				if len(otc.RangesSelectedForReport) > 0 {
					parts = append(parts, fmt.Sprintf("RangesSelectedForReport: %s",
						commitFormatChainRanges(otc.RangesSelectedForReport)))
				}
				if len(otc.OffRampNextSeqNums) > 0 {
					parts = append(parts, fmt.Sprintf("OffRampNextSeqNums: %s",
						commitFormatSeqNumChains(otc.OffRampNextSeqNums)))
				}
				if otc.ReportTransmissionCheckAttempts != 0 {
					parts = append(parts, fmt.Sprintf("ReportTransmissionCheckAttempts: %d",
						otc.ReportTransmissionCheckAttempts))
				}

				style := lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.Color("#FF9933"))
				buf.WriteString("Outcome [")
				buf.WriteString(style.Render(outcomeTypeName))
				buf.WriteString("]")
				if len(parts) > 0 {
					buf.WriteString(bullet)
					buf.WriteString(strings.Join(parts, bullet))
				}
				return fmt.Sprintf(buf.String())
			}
		}
	}

	return ""
}

func commitReportSummary(logs []*parse.Data) string {
	var numReports string
	var reportParts []string
	for _, log := range logs {
		message := log.GetMessage()
		if message == "generating report" {
			if roots, ok := log.RawLoggerFields["roots"].([]interface{}); ok {
				reportParts = append(reportParts, fmt.Sprintf("Roots: %d", len(roots)))
			}
			if prices, ok := log.RawLoggerFields["tokenPriceUpdates"].(map[string]interface{}); ok {
				reportParts = append(reportParts, fmt.Sprintf("TokenPriceUpdates: %d", len(prices)))
			}
			if gasUpdates, ok := log.RawLoggerFields["gasPriceUpdates"].([]interface{}); ok {
				reportParts = append(reportParts, fmt.Sprintf("GasPriceUpdates: %d", len(gasUpdates)))
			}
			if sigs, ok := log.RawLoggerFields["rmnSignatures"].(map[string]interface{}); ok {
				reportParts = append(reportParts, fmt.Sprintf("RMNSignatures: %d", len(sigs)))
			}
		} else {
			reportsMatches := reportRegex.FindStringSubmatch(message)
			if len(reportsMatches) > 1 {
				numReports = reportsMatches[1]
			}
		}
	}
	if len(reportParts) > 0 || numReports != "" {
		var buf strings.Builder
		buf.WriteString(padding)
		buf.WriteString(section.Render("Reports"))
		buf.WriteString(": ")
		buf.WriteString(number.Render(numReports))
		if numReports != "" && len(reportParts) == 0 {
			fmt.Println("???")
		}
		if len(reportParts) > 0 {
			buf.WriteString(bullet)
			buf.WriteString(strings.Join(reportParts, bullet))
		}
		return buf.String()
	}

	return ""
}

func (es commitSummary) String() string {
	var b strings.Builder
	b.WriteString(divider.Render(fmt.Sprintf("%3d: %d logs                       ", es.seqNumber, len(es.logs))))
	if obs := commitObservationSummary(es.logs); obs != "" {
		b.WriteString("\n")
		b.WriteString(obs)
	}
	if ocm := commitOutcomeSummary(es.logs); ocm != "" {
		b.WriteString("\n")
		b.WriteString(ocm)
	}
	if rpt := commitReportSummary(es.logs); rpt != "" {
		b.WriteString("\n")
		b.WriteString(rpt)
	}

	return b.String()
}

// commitCollector grabs merkle root OCR data and stores it in the summaryFormatter.
func (sr summaryFormatter) commitCollector(data *parse.Data) {
	// a regex that matches "building 1 reports" and captures the number.
	mark := false

	message := data.GetMessage()
	switch message {
	case merkleroot.SendingObservation:
		// TODO: collect observation
		mark = true
	case merkleroot.SendingOutcome:
		mark = true
	case "generating report":
		mark = true
	default:
		reportsMatches := reportRegex.FindStringSubmatch(message)
		if len(reportsMatches) > 1 {
			mark = true
		}
		if strings.HasPrefix(message, "built ") {
			mark = true
		}
	}

	if mark && data.SequenceNumber != 0 {
		if sr.commits[data.DONID] == nil {
			sr.commits[data.DONID] = make(map[int]commitSummary)
		}
		summary := sr.commits[data.DONID][data.SequenceNumber]
		summary.logs = append(summary.logs, data)
		summary.seqNumber = data.SequenceNumber
		sr.commits[data.DONID][data.SequenceNumber] = summary
	}
	if mark && data.SequenceNumber == 0 {
		fmt.Println("sequence number is 0")
	}
}

func commitFormatChainRanges(chains []plugintypes.ChainRange) string {
	if len(chains) == 0 {
		return ""
	}
	var nums []string
	for _, chain := range chains {
		nums = append(nums, fmt.Sprintf("%d->%s", chain.ChainSel, chain.SeqNumRange))
	}
	return strings.Join(nums, ", ")
}

func commitFormatSeqNumChains(chains []plugintypes.SeqNumChain) string {
	if len(chains) == 0 {
		return ""
	}
	var nums []string
	for _, chain := range chains {
		nums = append(nums, fmt.Sprintf("%d->%d", chain.ChainSel, chain.SeqNum))
	}
	return strings.Join(nums, ", ")
}
