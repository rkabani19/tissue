package search

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"sync"

	. "github.com/rkabani19/tissue/todo"
)

const (
	// TODO: Support diff file types
	comment    = "//"
	todoString = "TODO"
)

var wg sync.WaitGroup

// GetTodos walks through directory tree from startDir and parses files to
// extract commented todos and return them
func GetTodos(startDir string) ([]Todo, error) {
	mutex := &sync.Mutex{}
	todos := make([]Todo, 0)

	wg.Add(1)
	e := traverseFiles(startDir, func(fp string) error {
		file, err := os.Open(fp)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		line := 1
		for scanner.Scan() {
			todo := getTodo(scanner.Text())
			if todo != "" {
				mutex.Lock()
				todos = append(todos, Todo{
					LineNum:  line,
					Filepath: fp,
					Todo:     todo,
				})
				mutex.Unlock()
			}
			line++
		}

		if err := scanner.Err(); err != nil {
			return err
		}

		return nil
	})
	wg.Wait()

	return todos, e
}

func getTodo(line string) (todo string) {
	split := strings.Split(line, comment)
	if len(split) < 2 {
		return ""
	}

	split = strings.Split(line, todoString)
	if len(split) < 2 {
		return ""
	}

	todo = split[1]
	if strings.HasPrefix(todo, ":") {
		todo = strings.TrimSpace(todo[1:])
	}

	return todo
}

func traverseFiles(searchDir string, f func(fp string) error) error {
	defer wg.Done()

	e := filepath.Walk(searchDir, func(path string, fi os.FileInfo, err error) error {
		if fi.Mode().IsRegular() {
			f(path)
		} else if fi.IsDir() && path != searchDir {
			wg.Add(1)
			go traverseFiles(path, f)
			return filepath.SkipDir
		}
		return err
	})

	if e != nil {
		return e
	}

	return nil
}
