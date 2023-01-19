package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Select(nameUser string) string {
	prompt := promptui.Select{
		Label: "Hi, " + nameUser + ". Choose which information you want to see:",
		Items: []string{"artists", "tracks"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	return result
}
