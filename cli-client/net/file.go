package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Airbag65/fileport/cli-client/fs"
)

func GetFilesList(path string) (fs.Inode, error) {
	ip, err := fs.GetCofigIP()
	if err != nil {
		return nil, err
	}
	requset, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8001/files/list", ip), &bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	auth, err := fs.GetLocalAuth()
	if err != nil {
		return nil, err
	}
	requset.Header.Add("Authorization", fmt.Sprintf("Bearer %s", auth.AuthToken))
	requset.Header.Set("target", path)
	response, err := client.Do(requset)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, nil
	}
	var dir fs.Inode
	// TODO: Cannot unmarshal this structure for some reason. Fix this in another way
	err = json.NewDecoder(response.Body).Decode(&dir)
	if err != nil {
		return nil, err
	}
	return dir, nil
}
