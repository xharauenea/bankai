/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new feature to your project",
	Long: `Add a new feature to your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		addFeature()
	},
}


func addFeature() {

}



func init() {
}
