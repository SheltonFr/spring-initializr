package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spring",
	Short: "A powerfull development-tool which provide an easy way of creating spring boot projects!",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
