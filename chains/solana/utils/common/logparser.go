package common

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	bin "github.com/gagliardetto/binary"
)

type AnchorInstruction struct {
	Name         string
	ProgramID    string
	Logs         []string
	ComputeUnits int
	InnerCalls   []*AnchorInstruction
	EventData    []*EventData
}

type EventMapping struct {
	Name string
	New  func() interface{}
}

type EventData struct {
	Base64Data  string
	DecodedData []byte
	EventName   string
	Data        interface{} // Decoded data using the provided type
}

func EventMappingFor[T any](name string) EventMapping {
	return EventMapping{
		Name: name,
		New: func() interface{} {
			return new(T)
		},
	}
}

func NormalizeData(d []byte) []byte {
	if d == nil {
		return []byte{}
	}
	return d
}

// holds the state while parsing log lines
type parser struct {
	instructions       []*AnchorInstruction
	stack              []*AnchorInstruction
	currentInstruction *AnchorInstruction
	expectedEvents     []EventMapping

	// compiled regexes for performance
	programInvokeRegex  *regexp.Regexp
	programSuccessRegex *regexp.Regexp
	computeUnitsRegex   *regexp.Regexp
}

// creates a parser instance with precompiled regex
func newParser(expectedEvents []EventMapping) *parser {
	return &parser{
		instructions:        []*AnchorInstruction{},
		stack:               []*AnchorInstruction{},
		expectedEvents:      expectedEvents,
		programInvokeRegex:  regexp.MustCompile(`Program (\w+) invoke`),
		programSuccessRegex: regexp.MustCompile(`Program (\w+) success`),
		computeUnitsRegex:   regexp.MustCompile(`Program (\w+) consumed (\d+) of \d+ compute units`),
	}
}

func (p *parser) handleProgramInvokeLine(line string) bool {
	match := p.programInvokeRegex.FindStringSubmatch(line)
	if len(match) <= 1 {
		return false
	}

	newInstruction := &AnchorInstruction{
		ProgramID:    match[1],
		Name:         "",
		Logs:         []string{},
		ComputeUnits: 0,
		InnerCalls:   []*AnchorInstruction{},
		EventData:    []*EventData{},
	}

	if len(p.stack) == 0 {
		p.instructions = append(p.instructions, newInstruction)
	} else {
		p.stack[len(p.stack)-1].InnerCalls = append(p.stack[len(p.stack)-1].InnerCalls, newInstruction)
	}

	p.stack = append(p.stack, newInstruction)
	p.currentInstruction = newInstruction
	return true
}

// check if line is a "Program X success" line and updates stack
func (p *parser) handleProgramSuccessLine(line string) bool {
	match := p.programSuccessRegex.FindStringSubmatch(line)
	if len(match) <= 1 {
		return false
	}

	if len(p.stack) > 0 {
		p.stack = p.stack[:len(p.stack)-1] // pop
		if len(p.stack) > 0 {
			p.currentInstruction = p.stack[len(p.stack)-1]
		} else {
			p.currentInstruction = nil
		}
	}
	return true
}

func (p *parser) handleInstructionNameLine(line string) bool {
	if !strings.Contains(line, "Instruction:") {
		return false
	}
	if p.currentInstruction != nil {
		p.currentInstruction.Name = strings.TrimSpace(strings.Split(line, "Instruction:")[1])
	}
	return true
}

func (p *parser) handleProgramDataLine(line string) bool {
	if !strings.Contains(line, "Program data:") {
		return false
	}
	if p.currentInstruction == nil {
		return true // line recognized but no current instruction, do nothing
	}

	base64Data := strings.TrimSpace(strings.TrimPrefix(line, "Program data:"))
	decodedData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return true // recognized line but decode failed, no event
	}

	for _, event := range p.expectedEvents {
		if IsEvent(event.Name, decodedData) {
			obj := event.New()
			// NOTE: skipping the first 8 bytes which are the discriminator
			if err := bin.UnmarshalBorsh(&obj, decodedData[8:]); err != nil {
				continue
			}

			eventData := &EventData{
				Base64Data:  base64Data,
				DecodedData: decodedData,
				EventName:   event.Name,
				Data:        obj,
			}
			p.currentInstruction.EventData = append(p.currentInstruction.EventData, eventData)
		}
	}
	return true
}

func (p *parser) handleProgramLogLine(line string) bool {
	if !strings.HasPrefix(line, "Program log:") {
		return false
	}
	if p.currentInstruction != nil {
		logMessage := strings.TrimSpace(strings.TrimPrefix(line, "Program log:"))
		p.currentInstruction.Logs = append(p.currentInstruction.Logs, logMessage)
	}
	return true
}

func (p *parser) handleComputeUnitsLine(line string) bool {
	match := p.computeUnitsRegex.FindStringSubmatch(line)
	if len(match) <= 1 {
		return false
	}
	programID := match[1]
	units, _ := strconv.Atoi(match[2])

	// Find the instruction in the stack that matches this program ID
	for i := len(p.stack) - 1; i >= 0; i-- {
		if p.stack[i].ProgramID == programID {
			p.stack[i].ComputeUnits = units
			break
		}
	}
	return true
}

func ParseLogMessages(logMessages []string, expectedEvents []EventMapping) []*AnchorInstruction {
	p := newParser(expectedEvents)

	for _, line := range logMessages {
		line = strings.TrimSpace(line)

		// Try each handler in turn
		if p.handleProgramInvokeLine(line) {
			continue
		}
		if p.handleProgramSuccessLine(line) {
			continue
		}
		if p.handleInstructionNameLine(line) {
			continue
		}
		if p.handleProgramDataLine(line) {
			continue
		}
		if p.handleProgramLogLine(line) {
			continue
		}
		if p.handleComputeUnitsLine(line) {
			continue
		}

		// if none matched, this line might be irrelevant
	}

	return p.instructions
}

// Pretty prints the given Anchor instructions.
// Example usage:
// parsed := utils.ParseLogMessages(result.Meta.LogMessages)
// output := utils.LogTxResult(parsed)
// t.Logf("Tx logs: %s", output)
func LogTxResult(instructions []*AnchorInstruction) string {
	var output strings.Builder

	var printInstruction func(*AnchorInstruction, int, string)
	printInstruction = func(instruction *AnchorInstruction, index int, indent string) {
		output.WriteString(fmt.Sprintf("%sInstruction %d: %s\n", indent, index, instruction.Name))
		output.WriteString(fmt.Sprintf("%s  Program ID: %s\n", indent, instruction.ProgramID))
		output.WriteString(fmt.Sprintf("%s  Compute Units: %d\n", indent, instruction.ComputeUnits))

		// Print Events
		if len(instruction.EventData) > 0 {
			output.WriteString(fmt.Sprintf("%s  Events:\n", indent))
			for _, event := range instruction.EventData {
				output.WriteString(fmt.Sprintf("%s    Event: %s:\n", indent, event.EventName))
				output.WriteString(fmt.Sprintf("%s      Base64Data: %+v\n", indent, event.Base64Data))
				output.WriteString(fmt.Sprintf("%s      DecodedData: %+v\n", indent, event.DecodedData))
				output.WriteString(fmt.Sprintf("%s      Data: %+v\n", indent, event.Data))
			}
		}

		output.WriteString(fmt.Sprintf("%s  Logs:\n", indent))
		for _, log := range instruction.Logs {
			output.WriteString(fmt.Sprintf("%s    %s\n", indent, log))
		}

		if len(instruction.InnerCalls) > 0 {
			output.WriteString(fmt.Sprintf("%s  Inner Calls:\n", indent))
			for i, innerCall := range instruction.InnerCalls {
				printInstruction(innerCall, i+1, indent+"    ")
			}
		}
	}

	for i, instruction := range instructions {
		printInstruction(instruction, i+1, "")
	}

	return output.String()
}
