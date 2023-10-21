package cmd

import (
	"github.com/spf13/cobra"
	"goforge/pkg/bases/component"
)

// newComponentCmd represents the component command with new as the parent
var newComponentCmd = &cobra.Command{
	Use:   "component [name]",
	Short: "Add a component to the project.",
	Long:  `Add a component to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Usage()
		} else {
			name := args[0]
			component.New(name)
		}
	},
}

// rmComponentCmd represents the component command with rm as the parent
var rmComponentCmd = &cobra.Command{
	Use:   "component [name]",
	Short: "Remove a component from the project.",
	Long:  "Remove a component from the project.",
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Usage()
		} else {
			name := args[0]
			component.Remove(name)
		}
	},
}

func init() {
	newCmd.AddCommand(newComponentCmd)
	rmCmd.AddCommand(rmComponentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// componentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// componentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
