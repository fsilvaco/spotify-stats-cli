package cmd

import (
	"fmt"
	"strconv"

	"github.com/fsilvaco/spotify-stats-cli/spotify"
	"github.com/spf13/cobra"
)

var topCmd = &cobra.Command{
	Use:   "top [item]",
	Short: "token artists or tracks",
	Long:  "A simple command to see token user",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		arg := args[0]

		data := spotify.GetTopItems(arg)

		for i := 0; i < len(data.Items); i++ {
			position := strconv.Itoa(i + 1)

			fmt.Println(position + " - " + data.Items[i].Name)
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}
