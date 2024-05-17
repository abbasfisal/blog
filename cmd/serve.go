package cmd

import (
	"github.com/abbasfisal/blog/pkg/bootstrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  "application will be served on host and port defined on config.yml file",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	bootstrap.Serve()
}
