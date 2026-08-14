package main

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
	"gopkg.in/yaml.v3"

	"github.com/smartcontractkit/chainlink-ccip/cmd/runbook/internal"
)

const defaultEndpoint = "http://localhost:8428"

type arguments struct {
	endpoint string
	bearer   string
	user     string
	pass     string
	suffix   string
	timeout  time.Duration
	raw      bool
	defines  []string
}

func makeCommand() *cli.Command {
	var args arguments
	return &cli.Command{
		Name:  "runbook",
		Usage: "Execute the machine-checkable YAML specs behind docs/runbooks against a Prometheus-compatible datasource, deterministically.",
		Commands: []*cli.Command{
			{
				Name:   "list",
				Usage:  "List the runbooks bundled into this binary.",
				Action: func(ctx context.Context, cmd *cli.Command) error { return listRunbooks() },
			},
			{
				Name:      "run",
				Usage:     "Run a runbook and emit its output contract.",
				UsageText: "runbook run [runbook] -D input=value ...",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					name := cmd.Args().First()
					if name == "" {
						return fmt.Errorf("runbook run requires a runbook name (see 'runbook list')")
					}
					return run(ctx, name, &args)
				},
			},
		},
		Flags: runFlags(&args),
	}
}

func runFlags(args *arguments) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint",
			Value:       defaultEndpoint,
			Usage:       "Prometheus-compatible datasource base URL (e.g. VictoriaMetrics).",
			Destination: &args.endpoint,
		},
		&cli.StringFlag{
			Name:        "bearer",
			Usage:       "Bearer token for the datasource.",
			Destination: &args.bearer,
		},
		&cli.StringFlag{
			Name:        "user",
			Usage:       "Basic-auth username for the datasource.",
			Destination: &args.user,
		},
		&cli.StringFlag{
			Name:        "pass",
			Usage:       "Basic-auth password for the datasource.",
			Destination: &args.pass,
		},
		&cli.StringFlag{
			Name:        "suffix",
			Value:       "auto",
			Aliases:     []string{"s"},
			Usage:       "Metric-name suffix policy: auto (probe), keep (as-written _total/_bucket), strip (base names).",
			Destination: &args.suffix,
		},
		&cli.DurationFlag{
			Name:        "timeout",
			Value:       20 * time.Second,
			Usage:       "Per-query timeout.",
			Destination: &args.timeout,
		},
		&cli.BoolFlag{
			Name:        "raw",
			Usage:       "Emit raw query results alongside verdicts (for agent/human verification).",
			Destination: &args.raw,
		},
		&cli.StringSliceFlag{
			Name:        "define",
			Aliases:     []string{"D"},
			Usage:       "Runbook input as name=value. Repeatable. See 'runbook list'/YAML for each runbook's inputs.",
			Destination: &args.defines,
		},
	}
}

func listRunbooks() error {
	names, err := internal.EmbeddedRunbooks()
	if err != nil {
		return err
	}
	for _, n := range names {
		rb, err := internal.LoadRunbook(strings.TrimSuffix(n, ".yaml"))
		if err != nil {
			fmt.Printf("  %s\t— (unreadable: %v)\n", n, err)
			continue
		}
		var ins []string
		for _, in := range rb.Inputs {
			ins = append(ins, in.Name)
		}
		fmt.Printf("  %-26s %s\n      inputs: %s\n", rb.Name, rb.Description, strings.Join(ins, ", "))
	}
	return nil
}

func run(ctx context.Context, name string, args *arguments) error {
	rb, err := internal.LoadRunbook(name)
	if err != nil {
		return err
	}

	ex := internal.NewExecutor(strings.TrimSuffix(args.endpoint, "/"), args.timeout)
	if args.bearer != "" {
		ex.SetBearer(args.bearer)
	}
	if args.user != "" {
		ex.SetBasicAuth(args.user, args.pass)
	}
	switch args.suffix {
	case "auto":
		// nil -> probe on first query
	case "keep", "total":
		ex.SetStrip(false)
	case "strip", "none":
		ex.SetStrip(true)
	default:
		return fmt.Errorf("unknown --suffix %q (auto|keep|strip)", args.suffix)
	}

	provided, err := parseDefines(args.defines)
	if err != nil {
		return err
	}
	in, err := internal.NewInputs(rb, provided)
	if err != nil {
		return err
	}

	switch rb.Type {
	case "checklist":
		rep, err := internal.RunChecklist(ctx, rb, ex, in)
		if err != nil {
			return err
		}
		return emitYAML(rep, args.raw)
	case "graph":
		rep, err := internal.RunGraph(ctx, rb, ex, in, 30)
		if err != nil {
			return err
		}
		return emitYAML(rep, args.raw)
	default:
		return fmt.Errorf("unsupported runbook type %q", rb.Type)
	}
}

func parseDefines(flags []string) (map[string]string, error) {
	out := map[string]string{}
	for _, d := range flags {
		k, v, ok := strings.Cut(d, "=")
		if !ok || k == "" {
			return nil, fmt.Errorf("--define must be name=value, got %q", d)
		}
		out[k] = v
	}
	return out, nil
}

// emitYAML writes the report (optionally with per-variable raw QueryEvidence
// keyed under a _evidence map so an agent can still verify a verdict against
// ground truth without ICEReparsing the verdict).
func emitYAML(v any, raw bool) error {
	wrap := struct {
		Report   any            `yaml:"report"`
		Evidence map[string]any `yaml:"evidence,omitempty"`
	}{Report: v}
	if raw {
		wrap.Evidence = extractEvidence(v)
	}
	out, err := yaml.Marshal(wrap)
	if err != nil {
		return err
	}
	_, err = io.WriteString(writer(), string(out))
	return err
}
