package feedfollows

import (
	"encoding/json"
	"fmt"
	"github.com/mike-kimani/fechronizo/v2/pkg/httpresponses"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mike-kimani/fechronizo/v2/internal/database"
	"github.com/mike-kimani/fechronizo/v2/internal/models"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow %v", err))
	}

	httpresponses.RespondWithJson(w, 200, models.DatabaseFeedFollowtoFeedFollow(feedFollow))

}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		UserID uuid.UUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), params.UserID)

	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows %v", err))
		return
	}

	httpresponses.RespondWithJson(w, 200, models.DatabaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_follow_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("error parsing json %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     params.FeedID,
		UserID: user.ID,
	})
	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("couldn't delete feed follow %v", err))
		return
	}
	httpresponses.RespondWithJson(w, 200, "Successfully deleted feed follow")
}

func (apiCfg *ApiConfig) HandlerDeleteFeedFollows2(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("couldn't parse the feed ID %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		httpresponses.RespondWithError(w, 400, fmt.Sprintf("couldn't delete feed follow %v", err))
		return
	}

	httpresponses.RespondWithJson(w, 200, struct{}{})
}
