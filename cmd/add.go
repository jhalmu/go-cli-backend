/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/gosimple/slug"
	"os"
	//"github.com/jhalmu/go-cli-backend/data"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BlogText struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Slug    string `gorm:"not null"`
	Content string `gorm:"not null"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new entry",
	Long:  `Add new entry to the database.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("Please enter (Unique) Title").Show()
		slug := slug.Make(title)
		content, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().WithDefaultText("Please enter Content").Show()

		fmt.Println("Title: ", title)
		fmt.Println("Slug: ", slug)
		fmt.Println("Content: ", content)

		var db *gorm.DB
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: os.Getenv("DSN_STR"),
			//PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			fmt.Println(err.Error())
			pterm.Error.Println("Error connecting to database.")
		}

		db.AutoMigrate(&BlogText{})
		db.Create(&BlogText{Title: title, Slug: slug, Content: content})

		if err != nil {
			fmt.Println(err.Error())
			pterm.Error.Println("Error adding post to database.")
		}
		pterm.Success.Println("Post added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
