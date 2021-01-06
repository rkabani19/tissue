package prompt

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rkabani19/ti/issue"
	"github.com/rkabani19/ti/message"
	. "github.com/rkabani19/ti/todo"
)

type Option struct {
	Option string
	Run    func(Todo, string, string, string) error
}

func Execute(todos []Todo, pat string) error {
	// TODO: create an edit option
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	owner, repo, err := getConfig()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		j, err := createPrompt(options, i+1, todo)
		if err != nil {
			return err
		}

		err = options[j].Run(todo, pat, owner, repo)
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
		Label:    "{{ .Title | cyan | bold }} {{ .Number }}: {{ .Todo.Todo }} {{ .Todo.Filepath | faint}}:{{ .Todo.LineNum | faint }}",
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

func open(todo Todo, pat string, owner string, repo string) error {
	is := issue.NewIssueService(pat, owner, repo)
	return issue.Create(todo, is)
}

func skip(todo Todo, pat string, owner string, repo string) error {
	fmt.Println(message.Warning(fmt.Sprintf(
		"Skipped TODO from %s:%d", todo.Filepath, todo.LineNum)))
	return nil
}

func exit(todo Todo, pat string, owner string, repo string) error {
	fmt.Println(message.Warning("Exiting program"))
	return nil
}

func getConfig() (string, string, error) {
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("This directory has no GitHub user.")
		return "", "", err
	}
	owner := strings.TrimSpace(string(out))

	cmd = exec.Command("git", "remote", "get-url", "origin")
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("This directory has no GitHub remote.")
		return "", "", err
	}
	split := strings.Split(string(out), owner+"/")
	repo := strings.TrimSpace(split[1][:len(split[1])-5])

	return owner, repo, nil
}
