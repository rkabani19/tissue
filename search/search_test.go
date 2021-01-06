package search

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetTodos(t *testing.T) {
	todoComment := "Test Comment"
	filepath := "./test.txt"

	tables := []struct {
		fileInput []byte
		todo      string
	}{
		{[]byte("hello\ngo\n// TODO:" + todoComment + "\n wow \n // comment"), todoComment},
		{[]byte("hello\ngo\nrandom text // TODO:" + todoComment + "\n wow \n // comment"), todoComment},
	}

	for _, table := range tables {
		err := ioutil.WriteFile(filepath, table.fileInput, 0644)
		if err != nil {
			t.Fatalf("Could not write file. Error: %s", err)
		}

		todos, err := GetTodos(filepath)
		os.RemoveAll(filepath)

		if err != nil {
			t.Errorf("GetTodos threw an error: %s", err)
		}

		if len(todos) != 1 {
			t.Errorf("GetTodos returned wrong list of todos: %v", todos)
		}

		if todos[0].Todo != todoComment {
			t.Errorf("GetTodos returned wrong todo: %s\n", todos[0].Todo)
		}
	}
}
