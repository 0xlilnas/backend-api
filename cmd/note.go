/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A place where your place notes and memorize",
	Long:  `A place where you write all the notes and memorize including command,syntax and logic`,
}

func init() {
	rootCmd.AddCommand(noteCmd)
}
