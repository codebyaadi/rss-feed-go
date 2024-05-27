package models

import (
	"time"

	"github.com/codebyaadi/rss-agg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID			uuid.UUID	`json:"id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	Name		string		`json:"name"`
	APIkey		string		`json:"api_key"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIkey: dbUser.ApiKey,
	}
}