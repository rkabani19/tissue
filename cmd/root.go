package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rkabani19/tissue/message"
	"github.com/rkabani19/tissue/prompt"
	"github.com/rkabani19/tissue/search"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tissue",
	Short: "A tool to allow you to convert TODO's in your code to GitHub issues.",
	Long: `This tool will find all TODO comments in your code and create a GitHub
issue for each TODO in the associated GitHub repository.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Searching all files for TODO comments...")
		todos, err := search.GetTodos(".")
		if err != nil {
			log.Fatalln("Unable to get todos.")
		}
		fmt.Printf("Found %s TODOs\n\n", message.Highlight(strconv.Itoa(len(todos))))

		err = prompt.Execute(todos, args[0])
		if err != nil {
			log.Fatalln(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
