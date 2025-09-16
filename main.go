package main

import (
	"fmt"
	"github.com/chhandakguria/fraud_detector/ai"
	"log"
	"net/http"

	"github.com/chhandakguria/fraud_detector/db"
	"github.com/chhandakguria/fraud_detector/router"
)

func main() {
	// Connect DB
	db.Connect()
	// Initialize AI client
	ai.InitAI()

	// Setup router
	r := router.SetupRouter()

	fmt.Println("Fraud Detector running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
