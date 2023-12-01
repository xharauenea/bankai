/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/giteneaxharau/bankai/lib"
	"github.com/giteneaxharau/bankai/lib/config"
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

func initializeProject(pkgmFlag string) {
	_, errNextProjectExists := os.Stat("next.config.*")
	if os.IsNotExist(errNextProjectExists) {
		fmt.Println("Next project not found. Creating new project...")
		pkgm := pkgmFlag
		if len(pkgmFlag) < 1 {
			pkgm = lib.PromptGetSelect(lib.PromptContect{
				ErrorMsg: "Please select a package manager",
				Label:    "Select a package manager",
			}, lib.PackageManagers)
		}
		// Create new project
		cmdCreate := strings.Split(lib.PackageManagersMap[pkgm], " ")
		cmd := exec.Command(cmdCreate[0], cmdCreate[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			var execErr *exec.ExitError
			if errors.As(err, &execErr) {
				fmt.Println("Error creating new project. Please try again.")
				os.Exit(1)
			}
		}
		isSrcDir := lib.CheckNextjsProjectDir()
		packageManager := pkgm
		alias := lib.PromptGetInput(lib.PromptContect{
			Label:   "Enter an alias for this project",
			ErrorMsg: "Please enter an alias for this project",
		})
		config.CreateConfig(config.Config{
			Alias:          alias,
			PackageManager: packageManager,
			HasSrcDir:       isSrcDir,
		})
	} else {
		fmt.Println("Next project found. Skipping...")
	}
}

func init() {
	CreateCmd.Flags().StringP("package-manager", "p", "", "Package manager to use")
	CreateCmd.Aliases = []string{"create"}
}
