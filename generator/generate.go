package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	dataFile := flag.String("data", "./data.yml", "File containing CV data")

	flag.Parse()

	tmpl := template.Must(template.ParseGlob("./tmpl/*.html"))

	data := map[string]interface{}{}
	buf, err := os.ReadFile(*dataFile)
	if err != nil {
		log.Fatalf("Could not read file: %s", err)
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		log.Fatalf("Could not unmarshal YAML: %s", err)
	}

	f, err := os.Create("out.html")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = tmpl.ExecuteTemplate(f, "main.html", data)

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}
