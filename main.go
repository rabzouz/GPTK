package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	// Lire la clé depuis l'environnement
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Définissez OPENAI_API_KEY puis relancez !")
		return
	}

	// Créer le client avec la clé
	client := openai.NewClient(option.WithAPIKey(apiKey))

	// Boucle interactive
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("GPTK prêt – posez vos questions (Ctrl+C pour quitter).")

	for scanner.Scan() {
		prompt := scanner.Text()
		if prompt == "" {
			continue
		}

		// Appel à l'API (modèle gpt-4o)
		resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
			},
			Model: openai.String("gpt-4o"),
		})
		if err != nil {
			fmt.Println("Erreur API :", err)
			continue
		}
		fmt.Println("\n--- Réponse GPTK ---")
		fmt.Println(resp.Choices.Message.Content)
		fmt.Print("\n>>> ")
	}
}
