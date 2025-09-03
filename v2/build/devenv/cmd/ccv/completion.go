package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func getCommands() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "up", Description: "Spin up the development environment"},
		{Text: "down", Description: "Tear down the development environment"},
		{Text: "reconfigure", Description: "Reconfigure development environment"},
		{Text: "bs", Description: "Manage the Blockscout EVM block explorer"},
		{Text: "obs", Description: "Manage the observability stack"},
		{Text: "db", Description: "Inspect Databases"},
		{Text: "exit", Description: "Exit the interactive shell"},
	}
}

func getSubCommands(parent string) []prompt.Suggest {
	switch parent {
	case "addresses":
		return []prompt.Suggest{
			{Text: "addresses", Description: "Print all smart-contract addresses data (CLDF)"},
		}
	case "bs":
		return []prompt.Suggest{
			{Text: "up", Description: "Spin up Blockscout"},
			{Text: "down", Description: "Spin down Blockscout"},
			{Text: "restart", Description: "Restart Blockscout"},
		}
	case "obs":
		return []prompt.Suggest{
			{Text: "up", Description: "Spin up observability stack"},
			{Text: "down", Description: "Spin down observability stack"},
			{Text: "restart", Description: "Restart observability stack"},
		}
	case "db":
		return []prompt.Suggest{
			{Text: "indexer", Description: "Inspect Indexer Database"},
			{Text: "aggregator", Description: "Inspect Aggregator Database"},
			{Text: "verifier", Description: "Inspect Verifier Database"},
			{Text: "executor", Description: "Inspect Executor Database"},
		}
	case "u":
		fallthrough
	case "up":
		fallthrough
	case "r":
		fallthrough
	case "reconfigure":
		return []prompt.Suggest{
			{Text: "Config", Description: `Select the configuration file, ex.: env.toml,override.toml`},
		}
	default:
		return []prompt.Suggest{}
	}
}

func executor(in string) {
	checkDockerIsRunning()
	in = strings.TrimSpace(in)
	if in == "" {
		return
	}
	if in == "exit" {
		fmt.Println("Goodbye!")
		os.Exit(0)
	}
	args := strings.Fields(in)
	os.Args = append([]string{"ccv"}, args...)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// completer provides autocomplete suggestions for multi-word commands
func completer(in prompt.Document) []prompt.Suggest {
	text := in.TextBeforeCursor()
	words := strings.Fields(text)
	lastCharIsSpace := len(text) > 0 && text[len(text)-1] == ' '

	switch {
	case len(words) == 0:
		return getCommands()
	case len(words) == 1:
		if lastCharIsSpace {
			return getSubCommands(words[0])
		} else {
			return prompt.FilterHasPrefix(getCommands(), words[0], true)
		}

	case len(words) >= 2:
		if lastCharIsSpace {
			return []prompt.Suggest{}
		} else {
			parent := words[0]
			currentWord := words[len(words)-1]
			return prompt.FilterHasPrefix(getSubCommands(parent), currentWord, true)
		}
	default:
		return []prompt.Suggest{}
	}
}

func StartShell() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("ccv> "),
		prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionTitle("CCIP Interactive Shell"),
		prompt.OptionMaxSuggestion(15),
		prompt.OptionShowCompletionAtStart(),
		prompt.OptionCompletionWordSeparator(" "),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.Black),
		prompt.OptionSuggestionTextColor(prompt.Green),
		prompt.OptionScrollbarThumbColor(prompt.DarkGray),
		prompt.OptionScrollbarBGColor(prompt.Black),
	)
	p.Run()
}
