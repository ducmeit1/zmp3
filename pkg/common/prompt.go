package common

import (
	"github.com/manifoldco/promptui"
	"strconv"
)

func PromptString(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateEmptyInput,
	}

	return prompt.Run()
}

func PromptInteger(name string) (int64, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: ValidateIntegerNumberInput,
	}

	promptResult, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	parseInt, _ := strconv.ParseInt(promptResult, 0, 64)
	return parseInt, nil
}