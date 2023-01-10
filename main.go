package main

import (
	"fmt"
	"log"
	"os"
	"social-media/database"
	"social-media/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectToDb()

	r := routes.Router()

	if PORT, ok := os.LookupEnv("PORT"); ok {
		r.Run(fmt.Sprintf("localhost:%s", PORT))
	} else {
		r.Run("localhost:2000")
	}

}
