package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

type RecommendationRequest struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}

type RecommendationResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func main() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	TOKEN := os.Getenv("OPEN_AI_TOKEN")
	client := openai.NewClient(TOKEN)

	e := echo.New()

	e.POST("/recommendation", func(c echo.Context) error {
		req := new(RecommendationRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		// Generate the recommendation using OpenAI GPT-3
		recommendation, err := generateRecommendation(client, req.Budget, req.Purpose)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error generating recommendation")
		}

		response := RecommendationResponse{
			Status: "success",
			Data:   recommendation,
		}

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func generateRecommendation(client *openai.Client, budget int, purpose string) (string, error) {
	// You can customize the user message based on the input data
	userMessage := fmt.Sprintf("I have a budget of Rp %d and I need a %s laptop.", budget, purpose)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant that provides laptop recommendations.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userMessage,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
