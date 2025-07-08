package openai

func BuildPrompt(question string) string {
    return "You are a financial advisor AI. Answer the following question:\n\n" + question
}
