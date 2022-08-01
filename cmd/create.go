/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/0xlilnas/cli-app/data"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new memory",
	Long:  `create new memory and start memorize`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	noteCmd.AddCommand(createCmd)
}

func promptGetInput(pc promptContent) string {
	//validate
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}

		return nil
	}

	//templates
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	//prompt
	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created memory 📖 %s\n", result)
	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"💻 linux", "🔥 programming", "💽 logic"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "🍢 Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input %s\n", result)

	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word that you want to memorize",
		"What do you want to memorize?",
	}

	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide simple acronyms for memorization",
		fmt.Sprintf("What is the simplest acronym  of %s?", word),
	}

	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category",
		fmt.Sprintf("What category does %s belong to?", word),
	}

	category := promptGetSelect(categoryPromptContent)

	data.InsertNote(word, definition, category)
}
