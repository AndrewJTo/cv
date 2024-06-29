package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	dataFile := flag.String("data", "./data.yml", "File containing CV data")
	templateDir := flag.String("template", "./tmpl", "File containing templates")

	flag.Parse()

	tmpl := template.Must(template.ParseGlob("./" + filepath.Clean(*templateDir) + "/*.html"))

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
