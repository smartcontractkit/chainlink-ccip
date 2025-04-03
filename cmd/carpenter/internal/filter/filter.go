//go:generate go-enum -f=$GOFILE --names --nocase

package filter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

// ENUM(Plugin, Component, LogLevel, Message, Caller, LoggerName, DONID, SequenceNumber)
type Field string

type matcher struct {
	antiMatcher bool
	re          *regexp.Regexp
}

// CompiledFilterFields is a collection of compiled Field filters.
type CompiledFilterFields map[Field][]matcher

// ENUM(AND, OR)
type FilterOP string

func NewFilterFields(rawFields []string) (CompiledFilterFields, error) {
	fields := make(CompiledFilterFields)
	var malformedFields []string
	var invalidFields []string
	var errs []error
	for _, field := range rawFields {
		antiMatch := false
		if strings.HasPrefix(field, "!") {
			antiMatch = true
			field = field[1:]
		}

		if !strings.Contains(field, ":") {
			malformedFields = append(malformedFields, field)
			continue
		}

		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			malformedFields = append(malformedFields, field)
			continue
		}

		f, err := ParseField(parts[0])
		if err != nil {
			invalidFields = append(invalidFields, parts[0])
		}

		re, err := regexp.Compile(parts[1])
		if err != nil {
			errs = append(errs,
				fmt.Errorf("could not compile regexp %s in %s: %w", parts[1], field, err))
		}

		fields[f] = append(fields[f], matcher{antiMatcher: antiMatch, re: re})
	}

	if len(malformedFields) > 0 {
		errs = append(errs,
			fmt.Errorf("malformed fields: %s, expected format: field:pattern",
				strings.Join(malformedFields, ", ")))
	}
	if len(invalidFields) > 0 {
		errs = append(errs,
			fmt.Errorf("invalid fields: %s not in [%s]",
				strings.Join(invalidFields, ", "),
				strings.Join(FieldNames(), ", ")))
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return fields, nil
}

// Filter decides if the data should be displayed based on the provided filters.
func Filter(data *parse.Data, filters CompiledFilterFields, op FilterOP) (bool, error) {
	// No filters, include by default
	if len(filters) == 0 {
		return true, nil
	}

	anyMatch := false
	allMatch := true

	for field, compiledFilters := range filters {
		var fieldStr string
		for _, compiledFilter := range compiledFilters {
			switch field {
			case FieldComponent:
				fieldStr = data.Component
			case FieldMessage:
				fieldStr = data.GetMessage()
			case FieldLogLevel:
				fieldStr = data.GetLevel()
			case FieldCaller:
				fieldStr = data.GetCaller()
			case FieldLoggerName:
				fieldStr = data.GetLoggerName()
			case FieldPlugin:
				fieldStr = data.Plugin
			case FieldDONID:
				fieldStr = fmt.Sprintf("%d", data.DONID)
			case FieldSequenceNumber:
				fieldStr = fmt.Sprintf("%d", data.SequenceNumber)
			}

			matches := compiledFilter.re.MatchString(fieldStr)
			if compiledFilter.antiMatcher {
				if matches {
					return false, nil
				} else {
					continue
				}
			}

			anyMatch = anyMatch || matches
			allMatch = allMatch && matches
		}
	}

	switch op {
	case FilterOPAND:
		return allMatch, nil
	case FilterOPOR:
		return anyMatch, nil
	default:
		return anyMatch, nil
	}
}
