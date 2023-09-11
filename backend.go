package main

import (
	"backend/app/route"
	"backend/app/shared/database"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/joho/godotenv"
)

// *****************************************************************************
// Application Logic
// *****************************************************************************

func init() {
	// Verbose logging with file name and line number
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func configureEnvironments() {
	os.Clearenv()

	err := godotenv.Load("api.env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {

	//load local env vars
	if os.Getenv("GO_ENV") != "production" {
		configureEnvironments()
	}

	// Connect to databases
	database.ConnectPostgresSQL()

	port := os.Getenv("PORT")
	// println(port)

	if len(port) == 0 {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe(":"+port, route.Load()))
}
