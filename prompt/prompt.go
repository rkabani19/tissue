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

func Execute(todos []Todo) error {
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	for _, todo := range todos {
		i, err := createPrompt(options, todo)
		if err != nil {
			return err
		}

		options[i].Run()
		if options[i].Option == options[len(options)-1].Option {
			break
		}
	}

	return nil
}

func createPrompt(options []Option, todo Todo) (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "",
		Active:   "\U00001433 {{ .Option }}",
		Inactive: "  {{ .Option | faint }}",
	}

	todoText := fmt.Sprintf("Issue: %s -- %s:%d",
		todo.Todo, todo.Filepath, todo.LineNum)

	prompt := promptui.Select{
		Label:        todoText,
		Items:        options,
		Templates:    templates,
		Size:         4,
		HideHelp:     true,
		HideSelected: true,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return -1, err
	}

	return i, nil
}

func open() {
	fmt.Println("Open issue.")
}

func skip() {
	fmt.Println("Skip issue.")
}

func exit() {
	fmt.Println("Exiting program.")
}
