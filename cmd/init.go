package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"github.com/SheltonFr/spring-initializr/core"
)

var createSpringBootCmd = &cobra.Command{
	Use:   "init",
	Short: "Start the creation of the Spring Boot projct!",
	Run:   runCreateSpringBootCmd,
}

var projectTypes []core.GenericType = core.GetProjectTypes()
var languages []core.GenericType = core.GetLanguages()
var packagingTypes []core.GenericType = core.GetPackagingTypes()
var javaVersions []core.GenericType = core.GetJavaVersions()
var bootVersions []core.GenericType = core.GetBootVersions()
var dependencies []core.GenericType = core.GetDependencies()

var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Project Name:"},
		Validate: survey.Required,
	},
	{
		Name: "bootVersion",
		Prompt: &survey.Select{
			Message: "Select a Spring Boot version",
			Options: core.GenericTypeToStringList(bootVersions),
		},
		Validate: survey.Required,
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "Select a project type",
			Options: core.GenericTypeToStringList(projectTypes),
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
			Options: core.GenericTypeToStringList(languages),
			Default: "Java",
		},
	},
	{
		Name: "javaVersion",
		Prompt: &survey.Select{
			Message: "Select a Java version:",
			Options: core.GenericTypeToStringList(javaVersions),
		},
		Validate: survey.Required,
	},
	{
		Name: "packaging",
		Prompt: &survey.Select{
			Message: "Packaging:",
			Options: core.GenericTypeToStringList(packagingTypes),
			Default: "Jar",
		},
	},
	{
		Name:   "groupId",
		Prompt: &survey.Input{Message: "GroupID:", Default: "com.example"},
	},
	{
		Name:   "artifactId",
		Prompt: &survey.Input{Message: "ArtifactID:", Default: "demo"},
	},
	{
		Name:   "packageName",
		Prompt: &survey.Input{Message: "Package Name:", Default: "com.example.demo"},
	},
	{
		Name:   "description",
		Prompt: &survey.Input{Message: "Description:", Default: "Demo project for spring boot"},
	},
	{
		Name: "dependencies",
		Prompt: &survey.MultiSelect{
			Message: "Choose your project's dependencies",
			Options: core.GenericTypeToStringList(dependencies),
			Description: func(value string, index int) string {
				return dependencies[index].Description
			},
		},
	},
}

type Answers struct {
	Name         string
	Type         string
	Language     string
	Packaging    string
	JavaVersion  string
	BootVersion  string
	GroupID      string
	ArtifactID   string
	Description  string
	PackageName  string
	Dependencies []string
}

func runCreateSpringBootCmd(cmd *cobra.Command, args []string) {
	answers := Answers{}

	if err := survey.Ask(qs, &answers); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(createSpringBootCmd)
}
