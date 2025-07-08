package router

import (
	"ai-financial-api/api/v1/handler"
	"ai-financial-api/api/v1/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Public routes (no JWT middleware)
    public := r.PathPrefix("/api/v1").Subrouter()
    public.HandleFunc("/auth/request-otp", handler.RequestOTP).Methods("POST")
    public.HandleFunc("/auth/verify-otp", handler.VerifyOTP).Methods("POST")

    // Authenticated routes (JWT protected)
    private := r.PathPrefix("/api/v1").Subrouter()
    private.Use(middleware.JWTMiddleware)
    private.HandleFunc("/questions/ask", handler.AskQuestion).Methods("POST")
    private.HandleFunc("/payments/create", handler.CreatePayment).Methods("POST")

    return r
}
