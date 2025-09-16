package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chhandakguria/fraud_detector/ai"
	"github.com/chhandakguria/fraud_detector/db"
	"github.com/chhandakguria/fraud_detector/models"
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

	// Get AI-based fraud score
	score, reason := ai.ScoreTransaction(tx)

	var resp FraudResponse
	if score > 0.7 {
		resp = FraudResponse{TransactionID: tx.ID, Status: "fraud_suspected", Reason: reason, Score: score}
	} else {
		resp = FraudResponse{TransactionID: tx.ID, Status: "ok", Reason: reason, Score: score}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
