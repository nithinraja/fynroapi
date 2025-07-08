package question

import (
	"ai-financial-api/config"
	"ai-financial-api/internal/openai"
	"ai-financial-api/models"
	"context"
	"log"

	"github.com/google/uuid"
)

func Ask(questionText string, ctx context.Context) (string, error) {

	log.Print("Preparing to send prompt to OpenAI: %s", questionText)
    // Create question record
    q := models.Question{
        QuestionText: questionText,
        QuestionUUID: uuid.New().String(),
    }
    config.DB.Create(&q)

    // Build prompt and ask OpenAI
    prompt := openai.BuildPrompt(questionText)
	log.Printf("Built prompt: %s", prompt)

    answer, err := openai.Ask(prompt, ctx)
    if err != nil {
		log.Fatalf("OpenAI API call failed: %v", err)
        return "", err
    }

    // Save response
    resp := models.OpenAIResponse{
        QuestionID: q.ID,
        ResponseType: "initial",
        ResponseText: answer,
    }
    config.DB.Create(&resp)

    return answer, nil
}
