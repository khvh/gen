package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v3"
)

type specEntityField struct {
	Ref      string `yaml:"ref"`
	Tag      string `yaml:"tag"`
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
}

type specService struct {
	ID  string            `yaml:"id"`
	In  []specEntityField `yaml:"in"`
	Out string            `yaml:"out"`
}

type specEntity struct {
	Ref      string            `yaml:"ref"`
	Fields   []specEntityField `yaml:"fields"`
	Sub      []specEntity      `yaml:"sub"`
	Services []specService     `yaml:"services"`
}

type specEntities struct {
	Entities []specEntity `yaml:"entities"`
}

var (
	in, out string
)

//go:embed main.tmpl
var mainTemplate embed.FS

func main() {
	flag.StringVar(&in, "in", "spec.yml", "Spec file")
	flag.StringVar(&out, "out", "internal/generated", "Output dir")

	log.Default().Print(in, " <-> ", out)

	bts, err := os.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	ents := specEntities{}

	yaml.Unmarshal(bts, &ents)

	err = os.RemoveAll("internal/generated/entities")
	if err != nil {
		log.Default().Print(err)
	}

	err = os.MkdirAll("internal/generated/entities", 0755)
	if err != nil {
		log.Default().Print(err)
	}

	// tmpl := template.Must(template.New("main").Funcs(sprig.FuncMap()).ParseFiles("main.tmpl"))

	for _, item := range ents.Entities {
		tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(mainTemplate, "main.tmpl")
		if err != nil {
			log.Fatal(err)
		}

		var processed bytes.Buffer

		err = tmpl.ExecuteTemplate(&processed, "main.tmpl", &item)
		if err != nil {
			log.Fatalf("Unable to parse data into template: %v\n", err)
		}

		formatted, err := format.Source(processed.Bytes())
		if err != nil {
			log.Fatalf("Could not format processed template: %v\n", err)
		}

		path := fmt.Sprintf("internal/generated/entities/%s.go", strings.ToLower(item.Ref))

		log.Default().Println(path)

		os.WriteFile(path, formatted, 0755)
	}

	// e := ent.New(ent.NewService(ent.WithRepository(generated.NewEntityMemoryRepository())))

	// fmt.Println(e.TestMe())
}