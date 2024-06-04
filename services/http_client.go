package services

import (
	"net/http"
	"time"
)

func NewHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return client
}
