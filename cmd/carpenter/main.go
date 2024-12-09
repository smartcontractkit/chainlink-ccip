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
	var timeStyle = lipgloss.NewStyle().Width(30).MaxWidth(30).Height(1).MaxHeight(1)

	out :=
		timeStyle.Render(data.Timestamp.Format(time.RFC3339))
	//fmt.Println(data)
	fmt.Println(out)
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
