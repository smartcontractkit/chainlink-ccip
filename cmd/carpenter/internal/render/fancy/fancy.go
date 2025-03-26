package fancy

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render"
)

func init() {
	render.Register("fancy", basicRendererFactory)
}

func basicRendererFactory(options render.Options) render.Renderer {
	return basicRenderer
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
func basicRenderer(data *parse.Data) {
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
		timeStyle.Render(data.GetTimestamp().Format(time.TimeOnly)),
		uidStyle.Render(uid),
		levelStyle.Render(truncateLevel(data.GetLevel())),
		messageStyle.Render(data.GetMessage()),
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

	if strings.ToLower(data.GetLevel()) == "error" {
		fields = fmt.Sprintf("err=%v", data.RawLoggerFields["err"])
	}

	if strings.HasPrefix(data.GetMessage(), "failed to get token prices outcome") {
		return fmt.Sprintf("err=%v", data.RawLoggerFields["err"])
	}

	if strings.HasPrefix(data.GetMessage(), "Get consensus observation failed, empty outcome") {
		return fmt.Sprintf("err=%v", data.RawLoggerFields["err"])
	}

	if strings.HasPrefix(data.GetMessage(), "Sending Outcome") {
		return fmt.Sprintf("nextState=%v outcome=%v",
			data.RawLoggerFields["nextState"], data.RawLoggerFields["outcome"])
	}

	if strings.HasPrefix(data.GetMessage(), "sending merkle root processor observation") {
		return fmt.Sprintf("observation=%v", data.RawLoggerFields["observation"])
	}

	if strings.HasPrefix(data.GetMessage(), "call to MsgsBetweenSeqNums returned unexpected") {
		return fmt.Sprintf(
			"%s expected=%v actual=%v chain=%v",
			fields,
			data.RawLoggerFields["expected"],
			data.RawLoggerFields["actual"],
			data.RawLoggerFields["chain"],
		)
	}
	if strings.HasPrefix(data.GetMessage(), "queried messages between sequence numbers") {
		return fmt.Sprintf("%s numMsgs=%v sourceChain=%v seqNumRange=%v",
			fields,
			data.RawLoggerFields["numMsgs"],
			data.RawLoggerFields["sourceChainSelector"],
			data.RawLoggerFields["seqNumRange"],
		)
	}
	if strings.HasPrefix(data.GetMessage(), "decoded messages between sequence numbers") {
		return fmt.Sprintf("%s sourceChain=%v seqNumRange=%v",
			fields, data.RawLoggerFields["sourceChainSelector"], data.RawLoggerFields["seqNumRange"])
	}

	return ""
}
