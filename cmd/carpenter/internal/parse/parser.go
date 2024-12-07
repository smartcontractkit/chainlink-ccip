package parse

import (
	"encoding/json"
	"fmt"
	"strings"
)

// DataFilter is used by Filter to identify lines that should be displayed.
type DataFilter func(object map[string]interface{}) *Data

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

func ParseLine(line string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return nil, nil
	}

	dec := json.NewDecoder(strings.NewReader(line))
	err := dec.Decode(&obj)
	if err != nil {
		return nil, fmt.Errorf("could not decode line from JSON (%s): %w", line, err)
	}

	return obj, nil
}

func Filter(line string) (*Data, error) {
	object, err := ParseLine(line)
	if err != nil {
		return nil, fmt.Errorf("ParseLine: %w", err)
	}

	for _, f := range filters {
		data := f.df(object)
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
	Timestamp       string
	Level           string
	Caller          string
	SequenceNumber  int
	Plugin          string
	PluginProcessor string

	// Additional detail space, can be unique to each filter.
	// i.e. an error message, observer details, number of messages, etc
	Details string
}
