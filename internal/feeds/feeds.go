package feeds

import (
	"encoding/json"
	"fmt"
	"github.com/mike-kimani/fechronizo/internal/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mike-kimani/fechronizo/internal/database"
	"github.com/mike-kimani/fechronizo/pkg/jsonresponses"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeeds(r.Context(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
		UserName:  user.Name,
	})
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	jsonresponses.RespondWithJson(w, 201, models.DatabaseFeedToFeed(feed))
}

func (apiCfg *ApiConfig) HandlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feeds %v", err))
	}
	jsonresponses.RespondWithJson(w, 200, models.DatabaseFeedstoFeeds(feeds))
}
