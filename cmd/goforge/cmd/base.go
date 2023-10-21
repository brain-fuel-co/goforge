/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"goforge/pkg/bases/base"
)

// addBaseCmd represents the base command
var addBaseCmd = &cobra.Command{
	Use:   "base",
	Short: "Add a base to the project.",
	Long:  `Add a base to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
		} else {
			name := args[0]
			base.New(name)
		}
	},
}

// rmBaseCmd represents the base command
var rmBaseCmd = &cobra.Command{
	Use:   "base",
	Short: "Remove a base from the project.",
	Long:  "Remove a base from the project.",
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
		} else {
			name := args[0]
			base.Remove(name)
		}
	},
}

func init() {
	newCmd.AddCommand(addBaseCmd)
	rmCmd.AddCommand(rmBaseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addBaseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addBaseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
