package services

import (
	"html/template"
	"os"
	log "github.com/sirupsen/logrus"
)

func executeTemplate(path string, name string, outputName string, config Config) {
	// Get template content as string
	templateString, err := config.Box.FindString(path + name)
	if err != nil {
		log.Error(err)
		return
	}

	// Create the directory if not exist
	if _, err := os.Stat(config.Path + "/" + path); os.IsNotExist(err) {
		os.Mkdir(config.Path + "/" + path, os.ModePerm)
	}

	// Create the file
	file, err := os.Create(config.Path + "/" + path + outputName)
	if err != nil {
		log.Error(err)
		return
	}

	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	err = parsedTemplate.Execute(file, config)
	if err != nil {
		log.Error(err)
		return
	}
}