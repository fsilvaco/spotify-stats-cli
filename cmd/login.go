/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fsilvaco/spotify-stats-cli/login"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Command to login with your Spotify Account.",
	Long:  `Command to login with your Spotify Account.`,
	Run: func(cmd *cobra.Command, args []string) {
		login.Inicialize()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
