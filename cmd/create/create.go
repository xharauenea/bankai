/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project with a boilerplate.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgm, err := cmd.Flags().GetString("package-manager")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		initializeProject(pkgm)
	},
	Example: `bankai init`,
}

type promptContect struct {
	errorMsg string
	label    string
}

var packageManagers []string = []string{"npm", "yarn", "pnpm", "bun"}

func promptGetSelect(pc promptContect) string {

	prompt := promptui.Select{
		Label: pc.label,
		Items: packageManagers,
	}
	_, res, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("You choose %s\n", res)
	return res
}

func initializeProject(pkgmFlag string) {
	_, errNextProjectExists := os.Stat("next.config.*")
	pkgmCreateCmd := map[string]string{
		"npm":  "npx create-next-app@latest",
		"yarn": "yarn create next-app",
		"pnpm": "pnpm create next-app",
		"bun":  "bunx create-next-app",
	}
	if os.IsNotExist(errNextProjectExists) {
		fmt.Println("Next project not found. Creating new project...")
		pkgm := pkgmFlag
		if len(pkgmFlag) < 1 {
			pkgm = promptGetSelect(promptContect{
				errorMsg: "Please select a package manager",
				label:    "Select a package manager",
			})
		}
		// Create new project
		cmdCreate := strings.Split(pkgmCreateCmd[pkgm], " ")
		cmd := exec.Command(cmdCreate[0], cmdCreate[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	}
}

func init() {
	CreateCmd.PersistentFlags().StringP("package-manager", "p", "npm", "Package manager to use")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
