/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"goforge/pkg/bases/app"
)

// addAppCmd represents the addApp command
var addAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Add an app to the project.",
	Long:  `Add an app to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
		} else {
			name := args[0]
			app.New(name)
		}
	},
}

// rmAppCmd represents the rmApp command
var rmAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Remove an app from the project.",
	Long:  "Remove an app from the project.",
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
		} else {
			name := args[0]
			app.Remove(name)
		}
	},
}

func init() {
	newCmd.AddCommand(addAppCmd)
	rmCmd.AddCommand(rmAppCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
