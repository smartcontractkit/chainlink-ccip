package format

import (
	"fmt"
	"maps"
	"slices"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

// Options is a struct that holds options for all formatters.
type Options struct {
}

// FormatterFactory is a function that returns a Formatter, implemented by formatter to apply options.
type FormatterFactory func(options Options) Formatter

// Formatter is an interface that can format data somehow.
type Formatter interface {
	Format(data *parse.Data)
}

// formatterFactories is a map of formatters by name.
var formatterFactories = make(map[string]FormatterFactory)

// Register registers a formatter with the given name.
func Register(name string, factory FormatterFactory, description string) {
	// panic if factory is nil or already registered
	if factory == nil {
		panic("formatter is nil")
	}
	if _, ok := formatterFactories[name]; ok {

	}

	formatterFactories[name] = factory
}

// WrappedFormat is a wrapper around a function that implements Formatter.
type wrappedFormat struct {
	F func(data *parse.Data)
}

func (w wrappedFormat) Format(data *parse.Data) {
	w.F(data)
}

func NewWrappedFormat(f func(data *parse.Data)) Formatter {
	return wrappedFormat{F: f}
}

// GetFormatter returns a formatter by name.
func GetFormatter(name string, options Options) (Formatter, error) {
	factory, ok := formatterFactories[name]
	if !ok {
		return nil, fmt.Errorf("formatter %s not found", name)
	}
	return factory(options), nil
}

func GetFormatters() []string {
	return slices.Collect(maps.Keys(formatterFactories))
}
