package commands

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type promptContent struct {
	label    string
	errorMsg string
}

func promptGetInput(pc promptContent, items []string) string {
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.label,
			Items: items,
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

	// fmt.Printf("Input: %s\n", result)

	return result
}
