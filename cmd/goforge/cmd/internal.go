/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goforge/pkg/bases/internalBase"
	"goforge/pkg/components/utils"

	"github.com/spf13/cobra"
)

var componentFlag string
var baseFlag string

// newInternalCmd represents the internalBase command
var newInternalCmd = &cobra.Command{
	Use:   "internal [name]",
	Short: "Add implementation detail to a project",
	Long:  "Add implementation detail to a project",
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Usage()
		} else {
			fileName := args[0]
			if componentFlag != "" && baseFlag == "" {
				componentExists := utils.ElementExists(componentFlag, utils.Component)
				if !componentExists {
					fmt.Printf("component %s does not exist", componentFlag)
				} else {
					internalBase.New(fileName, componentFlag, utils.Component)
				}
			} else if baseFlag != "" && componentFlag == "" {
				baseExists := utils.ElementExists(baseFlag, utils.Base)
				if !baseExists {
					fmt.Printf("base %s does not exist", baseFlag)
				} else {
					internalBase.New(fileName, baseFlag, utils.Base)
				}
			} else {
				cmd.Usage()
			}
		}
	},
}

func init() {
	newCmd.AddCommand(newInternalCmd)
	newInternalCmd.PersistentFlags().StringVarP(&componentFlag, "component", "c", "", "Specify the componentFlag to which you add an implementation detail")
	newInternalCmd.PersistentFlags().StringVarP(&baseFlag, "base", "b", "", "Specify the baseFlag to which you add an implementation detail")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// internalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// internalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
