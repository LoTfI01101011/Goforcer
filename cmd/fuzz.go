/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

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

	fuzzCmd.Flags().StringP("filepath", "f", "", "Give a path to the routes file")

}

func fuzzGenerate(cmd *cobra.Command, args []string) {
	// url, _ := cmd.Flags().GetString("Url")
	filepath, _ := cmd.Flags().GetString("filepath")

	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("there was a problem with your file")
	}
	defer file.Close()
	arrFile := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arrFile = append(arrFile, scanner.Text())
	}
	fmt.Print(arrFile)

}
