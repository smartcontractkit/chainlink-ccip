// Package basic is a simple renderer that prints the data message to the
// console. Long lines are truncated when printing to a terminal.
package basic

import (
	"fmt"

	"golang.org/x/term"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render"
)

func init() {
	render.Register("basic", basicRendererFactory, "Print logs to the console with minimal processing.")

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
	return render.NewWrappedRender(basicRenderer)
}

func basicRenderer(data *parse.Data) {
	tryUpdateTermWidth()

	var line string
	if len(data.GetMessage()) > 0 {
		line = data.GetMessage()
	} else {
		line = fmt.Sprintf("%v", data)
	}

	// Truncate line if it's too long.
	if termWidth != 0 && len(line) > termWidth {
		line = line[:termWidth]
	}

	fmt.Println(line)
}
