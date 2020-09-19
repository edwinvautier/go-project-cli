package services

import (
	"html/template"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	hasRouter bool
	hasDB     bool
	hasDocker bool
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
		hasDB:     false,
		hasRouter: false,
		hasDocker: false,
	}
	for _, module := range modules {
		switch module {
		case "Router":
			config.hasRouter = true
		case "Docker":
			config.hasDocker = true
		case "Database":
			config.hasDB = true
		}
	}

	generateTemplates(config)
}

func generateTemplates(config Config) {
	box := packr.New("My Box", "../templates")
	main, err := box.FindString("main.txt")
	if err != nil {
		log.Error(err)
		return
	}

	mainTmp := template.Must(template.New("main").Parse(main))
	mainTmp.Execute(os.Stdout, config)
}

func removeAll(path string) {
	os.RemoveAll(path)
	os.Mkdir(path, os.ModePerm)
}
