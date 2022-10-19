package main

import (
	"final_project_hacktiv8/server"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
