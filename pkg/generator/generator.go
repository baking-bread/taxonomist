package generator

import (
	"strings"

	"github.com/baking-bread/taxonomist/internal"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type NameGenerator struct {
	Adjectives []string
	Nouns      []string
}

func NewNameGenerator(adjectives []string, nouns []string) *NameGenerator {
	return &NameGenerator{
		Adjectives: adjectives,
		Nouns:      nouns,
	}
}

func (g *NameGenerator) GenerateName(format string, prefix string, suffix string) string {
	var name string
	var parts []string
	var adjective = g.Adjectives[internal.Random(len(g.Adjectives))]
	var noun = g.Nouns[internal.Random(len(g.Nouns))]

	if prefix != "" {
		parts = append(parts, prefix)
	}

	parts = append(parts, adjective, noun)

	if suffix != "" {
		parts = append(parts, suffix)
	}

	// camelCase
	// PascalCase
	// snake_case
	// kebab-case
	// UPER_CASE
	// COBOL-CASE

	for i := range parts {
		if format == "camel" && i < 1 {
			continue
		}

		if format == "camel" || format == "pascal" {
			parts[i] = cases.Title(language.English).String(parts[i])
		} else if format == "uper" || format == "cobol" {
			parts[i] = cases.Upper(language.English).String(parts[i])
		} else {
			parts[i] = cases.Lower(language.English).String(parts[i])
		}
	}

	if format == "camel" || format == "pascal" {
		name = strings.Join(parts, "")
	} else if format == "snake" || format == "uper" {
		name = strings.Join(parts, "_")
	} else {
		name = strings.Join(parts, "-")
	}

	return name
}
