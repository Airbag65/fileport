package fs

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Global Global `toml:"global"`
}

type Global struct {
	IpAddr     string `toml:"ip_addr"`
	SourcePath string `toml:"source_path"`
}

func GetConfiguration() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := homeDir + "/.fileport/config.toml"
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	_, err = toml.Decode(string(configFile), &config)
	return &config, nil
}

func GetTitle() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := homeDir + "/.fileport/fileport_title.txt"

	content, err := os.ReadFile(path)
	return string(content), err
}
