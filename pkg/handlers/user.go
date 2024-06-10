package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/codebyaadi/rss-agg/internal/auth"
	"github.com/codebyaadi/rss-agg/internal/database"
	"github.com/codebyaadi/rss-agg/internal/models"
	"github.com/codebyaadi/rss-agg/pkg/helpers"
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

	helpers.ResponseWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (cfg *ApiConfig) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		helpers.ResponseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := cfg.DB.GetUserByApikey(r.Context(), apiKey)
	if err != nil {
		helpers.ResponseWithError(w, 403, fmt.Sprintf("Couldn't get the user: %v", err))
		return
	}

	helpers.ResponseWithJSON(w, 200, models.DatabaseUserToUser(user))
}