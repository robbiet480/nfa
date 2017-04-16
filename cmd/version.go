package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the version of nfa",
	Long:  `Return the version of nfa`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("1.0.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
