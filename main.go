package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Définissez la variable environnementale OPENAI_API_KEY puis relancez !")
		return
	}

	client := openai.NewClient(apiKey)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("GPTK prêt – posez vos questions (Ctrl+C pour quitter).")

	for scanner.Scan() {
		prompt := scanner.Text()
		if prompt == "" {
			continue
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4o, // ou "gpt-4o" si erreur
				Messages: []openai.ChatCompletionMessage{
					{Role: "user", Content: prompt},
				},
			},
		)

		if err != nil {
			fmt.Println("Erreur API :", err)
			continue
		}

		fmt.Println("\n--- Réponse GPTK ---")
		fmt.Println(resp.Choices[0].Message.Content)
		fmt.Print("\n>>> ")
	}
}
