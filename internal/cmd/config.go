package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/irvingdinh/example-go/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the configuration in JSON format",
	Run: func(cmd *cobra.Command, args []string) {
		buff, err := json.MarshalIndent(config.AsMap(), "", "  ")
		if err != nil {
			log.Fatalf("json.MarshalIndent returns error: %s", err.Error())
		}

		fmt.Println(string(buff))
	},
}
