package handlers

import (
	"encoding/json"
	"net/http"
)

func GetFeed(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]string{
		{
			"title": "First post",
			"url":   "https://example.com",
		},
	})
}