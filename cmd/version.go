package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of shadowCloud",
	Long:  `All software has versions. This is shadowCloud's version number`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("shadowCloud version 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
