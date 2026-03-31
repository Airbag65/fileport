package main

import (
	"fmt"
	"os"
)

// GetUserDirPath looks for the path string for the user
// with the given email. If it does not exists, the directory
// will be created
func GetUserDirPath(email string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	userDir := fmt.Sprintf("%s/.fileport/users/%s", homedir, email)
	if _, err = os.Stat(userDir); os.IsNotExist(err) {
		if err = os.MkdirAll(userDir, 0755); err != nil {
			return "", err
		}
	}
	return userDir, nil
}

type INodeType int

const (
	Directory INodeType = iota
	File
)

type INode interface {
	Type() INodeType
	Name() string
	Print(string, bool)
}

type DirectoryINode struct {
	name  string
	Nodes []INode
}

func (d *DirectoryINode) Type() INodeType {
	return Directory
}

func (d *DirectoryINode) Name() string {
	return d.name
}

func (d *DirectoryINode) Print(indent string, lastINode bool) {
	fmt.Printf("%s+- %s (directory) \n", indent, d.Name())

	if lastINode {
		indent += "   "
	} else {
		indent += "|  "
	}

	for i, node := range d.Nodes {
		node.Print(indent, i == len(d.Nodes)-1)
	}
}

type FileINode struct {
	name    string
	Content []byte
}

func (f *FileINode) Type() INodeType {
	return File
}

func (f *FileINode) Name() string {
	return f.name
}

func (f *FileINode) Print(indent string, lastINode bool) {
	fmt.Printf("%s+- %s\n", indent, f.Name())
}

func NewDirectoryINode(name string, nodes []INode) *DirectoryINode {
	return &DirectoryINode{
		name:  name,
		Nodes: nodes,
	}
}

func NewFileINode(name string) *FileINode {
	return &FileINode{name: name}
}
