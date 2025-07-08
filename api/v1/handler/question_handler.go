package handler

import (
	"encoding/json"
	"net/http"

	"ai-financial-api/internal/question"
	"ai-financial-api/utils"
)

func AskQuestion(w http.ResponseWriter, r *http.Request) {
    type request struct {
        QuestionText string `json:"question"`
    }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.QuestionText == "" {
        utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    answer, err := question.Ask(req.QuestionText, r.Context())
    if err != nil {
        utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]string{"answer": answer})
}
