package summary

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/go-viper/mapstructure/v2"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

type commitSummary struct {
	logs      []*parse.Data
	seqNumber int
}

func commitObservationSummary(logs []*parse.Data) string {
	seqNumStr := func(chains []plugintypes.SeqNumChain) string {
		if len(chains) == 0 {
			return ""
		}
		var nums []string
		for _, chain := range chains {
			nums = append(nums, fmt.Sprintf("%d->%d", chain.ChainSel, chain.SeqNum))
		}
		return strings.Join(nums, ", ")
	}

	var buf strings.Builder
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
					parts = append(parts, fmt.Sprintf("OnRampMaxSeqNums: %s", seqNumStr(obs.OnRampMaxSeqNums)))
				}
				if len(obs.OffRampNextSeqNums) > 0 {
					parts = append(parts, fmt.Sprintf("OffRampNextSeqNums: %s", seqNumStr(obs.OffRampNextSeqNums)))
				}

				if len(parts) == 0 {
					return "Observation: no seqNum data"
				} else {
					return fmt.Sprintf("Observation: \n\t* %s", strings.Join(parts, "\n\t* "))
				}
			}
		}
	}

	return buf.String()
}

func commitOutcomeSummary(logs []*parse.Data) string {
	var buf strings.Builder
	var outcomeTypeName string
	var parts []string
	for _, log := range logs {
		if log.GetMessage() == merkleroot.SendingOutcome {
			if raw, ok := log.RawLoggerFields["outcome"].(map[string]interface{}); ok {
				if t, ok := raw["outcomeType"]; ok {
					if outcomeType, ok := t.(float64); ok {
						switch merkleroot.OutcomeType(int(outcomeType)) {
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
					}
				}
				if reportRange, ok := raw["rangesSelectedForReport"].([]any); ok {
					for _, rng := range reportRange {
						if v, ok := rng.(map[string]any); ok {
							i := int(v["chain"].(float64))
							chain := fmt.Sprintf("%d", i)
							ranges := fmt.Sprintf("%v", v["seqNumRange"])
							parts = append(parts, fmt.Sprintf("%s %s", chain, ranges))
						}
					}
				}
			}
		}
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF9933"))
	buf.WriteString("Outcome [")
	buf.WriteString(style.Render(outcomeTypeName))
	buf.WriteString("]")
	if len(parts) > 0 {
		buf.WriteString(": ")
		buf.WriteString(strings.Join(parts, ", "))
	}
	return fmt.Sprintf(buf.String())
}

func commitReportSummary(logs []*parse.Data) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#339966"))

	for _, log := range logs {
		reportsMatches := reportRegex.FindStringSubmatch(log.GetMessage())
		if len(reportsMatches) > 1 {
			return fmt.Sprintf("Number of reports: %s", style.Render(reportsMatches[1]))
		}
	}
	return ""
}

func (es commitSummary) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%3d: %d logs", es.seqNumber, len(es.logs)))
	if obs := commitObservationSummary(es.logs); obs != "" {
		b.WriteString("\n     ")
		b.WriteString(obs)
	}
	if ocm := commitOutcomeSummary(es.logs); ocm != "" {
		b.WriteString("\n     ")
		b.WriteString(ocm)
	}
	if rpt := commitReportSummary(es.logs); rpt != "" {
		b.WriteString("\n     ")
		b.WriteString(rpt)
	}

	return b.String()
}

// commitCollector grabs merkle root OCR data and stores it in the summaryRenderer.
func (sr summaryRenderer) commitCollector(data *parse.Data) {
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
