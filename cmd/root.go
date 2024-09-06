package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spotify-stats",
	Short: "Spotify Stats shows your top Spotify artists and songs",
	Long:  "Spotify Stats is a CLI tool that displays your top Spotify artists, songs, and genres.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Spotify Stats!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
