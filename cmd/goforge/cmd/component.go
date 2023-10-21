package cmd

import (
	"github.com/spf13/cobra"
	"goforge/pkg/bases/component"
)

// newComponentCmd represents the component command with new as the parent
var newComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
		} else {
			name := args[0]
			component.New(name)
		}
	},
}

// rmComponentCmd represents the component command with rm as the parent
var rmComponentCmd = &cobra.Command{
	Use:   "component",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		howManyArgs := len(args)
		if howManyArgs != 1 {
			cmd.Parent().Usage()
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
