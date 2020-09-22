package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "This command is used to generate new models inside your application.",
	Long: `This command is used to generate new models inside your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Fatal("2 parameters required.")
		}

		subcommand := args[0]
		if subcommand != "entity" {
			log.Fatal("Unknown command make " + subcommand)
		}

		entity := Entity{
			Name: args[0],
			Fields: make([]Field, 0),
		}

		promptUserForEntityFields(&entity)
		log.Info(entity)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

// Field represents one field of the entity the user wants to create
type Field struct {
	Name string
	Type string
}

// Entity is the struct that represents the entity the user wants to create
type Entity struct {
	Name 	string
	Fields 	[]Field
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
		field := Field {
			Name: fieldName,
		}

		fType := ""
		typePrompt := &survey.Select{
			Message: "Choose a type for " + fieldName + ":",
			Options: []string{"string", "boolean", "int", "float", "date"},
		}
		survey.AskOne(typePrompt, &fType)
		field.Type = fType

		entity.Fields = append(entity.Fields, field)
	}
}