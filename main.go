package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")

	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Server is running on port: ", PORT)
}
