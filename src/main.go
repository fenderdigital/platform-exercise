package main

import (
	"os"

	"github.com/spf13/cobra"

	"platform-exer/src/cmd"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "platform-exercise",
	}

	rootCmd.SetOut(os.Stdout)
	rootCmd.AddCommand(
		cmd.MigrateCmd,
		cmd.ServerCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
