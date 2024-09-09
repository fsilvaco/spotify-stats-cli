package cmd

import (
	"fmt"

	"github.com/fsilvaco/spotify-stats-cli/token"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token [name]",
	Short: "token user",
	Long:  "A simple command to see token user",
	Run: func(cmd *cobra.Command, args []string) {

		tokenData := token.GetTokenData()

		fmt.Printf("Your token is: %s\n", tokenData.AccessToken)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
