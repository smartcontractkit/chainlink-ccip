package parse

import "regexp"

// FilterFields is a struct that holds the filters for the data.
type FilterFields struct {
	Filters          []string
	ComponentFilters []string
}

func (f *FilterFields) Compile() (CompiledFilterFields, error) {
	compile := func(filters []string) ([]regexp.Regexp, error) {
		compiledFilters := make([]regexp.Regexp, len(filters))
		for i, filter := range filters {
			compiled, err := regexp.Compile(filter)
			if err != nil {
				return nil, err
			}
			compiledFilters[i] = *compiled
		}
		return compiledFilters, nil
	}

	var cff CompiledFilterFields
	var err error

	cff.Filters, err = compile(f.Filters)
	if err != nil {
		return CompiledFilterFields{}, err
	}
	cff.ComponentFilters, err = compile(f.ComponentFilters)
	if err != nil {
		return CompiledFilterFields{}, err
	}

	return cff, nil
}

// CompiledFilterFields is a struct that holds the compiled filters for the data.
type CompiledFilterFields struct {
	Filters          []regexp.Regexp
	ComponentFilters []regexp.Regexp
}

// Filter decides if the data should be displayed based on the provided filters.
func Filter(data *Data, filters CompiledFilterFields) (bool, error) {
	// No filters, include all data
	if len(filters.Filters) == 0 && len(filters.ComponentFilters) == 0 {
		return true, nil
	}

	// Check component filters
	component := data.Component
	for _, componentFilter := range filters.ComponentFilters {
		if componentFilter.MatchString(component) {
			return true, nil
		}
	}

	return false, nil
}
