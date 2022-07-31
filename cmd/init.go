/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/0xlilnas/cli-app/data"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initiliaze memorize",
	Long:  `Initiliaze memorize database and data table`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		data.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
