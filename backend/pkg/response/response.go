package response

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type APIError struct {
	Field  string `json:"field,omitempty"`
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

type Meta struct {
	RequestID  string      `json:"request_id"`
	Timestamp  time.Time   `json:"timestamp"`
	Pagination interface{} `json:"pagination,omitempty"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []APIError  `json:"errors,omitempty"`
	Meta    Meta        `json:"meta"`
}

func NewMeta() Meta {
	return Meta{
		RequestID: uuid.New().String(),
		Timestamp: time.Now().UTC(),
	}
}

func Success(c *gin.Context, status int, message string, data interface{}, meta ...Meta) {
	m := NewMeta()
	if len(meta) > 0 {
		m = meta[0]
	}
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    m,
	})
}

func Error(c *gin.Context, status int, message string, detail interface{}, field ...string) {
	m := NewMeta()
	apiErr := APIError{
		Code:   http.StatusText(status),
		Detail: fmt.Sprintf("%v", detail),
	}
	if len(field) > 0 {
		apiErr.Field = field[0]
	}

	c.JSON(status, APIResponse{
		Success: false,
		Message: message,
		Errors:  []APIError{apiErr},
		Meta:    m,
	})
}
