package cmds

import (
	"fmt"

	"github.com/lassejlv/tailwindcss-in-go/utils"
)

func VersionCommand(args []string) {
	fmt.Println(utils.CurrentVersion)
}

func init() {
	Commands = append(Commands, Command{
		Name:        "version",
		Description: "Print the version of the program",
		Handler:     VersionCommand,
	})
}
