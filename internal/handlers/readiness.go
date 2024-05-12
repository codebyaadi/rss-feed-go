package handlers

import (
	"net/http"

	"github.com/codebyaadi/rss-agg/internal/helpers"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseWithJSON(w, 200, struct{}{})
}