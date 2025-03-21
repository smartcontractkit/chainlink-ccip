package basic

import (
	"fmt"

	"golang.org/x/term"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render"
)

func init() {
	render.Register("basic", basicRendererFactory)

	tryUpdateTermWidth()
}

var termWidth int = 80

func tryUpdateTermWidth() {
	if !term.IsTerminal(0) {
		// No width
		termWidth = 0
	}
	width, _, err := term.GetSize(0)
	if err != nil {
		return
	}

	termWidth = width
}

func basicRendererFactory(options render.Options) render.Renderer {
	return basicRenderer
}

func basicRenderer(data *parse.Data) {
	tryUpdateTermWidth()

	// Truncate line if it's too long.
	line := fmt.Sprintf("%v", data)
	if termWidth != 0 && len(line) > termWidth {
		line = line[:termWidth]
	}

	fmt.Println(line)
}
