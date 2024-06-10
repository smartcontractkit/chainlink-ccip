package main

import (
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

const tmplFile = "values.tmpl"
const configPath = "config/"
const outputPath = "output/"

func main() {
	var c map[string]interface{}
	yamlFile, err := os.ReadFile(configPath + os.Args[1])
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	tmpl, err := template.New("values.tmpl").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(outputPath + os.Args[2])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, c)
	if err != nil {
		panic(err)
	}
}
