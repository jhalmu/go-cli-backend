/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	//"fmt"
	"github.com/jhalmu/go-cli-backend/cmd"
	"github.com/joho/godotenv"
	"log"
	//"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//dsnStr := os.Getenv("DSN_STR")
	//fmt.Println(dsnStr)

	cmd.Execute()

}
