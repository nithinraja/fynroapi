package dto

type QuestionRequest struct {
	Question string `json:"question" binding:"required"`
}

type QuestionNameRequest struct {
	UUID string `json:"uuid" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type LinkQuestionsRequest struct {
	QuestionUUIDs []string `json:"question_uuids" binding:"required"`
}

// Request when a user submits a question (anonymous or logged-in)
type AskQuestionRequest struct {
	Content string `json:"content" binding:"required,min=5"` // Question text
}

// Response after storing the question
type AskQuestionResponse struct {
	QuestionID string `json:"question_id"`
	SessionID  string `json:"session_id"`
	Message    string `json:"message"`
}

// Request to provide name (after question is submitted)
type ProvideNameRequest struct {
	QuestionID string `json:"question_id" binding:"required"`
	Name       string `json:"name" binding:"required,min=2,max=100"`
}

// Response after calling OpenAI and saving the result
type AIResponse struct {
	AnswerID   string `json:"answer_id"`
	QuestionID string `json:"question_id"`
	Content    string `json:"content"`
	Type       string `json:"type"` // e.g., "initial", "recommendation"
}

// Request to get financial suggestions (post-login)
type RequestSuggestions struct {
	QuestionID string `json:"question_id" binding:"required"`
}

// Response containing list of recommendations
type RecommendationListResponse struct {
	Recommendations []string `json:"recommendations"`
}
