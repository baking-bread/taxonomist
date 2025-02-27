package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/baking-bread/taxonomist/internal/config"
	"github.com/spf13/cobra"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	configPath string
	count      int
	prefix     string
	sufix      string
	separator  string
	adjCount   int
	format     string
)

var BaseCmd = &cobra.Command{
	Use:   "taxonomist",
	Short: "A Name-Generator CLI Tool",
	Long:  "Taxonomy: A simple name generator tool that can name whatever you can think of",
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg *config.Config
		var err error

		if configPath != "" {
			cfg, err = config.LoadConfig(configPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Failed to load config from %s: %v\n", configPath, err)
				fmt.Fprintf(os.Stderr, "Using default configuration...\n")
				cfg = config.GetDefaultConfig()
			}
		} else {
			cfg = config.GetDefaultConfig()
		}

		// Validate config has required data
		if len(cfg.Adjectives) == 0 || len(cfg.Nouns) == 0 {
			return fmt.Errorf("invalid configuration: must have at least one adjective and one noun")
		}

		for i := 0; i < count; i++ {
			adjectives := make([]string, adjCount)
			for j := 0; j < adjCount; j++ {
				adjectives[j] = cfg.GetRandomAdjective()
			}

			var name string
			switch format {
			case "kebab":
				name = strings.Join(append(adjectives, cfg.GetRandomNoun()), separator)
			case "camel":
				parts := append(adjectives, cfg.GetRandomNoun())
				for i := range parts {
					if i > 0 {
						parts[i] = cases.Title(language.English).String(parts[i])
					}
				}
				name = strings.Join(parts, "")
			case "snake":
				name = strings.Join(append(adjectives, cfg.GetRandomNoun()), "_")
			}

			if prefix != "" {
				name = prefix + separator + name
			}
			if sufix != "" {
				name = name + separator + sufix
			}
			fmt.Println(name)
		}
		return nil
	},
}

func Execute() error {
	BaseCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file (optional)")
	BaseCmd.PersistentFlags().IntVarP(&count, "count", "n", 1, "Number of names to generate")
	BaseCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Prefix to add to generated names")
	BaseCmd.PersistentFlags().StringVarP(&sufix, "sufix", "s", "", "sufix to add to generated names")
	BaseCmd.PersistentFlags().StringVarP(&separator, "separator", "e", "-", "Separator to use between prefix, generated name, and sufix")
	BaseCmd.PersistentFlags().IntVarP(&adjCount, "adjectives", "a", 1, "Number of adjectives to use in the name")
	BaseCmd.PersistentFlags().StringVarP(&format, "format", "f", "kebab", "Output format (kebab, camel, snake)")

	if envConfig := os.Getenv("CONFIG_FILE"); envConfig != "" {
		configPath = envConfig
	}

	return BaseCmd.Execute()
}
