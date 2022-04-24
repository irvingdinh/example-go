package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/irvingdinh/example-go/internal/config"
)

var rootCmd = &cobra.Command{
	Short: "Example HTTP service with Go",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	config.Load()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
