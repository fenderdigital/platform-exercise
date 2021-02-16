package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"platform-exer/src/app"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run the web server",
	RunE:  server,
	Args:  cobra.ExactArgs(0),
}

func server(_ *cobra.Command, _ []string) error {
	s, err := app.InitServices()
	if err != nil {
		return err
	}

	r, err := app.InitRouter(s, os.Getenv("GIN_MODE"))
	if err != nil {
		return err
	}

	return r.Run(fmt.Sprintf(":%s", os.Getenv("LISTEN_ADDR")))
}
