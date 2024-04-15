package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mike-kimani/rssagg/internal/database"
)

type jsonNullInt32 struct{
	sql.NullInt32
}

func(v jsonNullInt32) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int32)
	} else {
		return json.Marshal(v.Int32)
	}

}

func(v *jsonNullInt32) UnmarshalJSON(data []byte) error {
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

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct{
		Name string `json:"name"`
		ChickenBought sql.NullInt32 `json:"chicken_bought"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err:= apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		ChickenBought: params.ChickenBought,
	} )
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts: %v", err))
		return
	}

	respondWithJson(w, 200, databasePostsToPosts(posts))
}