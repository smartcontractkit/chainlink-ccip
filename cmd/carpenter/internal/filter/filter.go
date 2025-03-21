//go:generate go-enum -f=$GOFILE --names --nocase

package filter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

// ENUM(Component, LogLevel, Message, Caller, LoggerName, )
type Field string

// CompiledFilterFields is a collection of compiled Field filters.
type CompiledFilterFields map[Field][]*regexp.Regexp

func NewFilterFields(rawFields []string) (CompiledFilterFields, error) {
	fields := make(CompiledFilterFields)
	var malformedFields []string
	var invalidFields []string
	var errs []error
	for _, field := range rawFields {
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

		fields[f] = append(fields[f], re)
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
func Filter(data *parse.Data, filters CompiledFilterFields) (bool, error) {
	// No filters, include by default
	if len(filters) == 0 {
		return true, nil
	}

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
			}

			if compiledFilter.MatchString(fieldStr) {
				return true, nil
			}
		}
	}
	return false, nil
}
