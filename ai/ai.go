package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"fraud-detector/models"
)

var client *openai.Client

func InitAI() {
	client = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

func ScoreTransaction(tx models.Transaction) (float64, string, error) {
	ctx := context.Background()

	prompt := fmt.Sprintf("Detect fraud risk. User: %s, Points: %d, Device: %s",
		tx.UserID, tx.Points, tx.DeviceID)

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4o-mini",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a fraud detection model. Respond with JSON {score:0-1, reason:string}"},
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return 0, "", err
	}

	return 0.85, resp.Choices[0].Message.Content, nil // mock score + AI reason
}
