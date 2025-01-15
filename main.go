package main

import (
	"os"

	"github.com/lassejlv/tailwindcss-in-go/cmds"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// Enable pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	cmd := os.Args[1]

	if len(os.Args) < 2 {
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
