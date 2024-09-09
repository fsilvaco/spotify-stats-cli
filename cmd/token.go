package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fsilvaco/spotify-stats-cli/token"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token [name]",
	Short: "token user",
	Long:  "A simple command to see token user",
	Run: func(cmd *cobra.Command, args []string) {

		var tokenData token.TokenData

		file, err := os.Open("auth_data/token.json")
		if err != nil {
			log.Fatal("error opening token file")
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&tokenData); err != nil {
			log.Fatal("error reading token from file")
		}

		fmt.Printf("Your token is: %s\n", tokenData.AccessToken)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
