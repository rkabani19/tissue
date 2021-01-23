package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rkabani19/tissue/message"
	"github.com/rkabani19/tissue/search"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all TODOs in your project",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Searching all files for TODO comments...")
		todos, err := search.GetTodos(".")
		if err != nil {
			log.Fatalln("Unable to get todos.")
		}
		fmt.Printf("Found %s TODOs\n\n", message.Highlight(strconv.Itoa(len(todos))))

		for i, todo := range todos {
			fmt.Printf("%s %s: %s %s\n",
				message.Highlight("TODO"), message.Highlight(strconv.Itoa(i+1)), todo.Todo,
				message.Faint(fmt.Sprintf("(%s:%d)", todo.Filepath, todo.LineNum)))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
