package lib

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptGetInput(pc PromptContect) string {
	validate := func(input string) error {
		if len(input) < 3 {
			return fmt.Errorf(pc.ErrorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", pc.ErrorMsg)
		return ""
	}

	return result
}