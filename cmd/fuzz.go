/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
		fmt.Println("fuzz called")
	},
}

func fuzzCmd(cmd *cobra.Command, args []string) {
	// url, _ := cmd.Flags().GetString("Url")
	routes, _ := cmd.Flags().GetString("Routes")

	content, err := os.ReadFile(routes)

	if err != nil {
		fmt.Printf("There is a problem with your file path")
	}
	fmt.Printf("Content of %s:\n%s\n", routes, string(content))

}

func init() {
	rootCmd.AddCommand(fuzzCmd)

	fuzzCmd.Flags().StringP("filepath", "f", "", "Give a path to the routes file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fuzzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fuzzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
