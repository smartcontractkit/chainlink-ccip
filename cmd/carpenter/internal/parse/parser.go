package parse

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

/*
Ideal log format

{
	// global log-level information
	"timestamp",
	"level",
	"caller", // package/file:linenum
	"logger", // name, i.e. CCIPCCIPExecOCR3.evm.1337.0xe6e...
	"version", // 2.18.0@732cc15
	"msg", // Unstructured log message.

	// app specific fields also go at this level.
	"oid",
	"oracleID",
	"donID",
	"seqNr",
	"e",
	"l"
	"proto",
	...

	// New - organize all chainlink-ccip specific data into this field.
	"plugin": {
		// static data here, automatically included by logger enhancements.
		"oid",
		"donID",
		"seqNum",
		"plugin", // commit|execute
		"processor", // merkleRoot, discovery, etc
		"logger stuff", // split up CCIPCCIPExecOCR3.evm.1337.0xe6e...

		"type", // used to serialize data
		"event", // a go struct 'lggr.LogEvent(CommitObservationLog{obs: observation})'
	},
}
*/

// DataFilter is used by Filter to identify lines that should be displayed.
type DataFilter func(data Data, object map[string]interface{}) *Data

type filter struct {
	name string
	df   DataFilter
}

var filters []filter

// RegisterDataFilter adds a filter
func RegisterDataFilter(name string, df DataFilter) {
	filters = append(filters, filter{
		name: name,
		df:   df,
	})
}

// sanitizeString removes any unwanted characters from the string which may
// be included by CI.
func sanitizeString(s string) string {
	if len(s) > 0 && s[0] != '{' {
		// Look for embedded tab
		idx := strings.LastIndex(s, `\t`)
		if idx != -1 {
			s = s[idx+2:]
			// Look for newline
			idx = strings.Index(s, `\n`)
			if idx != -1 {
				s = s[:idx]
			}
		}
	}

	if len(s) > 0 && s[0] != '{' {
		return ""
	}

	s = strings.ReplaceAll(s, "\\", "")
	s = strings.TrimSpace(s)
	return s
}

func ParseLine(line string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(line))
	err := dec.Decode(&obj)
	if err != nil {
		return nil, fmt.Errorf("could not decode line from JSON (%s): %w", line, err)
	}

	return obj, nil
}

func Filter(line string) (*Data, error) {
	line = sanitizeString(line)

	if len(line) == 0 {
		return nil, nil
	}

	object, err := ParseLine(line)
	if err != nil {
		return nil, fmt.Errorf("ParseLine: %w", err)
	}

	var data Data
	if err := json.Unmarshal([]byte(line), &data); err != nil {
		return nil, fmt.Errorf("unparsable line: %w", err)
	}

	// This isn't ours.
	if (data == Data{}) {
		return nil, nil
	}

	for _, f := range filters {
		data := f.df(data, object)
		if data != nil {
			data.FilterName = f.name
			return data, nil
		}
		// TODO: multiple matches?
	}
	return nil, nil
}

type Data struct {
	// FilterName that generated the data.
	FilterName string

	// TODO: automatically parse the global fields:
	// level, ts, logger, caller, msg, version, donID, oracleID
	// Data that we expect from most logs.
	// This is all part of the primary display.
	LoggerName     string    `json:"logger"`
	Timestamp      time.Time `json:"ts"`
	Level          string    `json:"level"`
	Caller         string    `json:"caller"`
	SequenceNumber int       `json:"ocrSeqNr"`
	Plugin         string    `json:"plugin"`
	Component      string    `json:"component"`
	OracleID       int       `json:"oracleID"`
	DONID          int       `json:"donID"`
	Message        string    `json:"msg"`
	Version        string    `json:"version"`
	ConfigDigest   string    `json:"configDigest"`

	// Additional detail space, can be unique to each filter.
	// i.e. an error message, observer details, number of messages, etc
	Details string
}
