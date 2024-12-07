package main

import (
	"bufio"
	"fmt"
	"os"

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

func renderData(data *parse.Data) {
	var timeStyle = lipgloss.NewStyle().Width(19).Height(1)
	// 2024-12-04T20:11:51.997Z |

	out :=
		timeStyle.Render(data.Timestamp)
	fmt.Println(data)
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
