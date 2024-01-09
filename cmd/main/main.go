package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pliavi/go-for-tickets/pkg/routes"
	"github.com/pliavi/go-for-tickets/pkg/routines"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	go routines.QueueProcess()

	router := routes.NewRouter()
	port := 8082 // take from env variable
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
