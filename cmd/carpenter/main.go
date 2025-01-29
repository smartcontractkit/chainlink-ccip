package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/filter"
)

func main() {
	err := makeCommand().Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Problem running command: %s\n", err.Error())
	}
}
