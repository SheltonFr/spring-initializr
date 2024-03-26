package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Start the creation of the Spring Boot projct!",
	Run:   testingThings,
}

var qs = []*survey.Question{
	{
		Name: "Type",
		Prompt: &survey.Select{
			Message: "Select the project type",
			Options: []string{"Gradle - Groove", "Gradle - Kotlin", "Gradle Config", "Maven", "Maven POM"},
			Default: "Maven",
		},
	},
}

func testingThings(cmd *cobra.Command, args []string) {
	answers := struct{}{}
	survey.Ask(qs, &answers)
}
func runSpringInitializr(cmd *cobra.Command, args []string) {

	result, err := http.Get("https://start.spring.io/metadata/client")
	if err != nil {
		panic(err)
	}

	defer result.Body.Close()

	file, err := os.Create("registries.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, result.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Done!!!")
}

func init() {
	rootCmd.AddCommand(createCmd)
}
