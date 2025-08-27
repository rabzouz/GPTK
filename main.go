package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// Type pour une entrée d'historique
type Entry struct {
	Prompt    string `json:"prompt"`
	Response  string `json:"response"`
	Timestamp string `json:"ts"` // yyyy-mm-ddThh:mm:ssZ
}

// historyPath retourne le chemin complet du fichier history.json
func historyPath() string {
	exe, _ := os.Executable() // Dossier de l’exécutable
	dir := filepath.Dir(exe)
	return filepath.Join(dir, "history.json")
}

// loadHistory lit le fichier s’il existe.
func loadHistory() (hist []Entry) {
	data, err := os.ReadFile(historyPath())
	if err != nil {
		return []Entry{}
	}
	_ = json.Unmarshal(data, &hist)
	return
}

// saveHistory remplace le fichier par la version mise à jour.
func saveHistory(hist []Entry) {
	data, _ := json.MarshalIndent(hist, "", "  ")
	_ = os.WriteFile(historyPath(), data, 0o644)
}

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Définissez la variable environnementale OPENAI_API_KEY puis relancez !")
		return
	}

	client := openai.NewClient(apiKey)

	// Flags : --show-history et --clear-history
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--show-history":
			for i, e := range loadHistory() {
				fmt.Printf("%d) [%s] %s ⇒ %s\n", i+1, e.Timestamp, e.Prompt, e.Response)
			}
			return
		case "--clear-history":
			os.Remove(historyPath())
			fmt.Println("Historique effacé.")
			return
		}
	}

	// Charger l'historique existant
	history := loadHistory()
	fmt.Printf("Historique chargé : %d entrées\n", len(history)) // Affichage pour debug

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
				Model: openai.GPT4o, // ou "gpt-4o"
				Messages: []openai.ChatCompletionMessage{
					{Role: "system", Content: "Réponds toujours en français, de manière claire et concise."}, // Force les réponses en français
					{Role: "user", Content: prompt},
				},
			},
		)

		if err != nil {
			fmt.Println("Erreur API :", err)
			continue
		}

		response := resp.Choices[0].Message.Content
		fmt.Println("\n--- Réponse GPTK ---")
		fmt.Println(response)
		fmt.Print("\n>>> ")

		// Ajouter à l'historique et sauvegarder
		history = append(history, Entry{
			Prompt:    prompt,
			Response:  response,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
		saveHistory(history)
	}
}
