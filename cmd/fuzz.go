/*
Copyright © 2024 NAME HERE <lotfi.kaddari.lp2@gmail.com>
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

var fuzzCmd = &cobra.Command{
	Use:   "fuzz",
	Short: "the fuzz command helps you to fuzz url",
	Long: `the fuzz command sends multiple requests to a url just by providing a domain and a file that containe the routes that you want to fuzz 
	. For example:
	fuzz -u https://example.com -f ~/list.txt 
	the list must contain the routes like this : /home , /users ...
	each route should be in one line `,
	Run: func(cmd *cobra.Command, args []string) {
		fuzzGenerate(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(fuzzCmd)
	fuzzCmd.Flags().StringP("Filepath", "f", "", "Provide a path to the routes file")
	fuzzCmd.Flags().StringP("Url", "u", "", "Provide the Domian that you want to fuzz")
	fuzzCmd.Flags().StringP("Method", "m", "", "Provide the method (Get , Post , Patch)")
}

type Response struct {
	url    string
	status int
}

func fuzzGenerate(cmd *cobra.Command, args []string) {
	domain, _ := cmd.Flags().GetString("Url")
	filepath, _ := cmd.Flags().GetString("Filepath")
	method, _ := cmd.Flags().GetString("Method")

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
				var res *http.Response
				var err error
				if method == "Get" {
					res, err = http.Get(fullUrl)
				}
				if method == "Post" {
					res, err = http.Post(fullUrl, "application/x-www-form-urlencoded", nil)
				}
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
