package main

import (
	"fmt"
	"os"
	"strings"
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

func DirectoryExists(path string) bool {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	targetDir := fmt.Sprintf("%s/.fileport/users/%s", homeDir, path)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return false
	}
	return true
}

type INodeType int

const (
	Directory INodeType = iota
	File
)

type INode interface {
	Type() INodeType
	Print(string, bool)
}

type DirectoryINode struct {
	Name  string  `json:"dir_name"`
	Nodes []INode `json:"dir_nodes"`
}

func (d *DirectoryINode) Type() INodeType {
	return Directory
}

func (d *DirectoryINode) Print(indent string, lastINode bool) {
	fmt.Printf("%s+- %s (directory) \n", indent, d.Name)

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
	Name string `json:"file_name"`
}

func (f *FileINode) Type() INodeType {
	return File
}

func (f *FileINode) Print(indent string, lastINode bool) {
	fmt.Printf("%s+- %s\n", indent, f.Name)
}

func NewDirectoryINode(name string, nodes []INode) *DirectoryINode {
	return &DirectoryINode{
		Name:  name,
		Nodes: nodes,
	}
}

func NewFileINode(name string) *FileINode {
	return &FileINode{Name: name}
}

func (d *DirectoryINode) AddINode(node INode) {
	d.Nodes = append(d.Nodes, node)
}

func GetDirectoryContent(email string) (*DirectoryINode, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	if !DirectoryExists(email) {
		return nil, nil
	}
	path := fmt.Sprintf("%s/.fileport/users/%s", homeDir, email)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	rootContent := []INode{}
	for _, entry := range entries {
		if entry.IsDir() {
			rootContent = append(rootContent, NewDirectoryINode(entry.Name(), []INode{}))
		} else {
			rootContent = append(rootContent, NewFileINode(entry.Name()))
		}
	}
	if strings.Split(email, "/")[1] == "." && len(strings.Split(email, "/")) == 2 {
		email = strings.Split(email, "/")[0]
	}
	dir := NewDirectoryINode(email, rootContent)
	return dir, nil
}

func GetDirectoryContentR(email string) (*DirectoryINode, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	if !DirectoryExists(email) {
		return nil, nil
	}
	path := fmt.Sprintf("%s/.fileport/users/%s", homeDir, email)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	rootContent := []INode{}
	for _, entry := range entries {
		if entry.IsDir() {

			newDir, err := GetDirectoryContentR(fmt.Sprintf("%s/%s", email, entry.Name()))
			if err != nil {
				return nil, err
			}
			rootContent = append(rootContent, newDir)
		} else {
			rootContent = append(rootContent, NewFileINode(entry.Name()))
		}
	}
	dirName := strings.Split(email, "/")[len(strings.Split(email, "/"))-1]
	dir := NewDirectoryINode((map[bool]string{true: strings.Split(email, "/")[0], false: dirName})[dirName == "."], rootContent)
	return dir, nil
}
