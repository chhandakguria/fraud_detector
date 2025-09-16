package main

import (
	"fmt"
	"log"
	"net/http"

	"fraud-detector/db"
	"fraud-detector/router"
)

func main() {
	// Connect DB
	db.Connect()

	// Setup router
	r := router.SetupRouter()

	fmt.Println("Fraud Detector running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
