package cmd

import (
	"os"
	"path/filepath"

	"github.com/loksonarius/gli/cfg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string
	Config     cfg.Config

	rootCmd = &cobra.Command{
		Use:   "gli",
		Short: "Handle daily dev tasks with GitLab",
		Long: `gli is a CLI to interact with GitLab. The intent is to be more
than a way of running arbitrary CRUD commands against GitLab API
resources, and act like an actual developer interface for typical
developer workflows with GitLab.`,
		Run: func(cmd *cobra.Command, args []string) {
			logger.Println(Config)
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(
		&Config.CurrentTarget,
		"target",
		"T",
		Config.CurrentTarget,
		"target GitLab instance to run commands against",
	)
	viper.BindPFlag("currentTarget", rootCmd.PersistentFlags().Lookup("target"))
}

func initConfig() {
	configDir, err := os.UserConfigDir()

	if err != nil {
		logger.Println("Unable to retrieve user's config dir")
	} else {
		configPath = filepath.Join(configDir, "gli")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath(configPath)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// config doesn't already exist, create it
				if err = os.MkdirAll(configPath, 0755); err != nil {
					logger.Fatalf(
						"Got error trying to create config dir: %v",
						err,
					)
				}
				err = viper.SafeWriteConfig()
				if err != nil {
					logger.Fatalf(
						"Got error trying to initialize config file: %v",
						err,
					)
				}
			} else {
				// config exists, but we failed to read it
				logger.Fatalf(
					"Got error trying to read config file: %v",
					err,
				)
			}
		}

		if err = viper.Unmarshal(&Config); err != nil {
			logger.Fatalf(
				"Got error unmarshalling config into expected structure: %v",
				err,
			)
		}
	}
}

// Execute kicks off the gli CLI
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("{{.Version}}\n")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
