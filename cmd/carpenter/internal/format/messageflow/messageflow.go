// Package messageflow tracks the complete lifecycle of a CCIP message
// from source chain reading through destination chain execution.
package messageflow

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	format.Register("messageflow", messageFlowFormatterFactory, "Track the complete lifecycle of a specific CCIP message by message ID.")
}

func messageFlowFormatterFactory(options format.Options) format.Formatter {
	return &messageFlowFormatter{
		messageID:       options.MessageID,
		events:          make([]MessageEvent, 0),
		merkleRoots:     make(map[string]bool),
		msgHashes:       make(map[string]bool),
		relatedSeqNums:  make(map[string]bool),
		seqNumRanges:    make(map[string]bool),
		ocrSequences:    make(map[int]bool),
		phaseTimestamps: make(map[string]time.Time),
	}
}

type MessageEvent struct {
	Timestamp time.Time
	Stage     string
	Message   string
	Data      *parse.Data
	Details   map[string]interface{}
}

type messageFlowFormatter struct {
	messageID       string
	events          []MessageEvent
	merkleRoots     map[string]bool      // Track merkle roots associated with this message
	msgHashes       map[string]bool      // Track message hashes associated with this message
	relatedSeqNums  map[string]bool      // Track sequence numbers on relevant chains
	seqNumRanges    map[string]bool      // Track sequence number ranges
	ocrSequences    map[int]bool         // Track OCR sequence numbers
	phaseTimestamps map[string]time.Time // Track phase start times for latency analysis
}

var (
	stageStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00AA00"))

	timestampStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF"))

	latencyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6600"))

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000"))
)

func (mf *messageFlowFormatter) Format(data *parse.Data) {
	if mf.messageID == "" {
		return // No message ID specified
	}

	// Check if this log is related to our message ID
	if !mf.isRelatedToMessage(data) {
		return
	}

	// Determine the stage and extract relevant information
	stage, details := mf.determineStage(data)
	if stage == "" {
		return // Not a relevant stage
	}

	event := MessageEvent{
		Timestamp: data.GetTimestamp(),
		Stage:     stage,
		Message:   data.GetMessage(),
		Data:      data,
		Details:   details,
	}

	// Track phase start times for latency analysis
	phaseKey := fmt.Sprintf("%s_%s", data.Plugin, strings.Split(stage, "_")[0])
	if _, exists := mf.phaseTimestamps[phaseKey]; !exists {
		mf.phaseTimestamps[phaseKey] = data.GetTimestamp()
	}

	mf.events = append(mf.events, event)
}

func (mf *messageFlowFormatter) isRelatedToMessage(data *parse.Data) bool {
	message := data.GetMessage()

	// Direct message ID match
	if strings.Contains(message, mf.messageID) {
		mf.extractRelatedIdentifiers(data)
		return true
	}

	// Check raw logger fields for message ID references
	if data.RawLoggerFields != nil {
		for key, value := range data.RawLoggerFields {
			if valueStr, ok := value.(string); ok {
				// Direct message ID match
				if strings.Contains(valueStr, mf.messageID) {
					mf.extractRelatedIdentifiers(data)
					return true
				}
			}
			// Check for message ID in specific fields
			if key == "messageID" || key == "msgId" || key == "message_id" {
				if valueStr, ok := value.(string); ok && valueStr == mf.messageID {
					mf.extractRelatedIdentifiers(data)
					return true
				}
			}
		}
	}

	// Check if this log contains any of our tracked related identifiers
	if mf.containsRelatedIdentifiers(data) {
		return true
	}

	return false
}

func (mf *messageFlowFormatter) extractRelatedIdentifiers(data *parse.Data) {
	if data.RawLoggerFields == nil {
		return
	}

	// Track OCR sequence numbers
	if ocrSeqNr, ok := data.RawLoggerFields["ocrSeqNr"]; ok {
		if seqInt, ok := ocrSeqNr.(float64); ok {
			mf.ocrSequences[int(seqInt)] = true
		}
	}

	// Extract merkle roots, message hashes, and sequence numbers that might be related
	for key, value := range data.RawLoggerFields {
		if valueStr, ok := value.(string); ok {
			switch key {
			case "merkleRoot", "root", "hash", "msgHash", "commitRoot", "computedRoot":
				if len(valueStr) > 10 { // Likely a hash if long enough
					mf.merkleRoots[valueStr] = true
				}
			case "hashPreimage":
				// Hash preimages might contain our message
				if strings.Contains(valueStr, mf.messageID) {
					mf.msgHashes[valueStr] = true
				}
			case "seqNum", "sequenceNumber":
				mf.relatedSeqNums[valueStr] = true
			case "sequenceNumberRange":
				mf.seqNumRanges[valueStr] = true
			}
		}

		// Handle nested structures like arrays of roots
		if key == "roots" {
			if rootsArray, ok := value.([]interface{}); ok {
				for _, rootItem := range rootsArray {
					if rootMap, ok := rootItem.(map[string]interface{}); ok {
						if rootHash, exists := rootMap["merkleRoot"]; exists {
							if rootHashStr, ok := rootHash.(string); ok {
								mf.merkleRoots[rootHashStr] = true
							}
						}
					}
				}
			}
		}

		// Handle hashes array (from "Computed merkle root" logs)
		if key == "hashes" {
			if hashesArray, ok := value.([]interface{}); ok {
				for _, hashItem := range hashesArray {
					if hashStr, ok := hashItem.(string); ok {
						mf.merkleRoots[hashStr] = true
					}
				}
			}
		}
	}
}

func (mf *messageFlowFormatter) containsRelatedIdentifiers(data *parse.Data) bool {
	message := data.GetMessage()

	// Check if this log belongs to a related OCR sequence
	if data.RawLoggerFields != nil {
		if ocrSeqNr, ok := data.RawLoggerFields["ocrSeqNr"]; ok {
			if seqInt, ok := ocrSeqNr.(float64); ok {
				if mf.ocrSequences[int(seqInt)] {
					return true
				}
			}
		}
	}

	// Check if message contains any of our tracked merkle roots
	for root := range mf.merkleRoots {
		if strings.Contains(message, root) {
			return true
		}
	}

	// Check if message contains any of our tracked message hashes
	for msgHash := range mf.msgHashes {
		if strings.Contains(message, msgHash) {
			return true
		}
	}

	// Check raw logger fields for related identifiers
	if data.RawLoggerFields != nil {
		for key, value := range data.RawLoggerFields {
			if valueStr, ok := value.(string); ok {
				// Check merkle roots
				for root := range mf.merkleRoots {
					if valueStr == root || strings.Contains(valueStr, root) {
						return true
					}
				}

				// Check message hashes
				for msgHash := range mf.msgHashes {
					if valueStr == msgHash || strings.Contains(valueStr, msgHash) {
						return true
					}
				}

				// Check sequence number ranges
				for seqRange := range mf.seqNumRanges {
					if strings.Contains(valueStr, seqRange) {
						return true
					}
				}
			}

			// Check nested structures
			switch key {
			case "roots", "merkleRoots":
				if rootsArray, ok := value.([]interface{}); ok {
					for _, rootItem := range rootsArray {
						if rootMap, ok := rootItem.(map[string]interface{}); ok {
							if rootHash, exists := rootMap["merkleRoot"]; exists {
								if rootHashStr, ok := rootHash.(string); ok {
									for trackedRoot := range mf.merkleRoots {
										if rootHashStr == trackedRoot {
											return true
										}
									}
								}
							}
						}
					}
				}
			case "hashes":
				if hashesArray, ok := value.([]interface{}); ok {
					for _, hashItem := range hashesArray {
						if hashStr, ok := hashItem.(string); ok {
							for trackedRoot := range mf.merkleRoots {
								if hashStr == trackedRoot {
									return true
								}
							}
						}
					}
				}
			case "rangesSelectedForReport", "rootsToReport", "sequenceNumberRange":
				// These might contain sequence number ranges or roots related to our message
				if ranges, ok := value.([]interface{}); ok {
					for _, rangeItem := range ranges {
						if rangeMap, ok := rangeItem.(map[string]interface{}); ok {
							// Check if any of our tracked sequence numbers fall within reported ranges
							for seqNum := range mf.relatedSeqNums {
								if strings.Contains(fmt.Sprintf("%v", rangeMap), seqNum) {
									return true
								}
							}
						}
					}
				}
			}
		}
	}

	return false
}

func (mf *messageFlowFormatter) determineStage(data *parse.Data) (string, map[string]interface{}) {
	message := data.GetMessage()
	details := make(map[string]interface{})

	// Copy relevant fields from raw logger fields
	if data.RawLoggerFields != nil {
		for key, value := range data.RawLoggerFields {
			switch key {
			case "sourceChain", "destChain", "seqNum", "messageID", "msgId", "message_id",
				"sender", "receiver", "tokenAmounts", "gasLimit", "strict", "feeToken",
				"plugin", "ocrPhase", "ocrSeqNr", "state", "component",
				"commitRoot", "computedRoot", "merkleRoot", "root", "hash":
				details[key] = value
			}
		}
	}

	// Extract OCR phase and state information
	ocrPhase := ""
	processorState := ""
	if data.RawLoggerFields != nil {
		if phase, ok := data.RawLoggerFields["ocrPhase"].(string); ok {
			ocrPhase = phase
		}
		if state, ok := data.RawLoggerFields["state"].(string); ok {
			processorState = state
		}
	}

	// Determine stage based on message content, plugin, phase, and state
	switch {
	// === SOURCE CHAIN EVENTS ===
	case strings.Contains(message, "CCIPMessageSent") || strings.Contains(message, "message sent"):
		return "📤 SOURCE_EMISSION", details

	// === COMMIT PLUGIN FLOW ===
	// Commit Observation Phase
	case strings.Contains(message, "Computed merkle root") && data.Plugin == "Commit" && ocrPhase == "obs":
		return "🌳 COMMIT_OBS_MERKLE_ROOT", details

	case strings.Contains(message, "commit plugin performing observation") && ocrPhase == "obs":
		return "🔍 COMMIT_OBS_START", details

	case strings.Contains(message, "sending merkle root processor observation") && data.Plugin == "Commit":
		return "📤 COMMIT_OBS_SENT", details

	case data.Plugin == "Commit" && ocrPhase == "obs" && strings.Contains(message, "observed"):
		return "🔍 COMMIT_OBS_DATA", details

	// Commit Outcome Phase
	case strings.Contains(message, "commit plugin performing outcome") && ocrPhase == "otcm":
		return "🎯 COMMIT_OUTCOME_START", details

	case strings.Contains(message, "Sending Outcome") && data.Plugin == "Commit":
		return "📤 COMMIT_OUTCOME_SENT", details

	case data.Plugin == "Commit" && ocrPhase == "otcm":
		return "🎯 COMMIT_OUTCOME_PROCESS", details

	// Commit Report Phase
	case strings.Contains(message, "generating report") && data.Plugin == "Commit" && ocrPhase == "rgen":
		return "📊 COMMIT_REPORT_GENERATION", details

	case strings.Contains(message, "built") && strings.Contains(message, "reports") && data.Plugin == "Commit":
		return "✅ COMMIT_REPORT_BUILT", details

	case strings.Contains(message, "transmitting report") && data.Plugin == "Commit":
		return "📡 COMMIT_REPORT_TRANSMISSION", details

	// === EXECUTE PLUGIN FLOW ===
	// Execute GetCommitReports State
	case strings.Contains(message, "execute plugin got observation") && processorState == "GetCommitReports":
		return "📋 EXEC_OBS_GET_COMMIT_REPORTS", details

	case data.Plugin == "Execute" && ocrPhase == "obs" && processorState == "GetCommitReports":
		return "📋 EXEC_OBS_COMMIT_REPORTS", details

	// Execute GetMessages State
	case strings.Contains(message, "execute plugin got observation") && processorState == "GetMessages":
		return "📨 EXEC_OBS_GET_MESSAGES", details

	case data.Plugin == "Execute" && ocrPhase == "obs" && processorState == "GetMessages":
		return "📨 EXEC_OBS_MESSAGES", details

	// Execute Filter State
	case strings.Contains(message, "Execute plugin performing observation") && processorState == "Filter":
		return "🔍 EXEC_OBS_FILTER", details

	case strings.Contains(message, "execute plugin got observation") && processorState == "Filter":
		return "🔍 EXEC_OBS_FILTER_RESULT", details

	case data.Plugin == "Execute" && ocrPhase == "obs" && processorState == "Filter":
		return "🔍 EXEC_OBS_FILTER_PROCESS", details

	// Execute Outcome Phase
	case data.Plugin == "Execute" && ocrPhase == "otcm":
		if processorState == "GetCommitReports" {
			return "🎯 EXEC_OUTCOME_COMMIT_REPORTS", details
		} else if processorState == "GetMessages" {
			return "🎯 EXEC_OUTCOME_MESSAGES", details
		} else if processorState == "Filter" {
			return "🎯 EXEC_OUTCOME_FILTER", details
		}
		return "🎯 EXEC_OUTCOME_PROCESS", details

	// Execute Report Phase
	case data.Plugin == "Execute" && ocrPhase == "rgen":
		return "📊 EXEC_REPORT_GENERATION", details

	case strings.Contains(message, "built") && strings.Contains(message, "reports") && data.Plugin == "Execute":
		return "✅ EXEC_REPORT_BUILT", details

	case strings.Contains(message, "transmitting report") && data.Plugin == "Execute":
		return "📡 EXEC_REPORT_TRANSMISSION", details

	// === MESSAGE PROCESSING ===
	case strings.Contains(message, "Hashed message, adding to tree leaves"):
		return "🌳 MESSAGE_MERKLE_TREE_BUILD", details

	case strings.Contains(message, "constructing merkle tree"):
		return "🌳 MESSAGE_MERKLE_TREE_CONSTRUCT", details

	case strings.Contains(message, "merkle root verified"):
		return "✅ MESSAGE_MERKLE_VERIFIED", details

	case strings.Contains(message, "selected messages from commit report for execution"):
		return "📋 MESSAGE_SELECTED_FOR_EXEC", details

	case strings.Contains(message, "generating merkle proofs"):
		return "🔐 MESSAGE_MERKLE_PROOFS", details

	case strings.Contains(message, "message already executed"):
		return "✅ MESSAGE_EXECUTED", details

	case strings.Contains(message, "message pseudo deleted"):
		return "🗑️ MESSAGE_CLEANUP", details

	case strings.Contains(message, "skipping pseudo-deleted message"):
		return "⏭️ MESSAGE_SKIPPED", details

	// === DESTINATION CHAIN EVENTS ===
	case strings.Contains(message, "ExecutionStateChanged") || strings.Contains(message, "message executed"):
		return "🎯 DEST_EXECUTION", details

	// === ERROR HANDLING ===
	case strings.Contains(message, "error") || strings.Contains(message, "failed"):
		if data.Plugin == "Commit" {
			if ocrPhase == "obs" {
				return "❌ COMMIT_OBS_ERROR", details
			} else if ocrPhase == "otcm" {
				return "❌ COMMIT_OUTCOME_ERROR", details
			} else if ocrPhase == "rgen" {
				return "❌ COMMIT_REPORT_ERROR", details
			}
			return "❌ COMMIT_ERROR", details
		} else if data.Plugin == "Execute" {
			if ocrPhase == "obs" {
				return "❌ EXEC_OBS_ERROR", details
			} else if ocrPhase == "otcm" {
				return "❌ EXEC_OUTCOME_ERROR", details
			} else if ocrPhase == "rgen" {
				return "❌ EXEC_REPORT_ERROR", details
			}
			return "❌ EXEC_ERROR", details
		}
		return "❌ ERROR", details

	default:
		// Categorize other logs by plugin and phase
		if data.Plugin == "Commit" {
			if ocrPhase == "obs" {
				return "📝 COMMIT_OBS_OTHER", details
			} else if ocrPhase == "otcm" {
				return "📝 COMMIT_OUTCOME_OTHER", details
			} else if ocrPhase == "rgen" {
				return "📝 COMMIT_REPORT_OTHER", details
			}
			return "📝 COMMIT_OTHER", details
		} else if data.Plugin == "Execute" {
			if ocrPhase == "obs" {
				return "📝 EXEC_OBS_OTHER", details
			} else if ocrPhase == "otcm" {
				return "📝 EXEC_OUTCOME_OTHER", details
			} else if ocrPhase == "rgen" {
				return "📝 EXEC_REPORT_OTHER", details
			}
			return "📝 EXEC_OTHER", details
		}

		// Check for other message-related activities
		if strings.Contains(message, "decoded messages") || strings.Contains(message, "queried messages") {
			return "📖 MESSAGE_QUERY", details
		}

		return "📝 OTHER", details
	}
}

func (mf *messageFlowFormatter) Close() error {
	if len(mf.events) == 0 {
		fmt.Printf("No events found for message ID: %s\n", mf.messageID)
		return nil
	}

	// Sort events by timestamp
	sort.Slice(mf.events, func(i, j int) bool {
		return mf.events[i].Timestamp.Before(mf.events[j].Timestamp)
	})

	fmt.Printf("Message Flow Analysis for Message ID: %s\n", stageStyle.Render(mf.messageID))
	fmt.Printf("═══════════════════════════════════════════════════════════════\n\n")

	var firstEvent, lastEvent *MessageEvent
	stageCount := make(map[string]int)
	commitEvents := []MessageEvent{}
	execEvents := []MessageEvent{}
	otherEvents := []MessageEvent{}

	// Categorize events by phase
	for i, event := range mf.events {
		if i == 0 {
			firstEvent = &event
		}
		lastEvent = &event
		stageCount[event.Stage]++

		if strings.Contains(event.Stage, "COMMIT") {
			commitEvents = append(commitEvents, event)
		} else if strings.Contains(event.Stage, "EXEC") {
			execEvents = append(execEvents, event)
		} else {
			otherEvents = append(otherEvents, event)
		}
	}

	// Organize events by OCR phases
	commitObsEvents := []MessageEvent{}
	commitOutcomeEvents := []MessageEvent{}
	commitReportEvents := []MessageEvent{}
	execObsEvents := []MessageEvent{}
	execOutcomeEvents := []MessageEvent{}
	execReportEvents := []MessageEvent{}

	for _, event := range mf.events {
		if strings.Contains(event.Stage, "COMMIT") {
			if strings.Contains(event.Stage, "OBS") {
				commitObsEvents = append(commitObsEvents, event)
			} else if strings.Contains(event.Stage, "OUTCOME") {
				commitOutcomeEvents = append(commitOutcomeEvents, event)
			} else if strings.Contains(event.Stage, "REPORT") {
				commitReportEvents = append(commitReportEvents, event)
			} else {
				commitEvents = append(commitEvents, event)
			}
		} else if strings.Contains(event.Stage, "EXEC") {
			if strings.Contains(event.Stage, "OBS") {
				execObsEvents = append(execObsEvents, event)
			} else if strings.Contains(event.Stage, "OUTCOME") {
				execOutcomeEvents = append(execOutcomeEvents, event)
			} else if strings.Contains(event.Stage, "REPORT") {
				execReportEvents = append(execReportEvents, event)
			} else {
				execEvents = append(execEvents, event)
			}
		}
	}

	// Print events grouped by plugin and phase
	if len(commitObsEvents) > 0 {
		fmt.Printf("🔹 COMMIT PLUGIN - OBSERVATION PHASE (%d events)\n", len(commitObsEvents))
		fmt.Printf("─────────────────────────────────────────────────\n")
		mf.printEvents(commitObsEvents)
		fmt.Println()
	}

	if len(commitOutcomeEvents) > 0 {
		fmt.Printf("🔹 COMMIT PLUGIN - OUTCOME PHASE (%d events)\n", len(commitOutcomeEvents))
		fmt.Printf("──────────────────────────────────────────────\n")
		mf.printEvents(commitOutcomeEvents)
		fmt.Println()
	}

	if len(commitReportEvents) > 0 {
		fmt.Printf("🔹 COMMIT PLUGIN - REPORT PHASE (%d events)\n", len(commitReportEvents))
		fmt.Printf("─────────────────────────────────────────────\n")
		mf.printEvents(commitReportEvents)
		fmt.Println()
	}

	if len(commitEvents) > 0 {
		fmt.Printf("🔹 COMMIT PLUGIN - OTHER (%d events)\n", len(commitEvents))
		fmt.Printf("───────────────────────────────────\n")
		mf.printEvents(commitEvents)
		fmt.Println()
	}

	if len(execObsEvents) > 0 {
		fmt.Printf("🔸 EXECUTE PLUGIN - OBSERVATION PHASE (%d events)\n", len(execObsEvents))
		fmt.Printf("────────────────────────────────────────────────\n")
		mf.printEvents(execObsEvents)
		fmt.Println()
	}

	if len(execOutcomeEvents) > 0 {
		fmt.Printf("🔸 EXECUTE PLUGIN - OUTCOME PHASE (%d events)\n", len(execOutcomeEvents))
		fmt.Printf("───────────────────────────────────────────────\n")
		mf.printEvents(execOutcomeEvents)
		fmt.Println()
	}

	if len(execReportEvents) > 0 {
		fmt.Printf("🔸 EXECUTE PLUGIN - REPORT PHASE (%d events)\n", len(execReportEvents))
		fmt.Printf("──────────────────────────────────────────────\n")
		mf.printEvents(execReportEvents)
		fmt.Println()
	}

	if len(execEvents) > 0 {
		fmt.Printf("🔸 EXECUTE PLUGIN - OTHER (%d events)\n", len(execEvents))
		fmt.Printf("────────────────────────────────────\n")
		mf.printEvents(execEvents)
		fmt.Println()
	}

	if len(otherEvents) > 0 {
		fmt.Printf("📋 OTHER ACTIVITIES (%d events)\n", len(otherEvents))
		fmt.Printf("─────────────────────────────────────\n")
		mf.printEvents(otherEvents)
		fmt.Println()
	}

	// Calculate and display latency information
	if firstEvent != nil && lastEvent != nil {
		totalLatency := lastEvent.Timestamp.Sub(firstEvent.Timestamp)
		fmt.Printf("═══════════════════════════════════════════════════════════════\n")
		fmt.Printf("Summary:\n")
		fmt.Printf("  Total Events: %d\n", len(mf.events))
		fmt.Printf("  Start Time: %s\n", firstEvent.Timestamp.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("  End Time: %s\n", lastEvent.Timestamp.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("  Total Latency: %s\n", latencyStyle.Render(totalLatency.String()))

		fmt.Printf("\nStage Breakdown:\n")
		for stage, count := range stageCount {
			fmt.Printf("  %-20s: %d events\n", stage, count)
		}

		// Show tracked identifiers that were discovered
		if len(mf.merkleRoots) > 0 || len(mf.msgHashes) > 0 || len(mf.relatedSeqNums) > 0 || len(mf.ocrSequences) > 0 {
			fmt.Printf("\nTracked Identifiers:\n")
			if len(mf.ocrSequences) > 0 {
				fmt.Printf("  OCR Sequences: %d\n", len(mf.ocrSequences))
				var ocrSeqs []int
				for seq := range mf.ocrSequences {
					ocrSeqs = append(ocrSeqs, seq)
				}
				sort.Ints(ocrSeqs)
				for _, seq := range ocrSeqs {
					fmt.Printf("    %d\n", seq)
				}
			}
			if len(mf.merkleRoots) > 0 {
				fmt.Printf("  Merkle Roots: %d\n", len(mf.merkleRoots))
				for root := range mf.merkleRoots {
					fmt.Printf("    %s\n", root)
				}
			}
			if len(mf.relatedSeqNums) > 0 {
				fmt.Printf("  Message Sequence Numbers: %d\n", len(mf.relatedSeqNums))
				for seqNum := range mf.relatedSeqNums {
					fmt.Printf("    %s\n", seqNum)
				}
			}
			if len(mf.seqNumRanges) > 0 {
				fmt.Printf("  Sequence Number Ranges: %d\n", len(mf.seqNumRanges))
				for seqRange := range mf.seqNumRanges {
					fmt.Printf("    %s\n", seqRange)
				}
			}
			if len(mf.msgHashes) > 0 {
				fmt.Printf("  Message Hashes: %d\n", len(mf.msgHashes))
				for hash := range mf.msgHashes {
					maxLen := 50
					if len(hash) < maxLen {
						maxLen = len(hash)
					}
					fmt.Printf("    %s...\n", hash[:maxLen])
				}
			}
		}

		// Show phase-specific latencies
		if len(mf.phaseTimestamps) > 1 {
			fmt.Printf("\nPhase Latencies:\n")
			var phases []string
			for phase := range mf.phaseTimestamps {
				phases = append(phases, phase)
			}
			sort.Strings(phases)

			for i, phase := range phases {
				if i > 0 {
					prevPhase := phases[i-1]
					latency := mf.phaseTimestamps[phase].Sub(mf.phaseTimestamps[prevPhase])
					fmt.Printf("  %s → %s: %s\n", prevPhase, phase, latencyStyle.Render(latency.String()))
				}
			}
		}

		// Check for potential issues
		if stageCount["ERROR"] > 0 {
			fmt.Printf("\n%s: %d error events detected!\n", errorStyle.Render("WARNING"), stageCount["ERROR"])
		}

		if totalLatency > 5*time.Minute {
			fmt.Printf("\n%s: High latency detected (>5 minutes)\n", errorStyle.Render("WARNING"))
		}
	}

	return nil
}

func (mf *messageFlowFormatter) printEvents(events []MessageEvent) {
	for _, event := range events {
		// Format timestamp
		timeStr := timestampStyle.Render(event.Timestamp.Format("15:04:05.000"))

		// Format stage with proper width
		stageStr := stageStyle.Render(fmt.Sprintf("%-35s", event.Stage))

		// Format message
		msgStr := messageStyle.Render(event.Message)

		fmt.Printf("%s %s %s\n", timeStr, stageStr, msgStr)

		// Print relevant details with OCR info highlighted
		if len(event.Details) > 0 {
			// Show OCR phase and state first if available
			if ocrPhase, ok := event.Details["ocrPhase"].(string); ok {
				fmt.Printf("    📊 OCR Phase: %s", latencyStyle.Render(ocrPhase))
				if state, ok := event.Details["state"].(string); ok {
					fmt.Printf(" | State: %s", latencyStyle.Render(state))
				}
				if seqNr, ok := event.Details["ocrSeqNr"]; ok {
					fmt.Printf(" | SeqNr: %v", seqNr)
				}
				fmt.Println()
			}

			// Show merkle root information prominently
			if commitRoot, ok := event.Details["commitRoot"].(string); ok {
				fmt.Printf("    🌳 Commit Root: %s\n", commitRoot)
			}
			if computedRoot, ok := event.Details["computedRoot"].(string); ok {
				fmt.Printf("    🌳 Computed Root: %s\n", computedRoot)
			}
			if merkleRoot, ok := event.Details["merkleRoot"].(string); ok {
				fmt.Printf("    🌳 Merkle Root: %s\n", merkleRoot)
			}
			if root, ok := event.Details["root"].(string); ok {
				fmt.Printf("    🌳 Root: %s\n", root)
			}

			// Show other relevant details
			for key, value := range event.Details {
				if value != nil && value != "" && key != "ocrPhase" && key != "state" && key != "ocrSeqNr" &&
					key != "commitRoot" && key != "computedRoot" && key != "merkleRoot" && key != "root" {
					switch key {
					case "messageID", "msgId", "message_id":
						fmt.Printf("    🆔 %s: %v\n", key, value)
					case "seqNum", "sequenceNumber":
						fmt.Printf("    🔢 %s: %v\n", key, value)
					case "sourceChain", "destChain":
						fmt.Printf("    🔗 %s: %v\n", key, value)
					case "hash":
						fmt.Printf("    🌳 Hash: %v\n", value)
					default:
						fmt.Printf("    %s: %v\n", key, value)
					}
				}
			}
		}
	}
}
