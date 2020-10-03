package services

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

// GetEntitiesList returns a slice of strings with all the entities names found in the models/ dir
func GetEntitiesList() []string {
	files, err := ioutil.ReadDir("./models")
    if err != nil {
		log.Fatal(err)
    }
	entities := make([]string, 0)
    for _, file := range files {
        if file.Name() != "DBConnector.go" && file.Name() != "Migrations.go" {
			entities = append(entities, strings.Split(file.Name(), "Struct.go")[0])
		}
	}
	
	return entities
}

func GetTypeOptions() []string {
	entitiesList := GetEntitiesList()
	options := []string{"string", "boolean", "int", "float", "date", "slice"}
	for _, entity := range entitiesList {
		options = append(options, entity)
	}

	return options
}