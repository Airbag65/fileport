package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

type Command interface {
	Execute()
}

type HelpCommand struct{}

type StatusCommand struct{}

type LoginCommad struct{}

type SignOutCommand struct{}

type RegisterCommand struct{}

type ListCommand struct{}

var (
	red      = color.RGB(255, 0, 0)
	green    = color.RGB(0, 255, 0)
	fpYellow = color.RGB(255, 249, 87)
	yellow   = color.RGB(255, 255, 0)
)

func GetCommand(args []string) Command {
	if len(args) < 1 {
		fmt.Println("Usage: fileport <command>")
		yellow.Println("Run 'fileport help' for further instructions")
		return nil
	}
	switch args[0] {
	case "help":
		return &HelpCommand{}
	case "status":
		return &StatusCommand{}
	case "login":
		return &LoginCommad{}
	case "signout":
		return &SignOutCommand{}
	case "register":
		return &RegisterCommand{}
	case "list":
		return &ListCommand{}
	default:
		fmt.Println("fileport: Invalid argument")
		return nil
	}
}
