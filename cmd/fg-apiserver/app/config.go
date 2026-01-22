package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultsHomeDir = ".cloudtest"
	defaultsConfigName = "fg-apiserver.yaml"
)

func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		for _, dir := range searchDir() {
			viper.AddConfigPath(dir)
		}

		viper.SetConfigFile("yaml")
		viper.SetConfigName(defaultsConfigName)
	}

	setupEnvironmentVariables()

	_ = viper.ReadInConfig()
}

func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CLOUDTEST")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func searchDir() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultsHomeDir), "."}
}

func filePath() string{
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultsConfigName, defaultsHomeDir)
}