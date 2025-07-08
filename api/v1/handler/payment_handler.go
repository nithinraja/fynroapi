package handler

import (
	"encoding/json"
	"net/http"

	"ai-financial-api/internal/payment"
	"ai-financial-api/utils"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    type request struct {
        QuestionUUID string  `json:"question_uuid"`
        Amount       float64 `json:"amount"`
    }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.QuestionUUID == "" || req.Amount <= 0 {
        utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    paymentLink, err := payment.Create(req.QuestionUUID, req.Amount)
    if err != nil {
        utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]string{"payment_link": paymentLink})
}
