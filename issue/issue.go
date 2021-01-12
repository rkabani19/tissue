package issue

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/google/go-github/github"
	"github.com/rkabani19/tissue/message"
	. "github.com/rkabani19/tissue/todo"
	"golang.org/x/oauth2"
)

type IssueService interface {
	createIssue(Todo) (*github.Issue, *github.Response, error)
}

type issueService struct {
	pat   string
	owner string
	repo  string
}

const label = "TODO"

// NeNewIssueService returns issueService struct
func NewIssueService(pat string) (IssueService, error) {
	owner, repo, err := getConfig()
	if err != nil {
		return issueService{}, err
	}

	return issueService{pat: pat, owner: owner, repo: repo}, nil
}

// Create will create a GitHub issue in the project's repository
func Create(todo Todo, is IssueService) error {
	issue, _, err := is.createIssue(todo)

	if err != nil {
		fmt.Println(message.Error("Unable to create issue."))
		return err
	}

	successMsg := fmt.Sprintf("Opened issue #%d: %s", *issue.Number, *issue.Title)
	fmt.Println(message.Success(successMsg))
	return nil
}

func (is issueService) authenticate() (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: is.pat},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}

func (is issueService) createIssue(todo Todo) (
	*github.Issue, *github.Response, error) {
	ctx, client := is.authenticate()
	body := fmt.Sprintf("%s:%d", todo.Filepath, todo.LineNum)
	return client.Issues.Create(
		ctx, is.owner, is.repo, &github.IssueRequest{
			Title:  &todo.Todo,
			Body:   &body,
			Labels: &[]string{label},
		})
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
