/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	//"github.com/jhalmu/go-cli-backend/data"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List alla blog entries",
	Long:  `List all blog entries in database.`,
	Run: func(cmd *cobra.Command, args []string) {
		var db *gorm.DB
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: os.Getenv("DSN_STR"),
			//PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			fmt.Println(err.Error())
			pterm.Error.Println("Error connecting to database.")
		}

		var blogtexts []BlogText
		//db.Debug().
		db.Find(&blogtexts)
		tabledata := pterm.TableData{{"ID", "Slug", "Title", "Created At", "Updated At"}}
		for _, post := range blogtexts {
			tabledata = append(tabledata, []string{fmt.Sprintf("%d", post.ID), post.Slug, post.Title, post.CreatedAt.Format("02.01.2006 15:04"), post.UpdatedAt.Format("02.01.2006 15:04:05")})
		}
		pterm.DefaultTable.WithHasHeader().WithData(tabledata).Render()

		if err != nil {
			fmt.Println(err.Error())
			pterm.Error.Println("Error listing from database.")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
