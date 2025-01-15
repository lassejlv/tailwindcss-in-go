package utils

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog/log"
)

// This will parse the config from, css.json
var ConfigFileName string = "css.json"

type Config struct {
	Extensions []string "json:\"extensions\""
	Minify     bool     "json:\"minify\""
}

func GetConfig() Config {

	workingDir, err := os.Getwd()
	if err != nil {
		log.Error().Msgf("Error getting working directory: %s", err)
	}

	path := workingDir + "/" + ConfigFileName

	fileExists, err := os.Stat(path)
	if err != nil {
		log.Error().Msgf("Error checking if file '%s' exists: %s", ConfigFileName, err)
	}

	if fileExists.IsDir() {
		log.Error().Msgf("config file '%s' is a directory", ConfigFileName)
	}

	configFile, err := os.Open(ConfigFileName)
	if err != nil {
		log.Error().Msgf("Error opening file '%s': %s", ConfigFileName, err)
	}

	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Error().Msgf("Error decoding file '%s': %s", ConfigFileName, err)
	}

	return config
}
