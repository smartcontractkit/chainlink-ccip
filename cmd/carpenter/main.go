package main

import (
	"context"
	"fmt"
	"os"

	// Register the formatters
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/basic"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/fancy"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format/summary"
)

func main() {
	err := makeCommand().Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem running command: %s\n", err.Error())
	}
}
