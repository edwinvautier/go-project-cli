package services

import (
	log "github.com/sirupsen/logrus"
	"os/user"
	"os"
)

func CreateConfigFile() *os.File {
	usr, err := user.Current()
    if err != nil {
        return nil
    }
	configPath := usr.HomeDir + "/.go-project-cli.yaml"
	// Create the file
	file, err := os.Create(configPath)
	if err != nil {
		log.Error(err)
		return nil
	}
	
	return file
}

func GetConfigFile() *os.File {
	usr, err := user.Current()
    if err != nil {
        return nil
    }
	configPath := usr.HomeDir + "/.go-project-cli.yaml"
	file, err := os.Open(configPath)
	
	return file
}