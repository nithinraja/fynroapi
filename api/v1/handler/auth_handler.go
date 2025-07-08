package handler

import (
	"encoding/json"
	"net/http"

	"ai-financial-api/internal/auth"
	"ai-financial-api/utils"
)

func RequestOTP(w http.ResponseWriter, r *http.Request) {
    type request struct {
        Mobile string `json:"mobile"`
    }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    err := auth.SendOTP(req.Mobile)
    if err != nil {
        utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to send OTP")
        return
    }

    utils.SuccessResponse(w, "OTP sent successfully")
}

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
    type request struct {
        Mobile string `json:"mobile"`
        OTP    string `json:"otp"`
    }
    var req request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    token, err := auth.VerifyOTP(req.Mobile, req.OTP)
    if err != nil {
        utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
        return
    }

    utils.JSONResponse(w, http.StatusOK, map[string]string{"token": token})
}
