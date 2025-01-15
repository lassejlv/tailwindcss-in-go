package cmds

import (
	"github.com/lassejlv/tailwindcss-in-go/utils"
)

func ParseCommand(args []string) {
	utils.Parser(utils.GetConfig())
}

func init() {
	Commands = append(Commands, Command{
		Name:        "parse",
		Description: "Parse all files to file classe's",
		Handler:     ParseCommand,
	})
}
