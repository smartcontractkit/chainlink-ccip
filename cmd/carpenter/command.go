package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/stream"
)

type arguments struct {
	files []string
}

// renderData
/*

2024-12-04T20:15:35Z | 1.1.1 |       Commit(MerkleRoot) | <processor details>
                       | | |           |     |-- processor
                       | | |           |-- OCR Plugin
                       | | |-- sequence number
                       | |-- DON ID
                       -- oracleID

*/
func renderData(data *parse.Data) {
	// simple color selection algorithm
	withColor := func(in interface{}, i int) string {
		color := fmt.Sprintf("%d", i%7+1)
		str := fmt.Sprintf("%v", in)

		return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(str)
	}

	var timeStyle = lipgloss.NewStyle().Width(10).Height(1).MaxHeight(1).
		Align(lipgloss.Center)
	var uidStyle = lipgloss.NewStyle().Width(15).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1).Bold(true)
	var messageStyle = lipgloss.NewStyle().Width(50).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1)

	uid := fmt.Sprintf("%s.%s.%s",
		withColor(data.OracleID, data.OracleID),
		withColor(data.DONID, data.DONID),
		withColor(data.SequenceNumber, data.SequenceNumber),
	)

	fmt.Printf("%s|%s|%s\n",
		timeStyle.Render(data.Timestamp.Format(time.TimeOnly)),
		uidStyle.Render(uid),
		messageStyle.Render(data.Message),
	)
}

func run(args arguments) error {
	var io stream.InputOptions

	// If no files are provided the stream will read from stdin.
	if len(args.files) != 0 {
		io.Filenames = args.files
	}

	inputStream, err := stream.InitializeInputStream(io)
	if err != nil {
		return fmt.Errorf("failed to initialize input stream: %w", err)
	}

	scanner := bufio.NewScanner(inputStream)
	for scanner.Scan() {
		line := scanner.Text()
		data, err := parse.Filter(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get data: %s\n", err)
			continue
		}
		if data == nil {
			// no data to display.
			continue
		}

		renderData(data)
		//fmt.Println(data)
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
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return run(args)
		},
	}
}
