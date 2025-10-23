package cmd

import (
	"fmt"
	"goframe/internal/app/routes"
	"goframe/internal/core/bootstrap"
	"goframe/internal/core/plugins"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		app := bootstrap.NewApp()
		r := app.Router
		cfg := app.Config

		routes.RegisterAPIRoutes(r, cfg)
		routes.RegisterWebRoutes(r)
		plugins.PrintRegisteredRoutes(r)


		addr := fmt.Sprintf(":%s", app.Config.AppPort)
		fmt.Printf("🚀 Server running at http://localhost%s\n", addr)
		r.Run(addr)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
