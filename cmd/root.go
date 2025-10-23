package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goframe",
	Short: "GoFrame is a lightweight Go application framework",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚙️ No subcommand provided, starting server by default...")
		serveCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
