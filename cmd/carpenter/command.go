package main

import (
	"bufio"
	"context"
	"fmt"
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
	logType      string
	rendererName string

	filter.CompiledFilterFields
}

func run(args arguments) error {
	var io stream.InputOptions

	// If no files are provided the stream will read from stdin.
	if len(args.files) != 0 {
		io.Filenames = args.files
	}

	renderer, err := render.GetRenderer(args.rendererName, render.Options{})
	if err != nil {
		return fmt.Errorf("failed to get renderer: %w", err)
	}

	inputStream, err := stream.InitializeInputStream(io)
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

		include, err := filter.Filter(data, args.CompiledFilterFields)
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

		renderer(data)
	}
	return nil
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
				Name:        "logType",
				Usage:       "Specify the type of log to parse, valid options: json, mixed, ci",
				Destination: &args.logType,
				Value:       "json",
				Validator: func(s string) error {
					if !parse.IsValidLogType(s) {
						return fmt.Errorf("invalid log type: %s, expected either %s or %s or %s",
							s,
							parse.LogTypeJSON,
							parse.LogTypeMixed,
							parse.LogTypeMixedGoTestJSON,
						)
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
						return fmt.Errorf("invalid renderer: %s, expected one of %s", s, choices)
					}
					return nil
				},
			},
			&cli.StringSliceFlag{
				Name:  "filter",
				Usage: fmt.Sprintf("Provide one or more filters to apply to the logs. Format as 'FieldName:Regexp', valid fields: [%s]", strings.Join(filter.FieldNames(), ", ")),
				//Destination: &args.FilterFields.Filters,
				Category: "filters",
				Action: func(ctx context.Context, command *cli.Command, fields []string) error {
					var err error
					args.CompiledFilterFields, err = filter.NewFilterFields(fields)
					if err != nil {
						return fmt.Errorf("invalid filter fields: %w", err)
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
