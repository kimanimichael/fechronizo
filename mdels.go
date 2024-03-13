package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mike-kimani/rssagg/internal/database"

	"database/sql"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ChickenBought sql.NullInt32 `json:"chicken_bought"`
	APIKey string `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User{
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ChickenBought: dbUser.ChickenBought,
		APIKey: dbUser.ApiKey,
	}
}