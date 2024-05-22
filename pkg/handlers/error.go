package handlers

import (
	"net/http"

	"github.com/codebyaadi/rss-agg/pkg/helpers"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	helpers.ResponseWithError(w, 400, "Something went wrong")
}