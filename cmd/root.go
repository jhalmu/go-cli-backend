/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"fmt"

	"github.com/spf13/cobra"
	//"log"
	"os"
)

const DSN_STR = "user=devader password=liukuPiri666 dbname=dev host=65.21.105.233 port=5432 sslmode=disable TimeZone=Europe/Helsinki"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-backend",
	Short: "Program to write blog posts in TUI",
	Long:  `Program made with Go and Cobra CLI. You can add, list and later delete blog posts from the terminal. Maeby even update.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
