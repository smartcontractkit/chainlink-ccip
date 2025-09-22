package messageflow

import (
	"fmt"
	"sort"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	format.Register("discover-ids", discoverIDsFormatterFactory, "Discover available message/transaction IDs in the logs for tracking.")
}

func discoverIDsFormatterFactory(options format.Options) format.Formatter {
	return &discoverIDsFormatter{
		identifiers: make(map[string][]string),
	}
}

type discoverIDsFormatter struct {
	identifiers map[string][]string // field name -> list of unique values
}

func (d *discoverIDsFormatter) Format(data *parse.Data) {
	// Look for potential identifiers in raw logger fields
	if data.RawLoggerFields != nil {
		for key, value := range data.RawLoggerFields {
			// Look for fields that might contain identifiers
			if d.isIdentifierField(key) {
				if valueStr, ok := value.(string); ok && valueStr != "" {
					if len(d.identifiers[key]) < 20 { // Limit to first 20 unique values
						if !contains(d.identifiers[key], valueStr) {
							d.identifiers[key] = append(d.identifiers[key], valueStr)
						}
					}
				}
			}
		}
	}

	// Also check for identifiers in the message itself
	message := data.GetMessage()
	if strings.Contains(message, "messageID") || strings.Contains(message, "txHash") ||
		strings.Contains(message, "requestID") || strings.Contains(message, "seqNum") {
		// Extract potential IDs from the message
		words := strings.Fields(message)
		for i, word := range words {
			if (strings.Contains(word, "ID") || strings.Contains(word, "Hash") ||
				strings.Contains(word, "Num")) && i+1 < len(words) {
				nextWord := words[i+1]
				if len(nextWord) > 10 { // Likely an ID if it's long enough
					key := "extracted_" + word
					if len(d.identifiers[key]) < 10 {
						if !contains(d.identifiers[key], nextWord) {
							d.identifiers[key] = append(d.identifiers[key], nextWord)
						}
					}
				}
			}
		}
	}
}

func (d *discoverIDsFormatter) isIdentifierField(fieldName string) bool {
	lowerField := strings.ToLower(fieldName)
	identifierPatterns := []string{
		"messageid", "msgid", "message_id", "msg_id",
		"txhash", "tx_hash", "transactionhash", "transaction_hash",
		"requestid", "request_id", "req_id",
		"seqnum", "seq_num", "sequencenumber", "sequence_number",
		"chainid", "chain_id",
		"blocknum", "block_num", "blocknumber", "block_number",
		"logindex", "log_index",
		"id", "hash", "digest",
	}

	for _, pattern := range identifierPatterns {
		if strings.Contains(lowerField, pattern) {
			return true
		}
	}
	return false
}

func (d *discoverIDsFormatter) Close() error {
	if len(d.identifiers) == 0 {
		fmt.Println("No potential identifiers found in the logs.")
		fmt.Println("The logs may not contain CCIP message tracking information.")
		return nil
	}

	fmt.Println("Discovered Potential Identifiers:")
	fmt.Println("═════════════════════════════════")

	// Sort field names for consistent output
	var fieldNames []string
	for field := range d.identifiers {
		fieldNames = append(fieldNames, field)
	}
	sort.Strings(fieldNames)

	for _, field := range fieldNames {
		values := d.identifiers[field]
		fmt.Printf("\n%s (%d unique values):\n", field, len(values))
		for i, value := range values {
			if i >= 5 { // Show max 5 examples
				fmt.Printf("  ... and %d more\n", len(values)-5)
				break
			}
			fmt.Printf("  %s\n", value)
		}
	}

	fmt.Println("\nUsage:")
	fmt.Println("Use any of these identifiers with the messageflow formatter:")
	fmt.Println("  ./carpenter --format messageflow --message-id <identifier_value> ...")

	return nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}



