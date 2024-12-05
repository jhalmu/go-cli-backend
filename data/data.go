package data

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BlogText struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Slug    string `gorm:"unique;not null"`
	Content string `gorm:"not null"`
}

var db *gorm.DB
var dsnStr = os.Getenv("DSN_STR")

func OpenDB() error {

	//var err error

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsnStr,
		//PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		pterm.Error.Println("Error connecting to database.")
	}
	return db.Error
}

func AddToDB(title string, slug string, content string) error {
	//var err error

	db.AutoMigrate(&BlogText{})
	db.Create(&BlogText{Title: title, Slug: slug, Content: content})

	/* if err != nil {
		fmt.Println(err.Error())
	} */
	return db.Error
}

func ListDB() error {
	var blogtexts []BlogText
	//db.Find(&blogtexts)
	db.Debug().Find(&blogtexts)
	tabledata := pterm.TableData{{"ID", "Slug", "Title", "Created At", "Updated At"}}
	for _, post := range blogtexts {
		tabledata = append(tabledata, []string{fmt.Sprintf("%d", post.ID), post.Slug, post.Title, post.CreatedAt.Format("02.01.2006 15:04"), post.UpdatedAt.Format("02.01.2006 15:04:05")})
	}
	pterm.DefaultTable.WithHasHeader().WithData(tabledata).Render()
	//pterm.Error.Println("Error listing from database.")

	//return db.RowsAffected
	return nil
}

// []string{fmt.Sprintf("%d", post.ID),
