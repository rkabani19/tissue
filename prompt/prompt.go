package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type Option struct {
	Option string
	Run    func()
}

func Execute() {
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	templates := &promptui.SelectTemplates{
		Label:    "",
		Active:   "\U00001433 {{ .Option }}",
		Inactive: "  {{ .Option | faint }}",
	}

	prompt := promptui.Select{
		Label:        "",
		Items:        options,
		Templates:    templates,
		Size:         4,
		HideHelp:     true,
		HideSelected: true,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	options[i].Run()
}

func open() {
	fmt.Println("Open issue.")
}
func skip() {
	fmt.Println("Skip issue.")
}
func exit() {
	fmt.Println("Exit program.")
}
