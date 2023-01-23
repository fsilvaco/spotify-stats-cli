package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var selectLabel string
var selectItems []string

func selectOption(label string, items []string) string {

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	return result
}

func SelectTypeForSearch(userName string) string {
	selectLabel = "Hi, " + userName + ". Choose which information you want to see"
	selectItems = []string{"artists", "tracks"}
	typeForSearch := selectOption(selectLabel, selectItems)

	return typeForSearch
}

func SelectTimeRange() (string, string) {
	var slugTimeRange string
	selectLabel = "In which period?"
	selectItems = []string{"Last 4 weeks", "Last 6 months", "Several years"}
	timeRange := selectOption(selectLabel, selectItems)

	switch timeRange {
	case "Several years":
		slugTimeRange = "long_term"
	case "Last 6 months":
		slugTimeRange = "medium_term"
	case "Last 4 weeks":
		slugTimeRange = "short_term"
	}

	return timeRange, slugTimeRange
}
