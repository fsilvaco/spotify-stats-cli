/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fsilvaco/spotify-stats-cli/spotify"
	"github.com/spf13/cobra"
)

const (
	PORT             string = ":8080"
	CLIENT_ID        string = "33cfd3aa25f144a18bc39ba4f7b6302c"
	SCOPE            string = "user-top-read"
	REDIRECT_URI     string = "http://localhost" + PORT + "/auth"
	SPOTIFY_AUTH_URL string = "https://accounts.spotify.com/authorize?" + "client_id=" + CLIENT_ID + "&response_type=token&scope=" + SCOPE + "&redirect_uri=" + REDIRECT_URI
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Command to login with your Spotify Account.",
	Long:  `Command to login with your Spotify Account.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := spotify.Server{Port: PORT, SpotifyAuthURL: SPOTIFY_AUTH_URL}
		server.Server()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
