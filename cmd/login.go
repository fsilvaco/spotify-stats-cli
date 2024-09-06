package cmd

import (
	"github.com/fsilvaco/spotify-stats-cli/server"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate the user with Spotify",
	Long:  "This command authenticates the user with Spotify, allowing access to restricted features and account management",
	Run: func(cmd *cobra.Command, args []string) {
		server.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
