package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mike-kimani/fechronizo/internal/models"
	"github.com/mike-kimani/fechronizo/pkg/jsonresponses"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mike-kimani/fechronizo/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type jsonNullInt32 struct {
	sql.NullInt32
}

func (v jsonNullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	} else {
		return json.Marshal(v.Int32)
	}

}

func (v *jsonNullInt32) UnmarshalJSON(data []byte) error {
	var x *int32
	err := json.Unmarshal(data, &x)

	if err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int32 = *x
	} else {
		v.Valid = false
	}
	return nil
}

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name          string        `json:"name"`
		ChickenBought sql.NullInt32 `json:"chicken_bought"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		Name:          params.Name,
		ChickenBought: params.ChickenBought,
	})
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	jsonresponses.RespondWithJson(w, 201, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	jsonresponses.RespondWithJson(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		jsonresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	jsonresponses.RespondWithJson(w, 200, models.DatabasePostsToPosts(posts))
}
