package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

const tmplFile = "values.yaml.tmpl"

type chain struct {
	NetworkId     int64
	FinalityDepth int
}
type config struct {
	Chains []chain
}

func main() {
	// Depending on how many variations of config we need, we can hardcode it here,
	// or read from some config files
	c := config{
		Chains: []chain{
			{NetworkId: 1337, FinalityDepth: 50},
			{NetworkId: 2337, FinalityDepth: 50},
			{NetworkId: 90000001, FinalityDepth: 100},
			{NetworkId: 90000002, FinalityDepth: 100},
		},
	}

	scriptsDir := os.Getenv("SCRIPT_DIR")
	tmplPath := path.Join(scriptsDir, tmplFile)

	_, err := os.ReadFile(tmplPath)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		os.Exit(1)
	}

	tmpl, err := template.New("values.yaml.tmpl").ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, c)
	if err != nil {
		panic(err)
	}
}
