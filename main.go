package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to RSS Aggr Service")

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not declared in the environment")
	}

	fmt.Println("Port:" + portString)
}