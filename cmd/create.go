package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/SheltonFr/spring-initializr/core"
)

var createSpringBootCmd = &cobra.Command{
	Use:   "create-sb-project",
	Short: "Start the creation of the Spring Boot projct!",
	Run:   runCreateSpringBootCmd,
}

var projectTypes []core.GenericType = core.GetProjectTypes()
var languages []core.GenericType = core.GetLanguages()
var packagingTypes []core.GenericType = core.GetPackagingTypes()

func genericTypeToStringList(items []core.GenericType) []string {
	var newItems []string
	for _, item := range items {
		newItems = append(newItems, item.Name)
	}
	return newItems
}

var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Project Name:"},
		Validate: survey.Required,
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Select a project type",
			Options: genericTypeToStringList(projectTypes),
			Description: func(value string, index int) string {
				return projectTypes[index].Description
			},
			Default: "Maven",
		},
	},
	{
		Name: "language",
		Prompt: &survey.Select{
			Message: "Select a language",
			Options: genericTypeToStringList(languages),
			Default: "Java",
		},
	},
	{
		Name: "packaging",
		Prompt: &survey.Select{
			Message: "Packaging:",
			Options: genericTypeToStringList(packagingTypes),
			Default: "Jar",
		},
	},
}

func runCreateSpringBootCmd(cmd *cobra.Command, args []string) {
	answers := struct {
		Name      string
		Type      string
		Language  string
		Packaging string
	}{}
	if err := survey.Ask(qs, &answers); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(createSpringBootCmd)
}
