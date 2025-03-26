package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/filter"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/stream"
)

type arguments struct {
	files        []string
	logType      parse.LogType
	rendererName string

	filter.CompiledFilterFields
	filterOP filter.FilterOP
}

func makeCommand() *cli.Command {
	var args arguments
	return &cli.Command{
		Name:  "carpenter",
		Usage: "A tool for parsing and displaying logs",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:        "filename",
				Usage:       "Provide one or more files to read. If not provided, reads from stdin.",
				Destination: &args.files,
			},
			&cli.StringFlag{
				Name:             "logType",
				Usage:            "Specify the type of log to parse, valid options: json, mixed, ci",
				Value:            parse.LogTypeJSON.String(),
				ValidateDefaults: true, // to make sure default is assigned.
				Validator: func(s string) error {
					var err error
					args.logType, err = parse.ParseLogType(s)
					if err != nil {
						return fmt.Errorf("expected one of [%s]",
							strings.Join(parse.LogTypeNames(), ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				OnlyOnce:    true,
				Name:        "renderer",
				Usage:       fmt.Sprintf("Select which rendering algorithm to use: [%s]", strings.Join(render.GetRenderers(), ", ")),
				Value:       "basic",
				Destination: &args.rendererName,
				Validator: func(s string) error {
					choices := render.GetRenderers()
					if !slices.Contains(choices, s) {
						return fmt.Errorf("expected one of [%s]",
							s, strings.Join(choices, ", "))
					}
					return nil
				},
			},
			&cli.StringSliceFlag{
				Name: "filter",
				Usage: fmt.Sprintf(
					"Line selection filters. Format as '[!]FieldName:Regexp', the optional ! prefix will omit logs that match the pattern, valid fields: [%s]",
					strings.Join(filter.FieldNames(), ", ")),
				Category: "filters",
				Validator: func(fields []string) error {
					var err error
					args.CompiledFilterFields, err = filter.NewFilterFields(fields)
					if err != nil {
						return err
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name: "filter-op",
				Usage: fmt.Sprintf(
					"Operation to use when combining filters. Valid options: [%s]",
					strings.Join(filter.FilterOPNames(), ", ")),
				Category:         "filters",
				Value:            string(filter.FilterOPAND),
				ValidateDefaults: true, // to make sure default is assigned.
				Validator: func(s string) error {
					var err error
					args.filterOP, err = filter.ParseFilterOP(s)
					if err != nil {
						return fmt.Errorf("expected one of %s", err,
							strings.Join(filter.FilterOPNames(), ", "))
					}
					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return run(args)
		},
	}
}

func run(args arguments) error {
	var options stream.InputOptions

	// If no files are provided the stream will read from stdin.
	if len(args.files) != 0 {
		options.Filenames = args.files
	}

	renderer, err := render.GetRenderer(args.rendererName, render.Options{})
	if err != nil {
		return fmt.Errorf("failed to get renderer: %w", err)
	}

	inputStream, err := stream.InitializeInputStream(options)
	if err != nil {
		return fmt.Errorf("failed to initialize input stream: %w", err)
	}

	scanner := bufio.NewScanner(inputStream)
	for scanner.Scan() {
		line := scanner.Text()
		data, err := parse.ParseLine(line, args.logType)
		if err != nil {
			return fmt.Errorf("ParseLine: %w", err)
		}

		include, err := filter.Filter(data, args.CompiledFilterFields, args.filterOP)
		if err != nil {
			msg := fmt.Sprintf("Unable to get data: %s\n", err)
			_, err2 := fmt.Fprintf(os.Stderr, msg)
			if err2 != nil {
				panic(msg)
			}
			return err
		}
		if !include {
			// no data to display.
			continue
		}

		renderer.Render(data)
	}

	// Check if renderer implements io.Closer and call Close if it does
	if closer, ok := renderer.(io.Closer); ok {
		if err := closer.Close(); err != nil {
			return fmt.Errorf("failed to close renderer: %w", err)
		}
	}

	return nil
}
