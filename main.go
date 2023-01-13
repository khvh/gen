package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go/format"
	"html/template"
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

//go:embed entity_gen.tmpl
var mainTemplate embed.FS

//go:embed entity_test_gen.tmpl
var mainTestTemplate embed.FS

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	flag.StringVar(&in, "in", "spec.yml", "Spec file")
	flag.StringVar(&out, "out", "internal/generated", "Output dir")

	log.Info().Msgf(in, " <-> ", out)

	bts, err := os.ReadFile(in)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	ents := specEntities{}

	if err := yaml.Unmarshal(bts, &ents); err != nil {
		log.Fatal().Err(err).Send()
	}

	err = os.RemoveAll("internal/generated/entities")
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = os.MkdirAll("internal/generated/entities", 0755)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for _, item := range ents.Entities {
		genEntity(&item)
		genEntityTests(&item)
	}
}

func genEntity(item *specEntity) {
	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(mainTemplate, "entity_gen.tmpl")

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	var processed bytes.Buffer

	err = tmpl.ExecuteTemplate(&processed, "entity_gen.tmpl", item)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	path := fmt.Sprintf("internal/generated/entities/%s.go", strings.ToLower(item.Ref))

	log.Info().Msgf("Writing %s -> %s", item.Ref, path)

	if err := os.WriteFile(path, formatted, 0755); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func genEntityTests(item *specEntity) {
	tmpl, err := template.New("").Funcs(sprig.FuncMap()).ParseFS(mainTestTemplate, "entity_test_gen.tmpl")

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	var processed bytes.Buffer

	err = tmpl.ExecuteTemplate(&processed, "entity_test_gen.tmpl", item)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	path := fmt.Sprintf("internal/generated/entities/%s_test.go", strings.ToLower(item.Ref))

	log.Info().Msgf("Writing %s -> %s", item.Ref, path)

	if err := os.WriteFile(path, formatted, 0755); err != nil {
		log.Fatal().Err(err).Send()
	}
}