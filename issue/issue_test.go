package issue

import (
	"errors"
	"testing"

	"github.com/google/go-github/github"
	. "github.com/rkabani19/tissue/todo"
)

var createIssueMock func(Todo) (*github.Issue, *github.Response, error)

type issueServiceMock struct{}

func (is issueServiceMock) createIssue(todo Todo) (*github.Issue, *github.Response, error) {
	return createIssueMock(todo)
}

func TestCreate(t *testing.T) {
	is := issueServiceMock{}
	issNum, issTitle := 1, "Test"

	createIssueMock = func(Todo) (*github.Issue, *github.Response, error) {
		return &github.Issue{Number: &issNum, Title: &issTitle}, nil, nil
	}

	err := Create(Todo{}, is)
	if err != nil {
		t.Fatal(err)
	}

	createIssueMock = func(Todo) (*github.Issue, *github.Response, error) {
		return nil, nil, errors.New("error")
	}

	err = Create(Todo{}, is)
	if err == nil {
		t.Error("Expected Create to throw an error but got nil")
	}
}
