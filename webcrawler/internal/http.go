package internal

import (
	"log"
	"net/http"
)

func get(url string) *http.Response {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to get server response:", err)
	}

	return res
}
