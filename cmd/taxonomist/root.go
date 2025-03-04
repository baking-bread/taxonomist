package cmd

import (
	"fmt"
	"os"
	"strings"

	config "github.com/baking-bread/taxonomist/internal"
	"github.com/sirupsen/logrus"
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
	debug      bool
	log        *logrus.Logger
)

func initLogger() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
}

var BaseCmd = &cobra.Command{
	Use:   "taxonomist",
	Short: "A Name-Generator CLI Tool",
	Long:  "Taxonomy: A simple name generator tool that can name whatever you can think of",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg *config.Config
		var err error

		if configPath != "" {
			log.WithField("config_path", configPath).Debug("Loading configuration file")
			cfg, err = config.LoadConfig(configPath)
			if err != nil {
				log.WithError(err).WithField("config_path", configPath).Warn("Failed to load config file")
				log.Info("Using default configuration")
				cfg = config.GetDefaultConfig()
			}
		} else {
			log.Debug("No config file specified, using default configuration")
			cfg = config.GetDefaultConfig()
		}

		if len(cfg.Adjectives) == 0 || len(cfg.Nouns) == 0 {
			log.Error("Invalid configuration: must have at least one adjective and one noun")
			return fmt.Errorf("invalid configuration: must have at least one adjective and one noun")
		}

		log.WithFields(logrus.Fields{
			"count":     count,
			"adj_count": adjCount,
			"format":    format,
			"prefix":    prefix,
			"suffix":    sufix,
			"separator": separator,
		}).Debug("Starting name generation")

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

			log.WithFields(logrus.Fields{
				"iteration": i + 1,
				"name":      name,
			}).Debug("Generated name")

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
	BaseCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	if envConfig := os.Getenv("CONFIG_FILE"); envConfig != "" {
		configPath = envConfig
		log.WithField("config_path", envConfig).Debug("Using config file from environment")
	}

	if err := BaseCmd.Execute(); err != nil {
		log.WithError(err).Error("Command execution failed")
		return err
	}
	return nil
}
