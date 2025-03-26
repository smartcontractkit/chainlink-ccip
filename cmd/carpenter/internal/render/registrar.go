package render

import (
	"fmt"
	"maps"
	"slices"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

// Options is a struct that holds options for all renderers.
type Options struct {
}

// RendererFactory is a function that returns a Renderer, implemented by renderers to apply options.
type RendererFactory func(options Options) Renderer

// Renderer is an interface that can render data somehow.
type Renderer interface {
	Render(data *parse.Data)
}

// rendererFactories is a map of renderers by name.
var rendererFactories = make(map[string]RendererFactory)

// Register registers a renderer with the given name.
func Register(name string, factory RendererFactory, description string) {
	// panic if factory is nil or already registered
	if factory == nil {
		panic("renderer is nil")
	}
	if _, ok := rendererFactories[name]; ok {

	}

	rendererFactories[name] = factory
}

// WrappedRender is a wrapper around a function that implements Renderer.
type wrappedRender struct {
	F func(data *parse.Data)
}

func (w wrappedRender) Render(data *parse.Data) {
	w.F(data)
}

func NewWrappedRender(f func(data *parse.Data)) Renderer {
	return wrappedRender{F: f}
}

// GetRenderer returns a renderer by name.
func GetRenderer(name string, options Options) (Renderer, error) {
	factory, ok := rendererFactories[name]
	if !ok {
		return nil, fmt.Errorf("renderer %s not found", name)
	}
	return factory(options), nil
}

func GetRenderers() []string {
	return slices.Collect(maps.Keys(rendererFactories))
}
