package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/rkabani19/tissue/issue"
	"github.com/rkabani19/tissue/message"
	. "github.com/rkabani19/tissue/todo"
)

type Option struct {
	Option string
	Run    func(Todo, issue.IssueService) error
}

func Execute(todos []Todo, pat string) error {
	// TODO: create an edit option
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	is, err := issue.NewIssueService(pat)
	if err != nil {
		return err
	}
	for i, todo := range todos {
		j, err := createPrompt(options, i+1, todo)
		if err != nil {
			return err
		}

		err = options[j].Run(todo, is)
		if err != nil {
			return err
		}

		if options[j].Option == options[len(options)-1].Option {
			break
		}
	}

	return nil
}

func createPrompt(options []Option, num int, todo Todo) (int, error) {
	todoText := struct {
		Title  string
		Number int
		Todo   Todo
	}{
		Title:  "Issue",
		Number: num,
		Todo:   todo,
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ .Title | cyan | bold }} {{ .Number | cyan | bold }}: {{ .Todo.Todo }} {{ .Todo.Filepath | faint}}:{{ .Todo.LineNum | faint }}",
		Active:   "\U000027A4 {{ .Option }}",
		Inactive: "  {{ .Option | faint }}",
	}

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

func open(todo Todo, is issue.IssueService) error {
	return issue.Create(todo, is)
}

func skip(todo Todo, is issue.IssueService) error {
	fmt.Println(message.Warning(fmt.Sprintf(
		"Skipped TODO from %s:%d", todo.Filepath, todo.LineNum)))
	return nil
}

func exit(todo Todo, is issue.IssueService) error {
	fmt.Println(message.Warning("Exiting program"))
	return nil
}
