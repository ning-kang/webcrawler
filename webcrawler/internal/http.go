package internal

import (
	"log"
	"net/http"
	"time"
)

var client http.Client

func get(url string) *http.Response {
	client.Timeout = 30 * time.Second

	res, err := client.Get(url)
	if err != nil {
		log.Fatal("Failed to get server response:", err)
	}

	return res
}
