package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "create-project",
	Short: "A powerfull development-tool which provide an easy way of creating spring boot projects!",
	Run:   runSpringInitializr,
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
