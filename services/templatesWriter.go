package services

import (
	"html/template"
	"os"
	log "github.com/sirupsen/logrus"
)

func generateFile(path string, name string, outputName string, config Config) {
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

	err = ExecuteTemplate(config, outputName, config.Path + "/" + path, templateString)
	if err != nil {
		log.Error(err)
		return
	}
}

// ExecuteTemplate takes a config, file string and computes it into a file
func ExecuteTemplate(config interface{}, outputName string, path string, templateString string) error{
	// Create the file
	file, err := os.Create(path + outputName)
	if err != nil {
		log.Error(err)
		return err
	}

	// Execute template and write file
	parsedTemplate := template.Must(template.New("template").Parse(templateString))
	err = parsedTemplate.Execute(file, config)

	return nil
}