// Package basic is a simple formatter that prints the data message to the
// console. Long lines are truncated when printing to a terminal.
package basic

import (
	"fmt"

	"golang.org/x/term"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	// Register the basic formatter by name.
	format.Register("basic", basicFormatterFactory, "Print logs to the console with minimal processing.")

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

func basicFormatterFactory(options format.Options) format.Formatter {
	// This simple formatter is a pure function, more advanced formats
	// can implement the Formatter interface.
	return format.NewWrappedFormat(basicFormatter)
}

func basicFormatter(data *parse.Data) {
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
