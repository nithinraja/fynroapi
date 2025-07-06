// =================
// pkg/utils/response.go
// =================
package utils

import (
    "github.com/gin-gonic/gin"
)

type APIResponse struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
    c.JSON(statusCode, APIResponse{
        Success: true,
        Message: message,
        Data:    data,
    })
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
    response := APIResponse{
        Success: false,
        Message: message,
    }

    if err != nil {
        response.Error = err.Error()
    }

    c.JSON(statusCode, response)
}

