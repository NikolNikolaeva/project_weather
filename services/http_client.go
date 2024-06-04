package services

import (
	"net/http"
	"time"
)

func NewHTTPClient(timeout ...time.Duration) *http.Client {
	return &http.Client{
		Timeout: append(timeout, 10*time.Second)[0],
	}
}
