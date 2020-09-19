package services

import (
	"os"
	log "github.com/sirupsen/logrus"
	"github.com/AlecAivazis/survey"
)

func CreateStructure(path string, modules []string)  {
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
		}
	}
}

func removeAll(path string) {
	os.RemoveAll(path)
	os.Mkdir(path, os.ModePerm)
}