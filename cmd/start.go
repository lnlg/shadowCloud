package cmd

import (
	"shadowCloud/app/route"
	"shadowCloud/internal/service"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start  server",
	Long:  `项目启动服务器命令`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runServer() {
	http := service.New()
	http.RegisterRoutes(route.New())
	http.Run()
}
