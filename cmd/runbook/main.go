// Command runbook executes the machine-checkable YAML specs embedded in
// docs/runbooks deterministically against a Prometheus-compatible datasource.
//
// Its purpose is to remove the mechanical walk (copy-paste the queries,
// apply the empty-result rule, follow the branches, format the output
// contract) from the AI-agent loop, so an agent only spends tokens on the
// judgment that actually remains. Every executed query and raw result is
// emitted so the AI layer (or a human) can still verify its interpretation
// against ground truth.
package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	cmd := makeCommand()
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "runbook: %s\n", err)
		os.Exit(1)
	}
}
