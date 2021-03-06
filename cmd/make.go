package cmd

import (
	"strings"
	"github.com/edwinvautier/go-project-cli/services"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"unicode"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "This command is used to generate new models inside your application.",
	Long:  `This command is used to generate new models inside your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("2 parameters required.")
		}

		subcommand := args[0]
		if subcommand != "entity" {
			log.Fatal("Unknown command make " + subcommand)
		}

		if !fileExists("./main.go") {
			log.Fatal("Please go inside at the root of go project to make an entity")
		}

		a := []rune(args[1])
		a[0] = unicode.ToLower(a[0])
		lowerName := string(a)
		upperName := upperCaseFirstChar(args[1])

		entity := Entity{
			Name:   		upperName,
			LowerName: 		lowerName,
			Fields: 		make([]Field, 0),
			HasDate:		false,
			HasCustomTypes: false,
		}

		filename := "./models/" + entity.Name + "Struct.go"
		// Check if a file exists for an entity with this name
		if fileExists(filename) {
			log.WithField("filename", filename).Fatal("A file already exists for this entity!")
		}

		promptUserForEntityFields(&entity)

		generateModelFile(entity)

		updateMigrations()
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

// Field represents one field of the entity the user wants to create
type Field struct {
	Name 		string
	Type 		string
	IsSlice 	bool
	SliceType	string
}

// Entity is the struct that represents the entity the user wants to create
type Entity struct {
	Name   			string
	LowerName 		string
	Fields 			[]Field
	HasDate			bool
	HasCustomTypes	bool
	ModulesPath		string
}

func promptUserForEntityFields(entity *Entity) {
	for true {
		fieldName := ""
		namePrompt := &survey.Input{
			Message: "Choose new field name (Press enter to stop adding fields)",
		}
		survey.AskOne(namePrompt, &fieldName)

		// If field name is empty then stop the function
		if fieldName == "" {
			break
		}
		field := Field{
			Name: upperCaseFirstChar(fieldName),
			IsSlice: false,
		}

		fType := ""
		typePrompt := &survey.Select{
			Message: "Choose a type for " + fieldName + ":",
			Options: services.GetTypeOptions(),
		}
		survey.AskOne(typePrompt, &fType)
		field.Type = fType

		if fType == "date" {
			entity.HasDate = true
		} else if fType == "slice" {
			field.IsSlice = true

			sliceTypePrompt := &survey.Select{
				Message: "Slice type :",
				Options: services.GetTypeOptions(),
			}
			survey.AskOne(sliceTypePrompt, &field.SliceType)

			if choosedCustomType(field.SliceType) {
				entity.HasCustomTypes = true
				field.SliceType = "models." + field.SliceType
			}
		}

		if choosedCustomType(field.Type) {
			entity.HasCustomTypes = true
			field.Type = "models." + field.Type
		}
		
		entity.Fields = append(entity.Fields, field)
	}
	wdPath, _ := os.Getwd()
	entity.ModulesPath = strings.Split(wdPath, "src/")[1]
}

func generateModelFile(entity Entity) {
	box := packr.New("My Box", "../templates")
	// Get template content as string
	templateString, err := box.FindString("models/NewModel.txt")
	if err != nil {
		log.Error(err)
		return
	}

	err = services.ExecuteTemplate(entity, entity.Name + "Struct.go","./models/", templateString)
	if err != nil {
		log.Error(err)
		return
	}
}

func upperCaseFirstChar(word string) string {
	a := []rune(word)
	a[0] = unicode.ToUpper(a[0])
	return string(a)
}

func fileExists(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func updateMigrations() {
	box := packr.New("My Box", "../templates")
	// Get template content as string
	templateString, err := box.FindString("models/migrations.txt")
	if err != nil {
		log.Error(err) 
		return
	}

	entitiesList := services.GetEntitiesList()
	err = services.ExecuteTemplate(entitiesList, "Migrations.go", "./models/", templateString)

	if err != nil {
		log.Error(err)
		return
	}
}

func choosedCustomType(cType string) bool{
	entitiesList := services.GetEntitiesList()
	log.Info(cType)
	for _, entityName := range entitiesList {
		log.Info(entityName)
        if entityName == cType {
            return true
        }
	}
	
	return false
}

