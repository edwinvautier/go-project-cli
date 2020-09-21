package services

import (
	log "github.com/sirupsen/logrus"
	"os/user"
	"os"
)

// CreateConfigFile is the function that creates the config file in the user filesystem if it doesn't exist.
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