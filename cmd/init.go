package cmd

import (
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init",
	Long:  "init",
	RunE: func(cmd *cobra.Command, args []string) error {
		//db := config.InitDB()
		//db.AutoMigrate()
		return nil
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
