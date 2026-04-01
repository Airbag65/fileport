package cmd

import "github.com/fatih/color"

type Command interface {
	Execute()
}

type HelpCommand struct{}

type StatusCommand struct{}

var (
	red      = color.RGB(255, 0, 0)
	fpYellow = color.RGB(255, 249, 87)
)

