package cmds

import "fmt"

func ParseCommand(args []string) {
	fmt.Println("Hello from parse command!")
}

func init() {
	Commands = append(Commands, Command{
		Name:        "parse",
		Description: "Parse all files to file classe's",
		Handler:     ParseCommand,
	})
}
