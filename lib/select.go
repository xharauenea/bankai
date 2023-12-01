package lib

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func PromptGetSelect(pc PromptContect, items []string) string {
	prompt := promptui.Select{
		Label: pc.Label,
		Items: items,
	}
	_, res, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", pc.ErrorMsg)
		os.Exit(1)
	}
	fmt.Printf("You choose %s\n", res)
	return res
}