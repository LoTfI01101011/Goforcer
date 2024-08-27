/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// fuzzCmd represents the fuzz command
var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fuzzGenerate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(fuzzCmd)

	fuzzCmd.Flags().StringP("filepath", "f", "", "Provide a path to the paths file")
	fuzzCmd.Flags().StringP("Url", "u", "", "Provide the Domian that you want to fuzz")

}

type Response struct {
	url    string
	status int
}

func fuzzGenerate(cmd *cobra.Command, args []string) {
	domain, _ := cmd.Flags().GetString("Url")
	filepath, _ := cmd.Flags().GetString("filepath")

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("There was a problem with your file path")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	responseCh := make(chan Response)
	maxConcurentRequest := 5
	sem := make(chan struct{}, maxConcurentRequest)
	var wg sync.WaitGroup
	go func() {
		defer close(responseCh)

		for scanner.Scan() {
			sem <- struct{}{}
			wg.Add(1)
			fullUrl := domain + scanner.Text()
			go func(fullUrl string) {
				defer func() { <-sem }()
				defer wg.Done()
				res, err := http.Get(fullUrl)
				if err != nil {
					return
				}
				responseCh <- Response{url: fullUrl, status: res.StatusCode}
			}(fullUrl)
		}
		wg.Wait()
	}()

	for response := range responseCh {
		fmt.Printf("This URL %s responded with the status of: %d\n", response.url, response.status)
	}

}
