package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "unpause a topic.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unpause called")
	},
}

func init() {
	rootCmd.AddCommand(unpauseCmd)
}
