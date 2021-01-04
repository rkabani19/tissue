package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
	. "github.com/rkabani19/ti/todo"
)

type Option struct {
	Option string
	Run    func()
}

func Execute(todos []Todo) {
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	for _, todo := range todos {
		i := createPrompt(options, todo)
		options[i].Run()
		if options[i].Option == options[len(options)-1].Option {
			break
		}
	}
}

func createPrompt(options []Option, todo Todo) int {
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
		panic(err)
	}

	return i
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
