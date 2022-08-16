package errors

import (
	"time"
)

type ErrorBodyResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Message   string    `json:"message"`
	Errors    error     `json:"errors"`
}
