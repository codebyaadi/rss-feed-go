package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/codebyaadi/rss-agg/internal/database"
	"github.com/codebyaadi/rss-agg/internal/helpers"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Error parsing Json: %s", err))
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		helpers.ResponseWithError(w, 400, fmt.Sprintf("Can't create user %s: ", err))
		return
	}

	helpers.ResponseWithJSON(w, 200, user)
}