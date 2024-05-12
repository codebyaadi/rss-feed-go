package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    if code != http.StatusNoContent {
        if err := json.NewEncoder(w).Encode(payload); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func ResponseWithError(w http.ResponseWriter, code int, message string) {
    ResponseWithJSON(w, code, map[string]string{"error": message})
}