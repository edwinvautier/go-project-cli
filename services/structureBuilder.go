package services

import (
	"os/exec"
	"os"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
)

// Config is the structure passed to the templates executors in order for them to access informations such as the username, its choices of modules etc..
type Config struct {
	HasRouter 	bool
	HasDB     	bool
	HasDocker 	bool
	Path 		string
	Box			*packr.Box
	Username 	string
	AppName 	string
}

// CreateStructure is the function used to create the configuration for the current command, create the repository and execute the templates inside.
func CreateStructure(path string, modules []string, username string, appName string) {
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
		Username: 	username,
		AppName: 	appName,
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

	initGitRepository(path)
	generateTemplates(config)
}

func generateTemplates(config Config) {
	// Generate packr Box
	config.Box = packr.New("My Box", "../templates")

	executeTemplate("", "main.txt", "main.go", config)
	executeTemplate("", "go.mod.txt", "go.mod", config)
	executeTemplate("", "readme.txt", "README.md", config)
	executeTemplate("", "gitignore.txt", ".gitignore", config)
	executeTemplate("services/", "lib.txt", "lib.go", config)
	
	if config.HasDB {
		executeTemplate("models/", "PlayerStruct.txt", "PlayerStruct.go", config)
		executeTemplate("models/", "DBConnector.txt", "DBConnector.go", config)
		executeTemplate("models/", "migrations.txt", "Migrations.go", config)
	}

	if config.HasRouter {
		executeTemplate("routes/", "Router.txt", "Router.go", config)
		executeTemplate("controllers/", "PlayersController.txt", "PlayersController.go", config)
		executeTemplate("controllers/", "RootController.txt", "RootController.go", config)
	}

	if config.HasDocker {
		executeTemplate("", "dockerfile.txt", "Dockerfile", config)
		executeTemplate("", "docker-compose.txt", "docker-compose.yml", config)
	}
}

func initGitRepository(path string) {
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
