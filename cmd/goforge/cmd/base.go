/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"goforge/pkg/bases/new/base"
)

// baseCmd represents the base command
var baseCmd = &cobra.Command{
	Use:   "base",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if cmd.Parent().Name() == "new" {
			if howManyArgs != 1 {
				cmd.Parent().Usage()
			} else {
				name := args[0]
				base.New(name)
			}
		}
	},
}

func init() {
	newCmd.AddCommand(baseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// baseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// baseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
