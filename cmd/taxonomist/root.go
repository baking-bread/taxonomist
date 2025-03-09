package cmd

import (
	"fmt"
	"os"

	config "github.com/baking-bread/taxonomist/internal"
	"github.com/baking-bread/taxonomist/pkg/generator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	configPath string
	prefix     string
	suffix     string
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
			"format": format,
			"prefix": prefix,
			"suffix": suffix,
		}).Debug("Starting name generation")

		gen := generator.NewNameGenerator(cfg.Adjectives, cfg.Nouns)

		fmt.Print(gen.GenerateName(format, prefix, suffix))

		return nil
	},
}

func Execute() error {
	BaseCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to the configuration file (optional)")
	BaseCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "Prefix to add to generated names")
	BaseCmd.PersistentFlags().StringVarP(&suffix, "suffix", "s", "", "suffix to add to generated names")
	BaseCmd.PersistentFlags().StringVarP(&format, "format", "f", "kebab", "Output format (kebab, camel, snake, pascal, uper, cobol)")
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
