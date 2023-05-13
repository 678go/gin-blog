package cmd

import "github.com/spf13/cobra"

var filePath string

var RootCmd = &cobra.Command{
	Use:   "gin-blog",
	Short: "gin-blog",
	Long:  "gin-blog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&filePath, "filePath", "f", "/config/dev.ini", "dev env")
}
