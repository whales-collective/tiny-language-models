package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// go test -v -run TestHawaiianPizzaExpert
func TestStarTrekExpert(t *testing.T) {
	ctx := context.Background()

	llmURL := os.Getenv("MODEL_RUNNER_BASE_URL") + "/engines/llama.cpp/v1/"
	model := os.Getenv("MODEL_RUNNER_LLM_CHAT")

	client := openai.NewClient(
		option.WithBaseURL(llmURL),
		option.WithAPIKey(""),
	)

	fmt.Println("🐳🤖 ENDPOINT:", llmURL)

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(`
			You are a Star Trek expert. Your name is Seven of Nine.
			USE ONLY THE INFORMATION PROVIDED IN THE KNOWLEDGE BASE.		
		`),
		openai.SystemMessage(`
			KNOWLEDGE BASE: 
			Star Trek is a science fiction media franchise that includes television series, films, books, and more.
			James T. Kirk is a fictional character in the Star Trek franchise, known for being the captain of the USS Enterprise.
			USS Enterprise is a starship in the Star Trek universe, known for its missions in space exploration.
			Spock is a fictional character in the Star Trek franchise, known for his Vulcan heritage and logical thinking.
			Leonard McCoy, also known as "Bones," is a fictional character in the Star Trek franchise, serving as the chief medical officer of the USS Enterprise.
			The best friend of James T. Kirk is Spock, who is known for his logical thinking and Vulcan heritage.
		`),
		openai.UserMessage("Who is James T. Kirk?"),
	}

	param := openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       model,
		Temperature: openai.Opt(0.8),
	}

	completion, err := client.Chat.Completions.New(ctx, param)

	if err != nil {
		t.Fatalf("😡 chat completion failed: %s", err)
	}

	fmt.Println("🐳🤖", completion.Choices[0].Message.Content)

	messages = append(messages,
		openai.AssistantMessage(completion.Choices[0].Message.Content),
		openai.UserMessage("Who is his best friend?"),
	)

	param = openai.ChatCompletionNewParams{
		Messages:    messages,
		Model:       model,
		Temperature: openai.Opt(0.8),
	}

	completion, err = client.Chat.Completions.New(ctx, param)

	if err != nil {
		t.Fatalf("😡 chat completion failed: %s", err)
	}

	fmt.Println("🐳🤖", completion.Choices[0].Message.Content)

	response := strings.ToLower(completion.Choices[0].Message.Content)

	expectedWords := []string{"spock"}
	for _, word := range expectedWords {
		if !strings.Contains(response, word) {
			t.Errorf("Expected response to contain word '%s', but it was not found", word)
		}
	}

	if len(completion.Choices[0].Message.Content) == 0 {
		t.Error("received empty response from model")
	}
}
