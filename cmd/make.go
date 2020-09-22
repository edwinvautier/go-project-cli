package cmd

import (
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

		entityName := args[1]

		entityStruct := promptUserForEntityFields()
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}

func promptUserForEntityFields() interface{} {
	var entity interface {}
}