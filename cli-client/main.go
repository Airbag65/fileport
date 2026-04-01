package main

import (
	"os"

	"github.com/Airbag65/fileport/cli-client/cmd"
)

func main() {
	command := cmd.GetCommand(os.Args[1:])
	if command == nil {
		return
	}
	command.Execute()
}
