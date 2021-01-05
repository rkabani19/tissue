package issue

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/rkabani19/ti/message"
	"github.com/rkabani19/ti/todo"
	"golang.org/x/oauth2"
)

const label = "TODO"

func Create(todo todo.Todo, pat string, repoOwner string, repo string) {
	ctx, client := authenticate(pat)
	body := fmt.Sprintf("%s:%d", todo.Filepath, todo.LineNum)

	issue, _, err := client.Issues.Create(
		ctx, repoOwner, repo, &github.IssueRequest{
			Title:  &todo.Todo,
			Body:   &body,
			Labels: &[]string{label},
		})

	if err != nil {
		fmt.Println(message.Error("Unable to create issue."))
		panic(err)
	}

	successMsg := fmt.Sprintf("Opened issue #%d: %s", *issue.Number, *issue.Title)
	fmt.Println(message.Success(successMsg))
}

func authenticate(pat string) (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}
