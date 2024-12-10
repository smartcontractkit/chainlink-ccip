package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"

	_ "github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/filter"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/stream"
)

func maybeExit(err error, s string, a ...interface{}) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, s+"\n", a...)
	os.Exit(1)
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

func main() {
	inputStream, err := stream.InitializeInputStream(stream.InputOptions{})
	maybeExit(err, "failed to initialize input stream: %s", err)

	scanner := bufio.NewScanner(inputStream)
	for scanner.Scan() {
		line := scanner.Text()
		data, err := parse.Filter(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to get data: %s", err)
			continue
		}
		if data == nil {
			// no data to display.
			continue
		}

		renderData(data)
		// TODO: render data.
		//fmt.Println(data)
	}
}
