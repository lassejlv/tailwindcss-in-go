package cmds

type Command struct {
	Name        string
	Description string
	Handler     func(args []string)
}

var Commands []Command
