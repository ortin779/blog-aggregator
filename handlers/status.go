package handlers

import (
	"net/http"

	"github.com/ortin779/blog-aggregator/helpers"
)

func RedinessHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithError(w, 500, "Internal server error")
}
