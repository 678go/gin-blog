package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "gin-blog",
	Short: "gin-blog",
	Long:  "gin-blog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
