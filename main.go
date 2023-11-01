package main

import (
	"log"
	"os"
	"time"

	"github.com/hjunior29/chain-processor/core"
	"github.com/hjunior29/chain-processor/db"
	"github.com/joho/godotenv"
)

func loadenv() {
	if os.Getenv("DOCKER_ENV") != "" {
		log.Println("Running in Docker, skipping .env file loading")
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	loadenv()

	err := db.InitDatabase()
	if err != nil {
		log.Fatalf("Error during db connecting: %v", err)
	}

	for {
		transaction, err := core.GetTransactions(os.Getenv("API_URL"))
		if err != nil {
			log.Fatalf("Error during request: %v", err)
		}

		core.Processor(transaction.Result)
		if err != nil {
			log.Fatalf("Error during processing: %v", err)
		}

		time.Sleep(5 * time.Second)
	}

}
