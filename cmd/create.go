package cmd

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/edwinvautier/go-project-cli/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is used to initialize a new application.",
	Long:  `This command is used to initialize a new application.`,
	Run: func(cmd *cobra.Command, args []string) {
		appName := strings.Join(args, "-")

		if len(args) != 1 {
			userRefused := promptUserForAppName(appName)

			if userRefused {
				return
			}
		}
		var username string
		if viper.GetString("username") != "" {
			username = viper.GetString("username")
		} else {
			// Ask the user for it's github username
			username = promptUserForGitUsername()
			viper.Set("username", username)
			viper.WriteConfig()
		}

		// Ask the user for the modules he wants
		modules := promptUserForModules()

		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Couldn't find the current directory.")
		}

		services.CreateStructure(path+"/"+appName, modules, username, appName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Info("Exiting...")
		os.Exit(1)
	}()
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func promptUserForAppName(appName string) bool {
	log.Warn("1 parameter expected: application name.")

	answer := true
	prompt := &survey.Confirm{
		Message: "Do you want to use this name: " + appName + " ?",
		Default: true,
	}
	survey.AskOne(prompt, &answer)

	if answer == false {
		return true
	}
	return false
}
func promptUserForModules() []string {
	modules := []string{}
	prompt := &survey.MultiSelect{
		Message: "What modules do you want:",
		Options: []string{"Router", "Database", "Docker"},
	}
	survey.AskOne(prompt, &modules)

	return modules
}

func promptUserForGitUsername() string{
	var name string
	prompt := &survey.Input{
		Message: "What's your github username ?",
	}
	survey.AskOne(prompt, &name)

	return name
}