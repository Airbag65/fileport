package cmd

import (
	"fmt"

	"github.com/Airbag65/fileport/cli-client/fs"
	"github.com/Airbag65/fileport/cli-client/net"
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
	ip, err := fs.GetIP()
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	authStatus, err := net.AuthServiceIsUp()
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	if !authStatus {
		red.Println("Could not connect to the server")
		fmt.Printf("Using IP: %s\n", ip)
		return
	}

	auth, err := fs.GetLocalAuth()
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	if auth.AuthToken == "" {
		red.Println("You are not signed in to fileport!")
		red.Println("Run 'fileport login' to sign in")
		fmt.Printf("Using IP: %s\n", ip)
		return
	}
	code, err := net.ValidateUserToken(auth.Email, auth.AuthToken)
	if err != nil {
		red.Println("Something went wrong")
		return
	}
	if code != net.OK {
		red.Println("You are not signed in to fileport!")
		red.Println("Run 'fileport login' to sign in")
		fmt.Printf("Using IP: %s\n", ip)
		return
	}
	green.Println("You are signed in to fileport! fileport is ready to use")
	fmt.Println("Your credentials:")
	fmt.Println("-----------------")
	fmt.Printf("Name: %s %s\n", auth.Name, auth.Surname)
	fmt.Printf("Email: %s\n", auth.Email)
	fmt.Printf("Using IP: %s\n", ip)

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
