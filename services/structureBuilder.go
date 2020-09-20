package services

import (
	"os/exec"
	"os"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	HasRouter 	bool
	HasDB     	bool
	HasDocker 	bool
	Path 		string
	Box			*packr.Box
}

func CreateStructure(path string, modules []string) {
	// Create the folder for the new project
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	} else {
		log.Warn("A directory with this name already exists.")

		wantsOverride := false
		prompt := &survey.Confirm{
			Message: "Override the folder ?",
			Default: false,
		}
		survey.AskOne(prompt, &wantsOverride)

		if wantsOverride {
			removeAll(path)
		} else {
			return
		}
	}

	// Generate the config object
	config := Config{
		HasDB:     	false,
		HasRouter: 	false,
		HasDocker: 	false,
		Path: 		path,
	}
	for _, module := range modules {
		switch module {
		case "Router":
			config.HasRouter = true
		case "Docker":
			config.HasDocker = true
		case "Database":
			config.HasDB = true
		}
	}

	initRepository(path)
	generateTemplates(config)
}

func generateTemplates(config Config) {
	// Generate packr Box
	config.Box = packr.New("My Box", "../templates")

	executeTemplate("", "main.txt", "main.go", config)
}

func initRepository(path string) {
	// Init git repository
	err := exec.Command("git", "init", path).Run()
	if err != nil {
		log.Error(err)
	}	
}

func removeAll(path string) {
	os.RemoveAll(path)
	os.Mkdir(path, os.ModePerm)
}
