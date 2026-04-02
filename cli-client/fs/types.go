package fs

import (
	"fmt"
	"os"
	"strings"
)

type Inode interface {
	Type() InodeType
	Print(string, bool)
}

type InodeType int

const (
	Directory InodeType = iota
	File
)

type DirectoryInode struct {
	Name  string  `json:"dir_name"`
	Nodes []Inode `json:"dir_nodes"`
}

func (d *DirectoryInode) Type() InodeType {
	return Directory
}

func (d *DirectoryInode) Print(indent string, lastINode bool) {
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

type FileInode struct {
	Name string `json:"file_name"`
}

func (f *FileInode) Type() InodeType {
	return File
}

func (f *FileInode) Print(indent string, lastINode bool) {
	fmt.Printf("%s+- %s\n", indent, f.Name)
}

func NewDirectoryInode(name string, nodes []Inode) *DirectoryInode {
	return &DirectoryInode{
		Name:  name,
		Nodes: nodes,
	}
}

func NewFileINode(name string) *FileInode {
	return &FileInode{Name: name}
}

func (d *DirectoryInode) AddINode(node Inode) {
	d.Nodes = append(d.Nodes, node)
}

func GetDirectoryContent(email string) (*DirectoryInode, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/.fileport/users/%s", homeDir, email)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	rootContent := []Inode{}
	for _, entry := range entries {
		if entry.IsDir() {
			rootContent = append(rootContent, NewDirectoryInode(entry.Name(), []Inode{}))
		} else {
			rootContent = append(rootContent, NewFileINode(entry.Name()))
		}
	}
	if strings.Split(email, "/")[1] == "." && len(strings.Split(email, "/")) == 2 {
		email = strings.Split(email, "/")[0]
	}
	dir := NewDirectoryInode(email, rootContent)
	return dir, nil
}

func GetDirectoryContentR(email string) (*DirectoryInode, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/.fileport/users/%s", homeDir, email)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	rootContent := []Inode{}
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
	dir := NewDirectoryInode((map[bool]string{true: strings.Split(email, "/")[0], false: dirName})[dirName == "."], rootContent)
	return dir, nil
}
