package main

import (
	"context"
	"fmt"
	"os"

	// Register the renderers
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render/basic"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render/fancy"
	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render/summary"
)

func main() {
	err := makeCommand().Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem running command: %s\n", err.Error())
	}
}
