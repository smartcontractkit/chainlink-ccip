package main

import (
	"os"
	"text/template"
)

const tmplFile = "values.tmpl"

const outputPath = "output/"

func main() {
	c := make(map[string]string)
	c["network"] = os.Args[1]
	c["chainId"] = os.Args[2]
	c["rpcEndpoint"] = os.Args[3]

	tmpl, err := template.New("values.tmpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(outputPath + os.Args[4])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, c)
	if err != nil {
		panic(err)
	}
}
