package openai

import (
	"fmt"
	"strings"
)

type PromptBuilder struct{}

// NewPromptBuilder creates a new instance of PromptBuilder
func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{}
}

// BuildInitialPrompt generates the initial OpenAI prompt based on user's question and name
func (pb *PromptBuilder) BuildInitialPrompt(name, question string) string {
	return fmt.Sprintf(`
Hello, my name is %s. I have a financial question: "%s".
Please help me by providing a clear and concise suggestion or response tailored to an individual in India.
`, name, question)
}

// BuildAnalysisPrompt generates a detailed financial analysis prompt for the second OpenAI call
func (pb *PromptBuilder) BuildAnalysisPrompt(name, question string) string {
	return fmt.Sprintf(`
User Name: %s
Question: %s

Now that the user is verified, provide a detailed analysis and breakdown of their financial situation.
Include insights, risks, and opportunities. Format the response as a list of actionable recommendations.
`, name, question)
}

// BuildSuggestionPrompt generates personalized suggestions after login
func (pb *PromptBuilder) BuildSuggestionPrompt(userName string, pastQuestions []string) string {
	questions := strings.Join(pastQuestions, "\n- ")
	return fmt.Sprintf(`
User: %s

Below are the questions previously asked:
- %s

Based on this, provide personalized financial improvement suggestions.
Include specific actions for budgeting, investing, and saving money.
`, userName, questions)
}
