package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "platform-exercise",
	}

	rootCmd.SetOut(os.Stdout)
	rootCmd.AddCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
