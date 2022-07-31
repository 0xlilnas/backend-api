/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "memorize",
	Short: "A tool to help you memorize better",
	Long:  `To become a super productive person,a good memory helps to make you work faster without any helps.So this app help you imprint the memory in your work flow in your day to day tasks`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
