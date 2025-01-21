package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v3"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/stream"
)

type arguments struct {
	files          []string
	logType        string
	disableFilters bool
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
	var uidStyle = lipgloss.NewStyle().Width(25).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1).Bold(true)
	var levelStyle = lipgloss.NewStyle().Width(4).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1).Italic(true)
	var messageStyle = lipgloss.NewStyle().Width(60).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1)
	var fieldsStyle = lipgloss.NewStyle().Width(100).Height(1).MaxHeight(1).
		Align(lipgloss.Left).PaddingLeft(1)

	uid := fmt.Sprintf("%s.%s.%s.%s.%s",
		withColor(data.OracleID, data.OracleID),
		withColor(data.DONID, data.DONID),
		withColor(data.SequenceNumber, data.SequenceNumber),
		withColor(data.Component, 0),
		withColor(data.OCRPhase, ocrPhaseToColor(data.OCRPhase)),
	)

	fmt.Printf("%s|%s|%s|%s|%s\n",
		timeStyle.Render(data.Timestamp.Format(time.TimeOnly)),
		uidStyle.Render(uid),
		levelStyle.Render(truncateLevel(data.Level)),
		messageStyle.Render(data.Message),
		fieldsStyle.Render(getRelevantFieldsForMessage(data)),
	)
}

func ocrPhaseToColor(phase string) int {
	switch phase {
	case "qry":
		return 1
	case "obs":
		return 2
	case "otcm":
		return 3
	case "rprt":
		return 4
	case "sacc":
		return 5
	case "strn":
		return 6
	default:
		return 0
	}
}

func truncateLevel(level string) string {
	switch lv := strings.ToLower(level); lv {
	case "info":
		return "ifo"
	case "debug":
		return "dbg"
	case "warn":
		return "wrn"
	case "error":
		return "err"
	case "critical":
		return "crt"
	default:
		return "unk"
	}
}

func getRelevantFieldsForMessage(data *parse.Data) string {
	var fields string

	if strings.ToLower(data.Level) == "error" {
		fields = fmt.Sprintf("err=%v", data.RawLoggerFields["err"])
	}

	if strings.HasPrefix(data.Message, "call to MsgsBetweenSeqNums returned unexpected") {
		return fmt.Sprintf(
			"%s expected=%v actual=%v chain=%v",
			fields,
			data.RawLoggerFields["expected"],
			data.RawLoggerFields["actual"],
			data.RawLoggerFields["chain"],
		)
	}
	if strings.HasPrefix(data.Message, "queried messages between sequence numbers") {
		return fmt.Sprintf("%s numMsgs=%v sourceChain=%v seqNumRange=%v",
			fields,
			data.RawLoggerFields["numMsgs"],
			data.RawLoggerFields["sourceChainSelector"],
			data.RawLoggerFields["seqNumRange"],
		)
	}
	if strings.HasPrefix(data.Message, "decoded messages between sequence numbers") {
		return fmt.Sprintf("%s sourceChain=%v seqNumRange=%v",
			fields, data.RawLoggerFields["sourceChainSelector"], data.RawLoggerFields["seqNumRange"])
	}

	return ""
}

func run(args arguments) error {
	var io stream.InputOptions

	// If no files are provided the stream will read from stdin.
	if len(args.files) != 0 {
		io.Filenames = args.files
	}

	fmt.Println("processing files", io.Filenames)

	inputStream, err := stream.InitializeInputStream(io)
	if err != nil {
		return fmt.Errorf("failed to initialize input stream: %w", err)
	}

	scanner := bufio.NewScanner(inputStream)
	for scanner.Scan() {
		line := scanner.Text()
		data, err := parse.Filter(line, args.logType, args.disableFilters)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get data: %s\n", err)
			return err
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
			&cli.StringFlag{
				Name:        "logType",
				Usage:       "Specify the type of log to parse, valid options: json, mixed, ci",
				Destination: &args.logType,
				Required:    true,
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
			&cli.BoolFlag{
				Name:        "disableFilters",
				Usage:       "Set to disable filter application on the logs. Defaults to false.",
				Destination: &args.disableFilters,
				Required:    false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return run(args)
		},
	}
}
