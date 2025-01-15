package main

import (
	"os"

	"github.com/lassejlv/tailwindcss-in-go/cmds"
	"github.com/lassejlv/tailwindcss-in-go/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Enable pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	config := utils.GetConfig()

	if len(config.Extensions) == 0 {
		log.Error().Msg("No extensions were found in the config file")
		return
	}

	cmd := os.Args[1]

	if len(os.Args) < 1 {
		log.Warn().Msg("No command was provided")
		return
	}

	for _, command := range cmds.Commands {
		if command.Name == cmd {
			command.Handler(os.Args[2:])
			return
		}
	}

	log.Error().Msgf("Command '%s' was not found", cmd)
}
