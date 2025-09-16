package ai

import (
	"context"
	"fmt"
	"github.com/chhandakguria/fraud_detector/models"
	"github.com/sashabaranov/go-openai"
	"log"
)

var client *openai.Client

func InitAI() {
	client = openai.NewClient("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")
}

func ScoreTransaction(tx models.Transaction) (float64, string) {
	ctx := context.Background()

	prompt := fmt.Sprintf("Detect fraud risk. User: %s, Points: %d, Device: %s",
		tx.UserID, tx.Points, tx.DeviceID)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a fraud detection AI."},
			{Role: "user", Content: prompt},
		},
	})

	if err != nil {
		log.Println("OpenAI API error:", err)
		return 0.5, "default"
	}
	answer := resp.Choices[0].Message.Content
	fmt.Println("AI response:", answer)

	// For demo: if points > 5000 → score high, else low
	if tx.Points > 5000 {
		return 0.85, answer
	}
	return 0.3, "Normal behavior per AI"

	//return 0.85, resp.Choices[0].Message.Content, nil // mock score + AI reason
}
