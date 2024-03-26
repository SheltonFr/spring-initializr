package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spring",
	Short: "Spring is Spring Boot project generator tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Im working")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
