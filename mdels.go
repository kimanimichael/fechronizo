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


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	URL string `json:"url"`
	UserID uuid.UUID `json:"user_id"`
	UserName string `json:"user_name"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		ID: dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
		UserName: dbFeed.UserName,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

