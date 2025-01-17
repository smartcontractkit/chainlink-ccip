package parse

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	LogTypeJSON  = "json"
	LogTypeMixed = "mixed"
	LogTypeCI    = "ci"
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
type DataFilter func(data Data) *Data

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

// sanitizeString removes any unwanted characters from the string based on the log type.
func sanitizeString(s, logType string) string {
	if logType == "mixed" {
		// Mixed log lines are usually produced when running tests on the command line, they look like this:
		// logger.go:146: 2025-01-17T13:51:40.330+0200 INFO    CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e        commit/plugin.go:80     creating new plugin instance    {"version": "unset@unset", "plugin": "Commit", "oracleID": 1, "donID": 2, "configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291", "p2pID": "12D3KooWBD42agWRU3khVJwYQTXnr5uG1Qmh5n1Hm2q6RUFaJRhu"}
		return strings.TrimSpace(s)
	} else if logType == LogTypeCI {
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
	} else if logType == LogTypeJSON {
		return s
	} else {
		panic(fmt.Sprintf("SanitizeString: unknown log type %s", logType))
	}
}

func ParseLine(line, logType string) (*Data, error) {
	if logType == LogTypeJSON {
		var obj map[string]interface{}
		dec := json.NewDecoder(strings.NewReader(line))
		err := dec.Decode(&obj)
		if err != nil {
			return nil, fmt.Errorf("could not decode line from JSON (%s): %w", line, err)
		}

		var data Data
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			return nil, fmt.Errorf("unparsable line: %w", err)
		}

		// This isn't ours.
		if data.IsEmpty() {
			return nil, nil
		}

		data.RawLoggerFields = obj

		return &data, nil
	} else if logType == LogTypeMixed {
		// log line looks like this:
		// logger.go:146: 2025-01-17T13:51:40.330+0200 INFO    CCIPCommitPlugin.evm.1337.3379446385462418246.0x075f98f19ef9873523cde0267ab8b0253904363e        commit/plugin.go:80     creating new plugin instance    {"version": "unset@unset", "plugin": "Commit", "oracleID": 1, "donID": 2, "configDigest": "000a7d1df8632e2b3479350dcca1ee46eeec889dc37eb2ab094e63a1820ba291", "p2pID": "12D3KooWBD42agWRU3khVJwYQTXnr5uG1Qmh5n1Hm2q6RUFaJRhu"}
		// From the JSON object we can get all the log fields
		// From the remainer of the log we can get:
		// - timestamp
		// - level
		// - logger name
		// - caller
		// - message
		// write a regular expression for this
		re := regexp.MustCompile(`(?P<timestamp>\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}(?:\+\d{4})?)\s+(?P<level>\w+)\s+(?P<logger>\S+)\s+(?P<caller>\S+)\s+(?P<message>.*?)\s+(?P<jsonFields>{.*})`)
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) == 0 {
			return nil, fmt.Errorf("could not parse line: %s", line)
		}

		// extract all the named matches
		obj := make(map[string]string)
		for _, match := range matches {
			for i, name := range re.SubexpNames() {
				if i != 0 && name != "" {
					obj[name] = match[i]
				}
			}
		}

		// assert that all named matches have been found
		if len(obj) != 6 {
			return nil, fmt.Errorf("could not parse line: %s", line)
		}

		// construct the data object
		var data Data

		// some fields in the mixed log need to be manually populated in the Data struct
		parsedTs, err := parseCustomLayout(obj["timestamp"])
		if err != nil {
			return nil, fmt.Errorf("could not parse timestamp: %w", err)
		}
		data.Timestamp = parsedTs
		data.Level = obj["level"]
		data.LoggerName = obj["logger"]
		data.Caller = obj["caller"]
		data.Message = obj["message"]

		// parse the json fields into the data object to get the remaining fields.
		if err := json.Unmarshal([]byte(obj["jsonFields"]), &data); err != nil {
			return nil, fmt.Errorf("could not parse json fields: %w, fields: %s", err, obj["jsonFields"])
		}

		// parse the json fields into a map[string]any so that we can have the raw
		// fields in the data struct.
		var rawFields = make(map[string]any)
		if err := json.Unmarshal([]byte(obj["jsonFields"]), &rawFields); err != nil {
			return nil, fmt.Errorf("could not parse json fields: %w, fields: %s", err, obj["jsonFields"])
		}

		data.RawLoggerFields = rawFields

		return &data, nil
	} else if logType == LogTypeCI {
		return nil, fmt.Errorf("CI log parsing not yet implemented")
	} else {
		return nil, fmt.Errorf("unknown log type %s", logType)
	}
}

func Filter(line, logType string) (*Data, error) {
	line = sanitizeString(line, logType)
	if len(line) == 0 {
		return nil, nil
	}

	data, err := ParseLine(line, logType)
	if err != nil {
		return nil, fmt.Errorf("ParseLine: %w", err)
	}

	for _, f := range filters {
		data := f.df(*data)
		if data != nil {
			data.FilterName = f.name
			return data, nil
		}
		// TODO: multiple matches?
	}
	return data, nil
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

	RawLoggerFields map[string]any `json:"-"`

	// Additional detail space, can be unique to each filter.
	// i.e. an error message, observer details, number of messages, etc
	Details string
}

func (d Data) IsEmpty() bool {
	return false // TODO: implement
}

func parseCustomLayout(s string) (time.Time, error) {
	const customLayout = "2006-01-02T15:04:05.000-0700"
	return time.Parse(customLayout, s)
}
