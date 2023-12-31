/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/giteneaxharau/bankai/cmd/add"
	"github.com/giteneaxharau/bankai/cmd/create"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bankai",
	Short: "Elevate your development workflow with Bankai",
	Long:  `Bankai is a CLI tool that helps you to automate your development workflow.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCommands() {
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.AddCommand(add.AddCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bankai.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	addCommands()

}
