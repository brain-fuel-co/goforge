package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// prereqsCmd represents the prereqs command
var prereqsCmd = &cobra.Command{
	Use:   "prereqs",
	Short: "Installs prerequisites for using GoForge (Windows Only)",
	Long: `Installs prerequisites for using GoForge (Windows Only):
  - Scoop
  - Make
  - Git
  - Cobra
Requires that you have Go v1.18.0 or later installed.`,
	Run: func(cmd *cobra.Command, args []string) {
		installScoopCmd := "iex (new-object net.webclient).downloadstring('https://get.scoop.sh')"
		installMakeCmd := "scoop install make"
		installGitCmd := "scoop install git"
		installCobraCmd := "go install github.com/spf13/cobra-cli@latest"

		fmt.Println("Installing Scoop")
		installScoop := exec.Command("powershell", "-Command", installScoopCmd)
		installScoopOut, installScoopErr := installScoop.CombinedOutput()
		if installScoopErr != nil {
			fmt.Println("Error installing Scoop")
			return
		} else {
			fmt.Println(strings.TrimSpace(string(installScoopOut)))
		}

		installMake := exec.Command("powershell", "-Command", installMakeCmd)
		installMakeOut, installMakeErr := installMake.CombinedOutput()
		if installMakeErr != nil {
			fmt.Println("Error installing Make")
			return
		} else {
			fmt.Println(strings.TrimSpace(string(installMakeOut)))
		}

		installGit := exec.Command("powershell", "-Command", installGitCmd)
		installGitOut, installGitErr := installGit.CombinedOutput()
		if installGitErr != nil {
			fmt.Println("Error installing Git")
			return
		} else {
			fmt.Println(strings.TrimSpace(string(installGitOut)))
		}

		installCobra := exec.Command("powershell", "-Command", installCobraCmd)
		installCobraOut, installCobraErr := installCobra.CombinedOutput()
		if installCobraErr != nil {
			fmt.Println("Error installing Cobra")
			return
		} else {
			fmt.Println(strings.TrimSpace(string(installCobraOut)))
		}
	},
}

func init() {
	installCmd.AddCommand(prereqsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prereqsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prereqsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
