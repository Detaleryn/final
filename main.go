package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Detaleryn/final/pkg/db"
	"github.com/Detaleryn/final/pkg/server"
)

func main() {

	dbFile := "scheduler.db"
	err := db.Init(dbFile)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer db.Close()

	server.Init()
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
