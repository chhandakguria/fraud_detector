package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"fraud-detector/ai"
	"fraud-detector/db"
	"fraud-detector/models"
)

type FraudResponse struct {
	TransactionID uint    `json:"transaction_id"`
	Status        string  `json:"status"`
	Reason        string  `json:"reason"`
	Score         float64 `json:"score"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var tx models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tx.Timestamp = time.Now()
	db.DB.Create(&tx)

	score, reason, err := ai.ScoreTransaction(tx)
	if err != nil {
		http.Error(w, "AI service failed", http.StatusInternalServerError)
		return
	}

	status := "ok"
	if score > 0.7 {
		status = "fraud_suspected"
	}

	resp := FraudResponse{
		TransactionID: tx.ID,
		Status:        status,
		Reason:        reason,
		Score:         score,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
