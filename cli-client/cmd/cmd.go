package cmd

import (
	"fmt"

	"github.com/Airbag65/fileport/cli-client/fs"
)

func (c *HelpCommand) Execute() {
	title, err := fs.GetTitle()
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	fpYellow.Println(title)
}

func (c *StatusCommand) Execute() {
	config, err := fs.GetConfiguration()
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	fmt.Println(config)
}

func GetCommand(args []string) Command {
	switch args[0] {
	case "help":
		return &HelpCommand{}
	case "status":
		return &StatusCommand{}
	}
	fmt.Println(args)
	return nil
}
