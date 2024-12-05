/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	//"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show one selected blog entry",
	Long:  `Show only one selected blog entry from database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("show called")

	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
