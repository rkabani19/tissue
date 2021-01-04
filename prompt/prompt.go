package prompt

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/rkabani19/ti/issue"
	. "github.com/rkabani19/ti/todo"
)

type Option struct {
	Option string
	Run    func(Todo, string, string, string)
}

func Execute(todos []Todo, pat string) error {
	options := []Option{
		{Option: "Open Issue", Run: open},
		{Option: "Skip Issue", Run: skip},
		{Option: "Exit", Run: exit},
	}

	repo, owner, err := getConfig()
	if err != nil {
		return err
	}

	for _, todo := range todos {
		i, err := createPrompt(options, todo)
		if err != nil {
			return err
		}

		options[i].Run(todo, pat, owner, repo)
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

func open(todo Todo, pat string, owner string, repo string) {
	issue.Create(todo, pat, owner, repo)
}

func skip(todo Todo, pat string, owner string, repo string) {
}

func exit(todo Todo, pat string, owner string, repo string) {
	fmt.Println("Exiting program.")
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
