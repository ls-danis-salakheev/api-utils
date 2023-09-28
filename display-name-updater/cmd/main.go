package main

import (
	"display-name-updater/internal/csv"
	"display-name-updater/internal/models"
	"display-name-updater/internal/rest"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	loadEnvVars()
	lines, errOccurred := csv.LoadCsvLines(os.Getenv("CLIENTS_FILE_PATH"))
	if errOccurred {
		return
	}
	clientData := models.CreateClientArr(lines, 1, len(lines))
	rest.Update(clientData)

	exit()
}

func exit() {
	fmt.Println("Finishing after 30 seconds...")
	select {
	case <-time.After(30 * time.Second):
		fmt.Println("Finishing timeout is over. Exiting...")
	}
}

func loadEnvVars() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Cannot load env vars")
		panic(envErr)
	}
}
