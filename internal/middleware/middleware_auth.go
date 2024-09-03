package middleware

import (
	"fmt"
	"github.com/mike-kimani/fechronizo/v2/pkg/httpresponses"
	"net/http"

	"github.com/mike-kimani/fechronizo/v2/internal/database"
	"github.com/mike-kimani/fechronizo/v2/pkg/auth"
)

type ApiConfig struct {
	DB *database.Queries
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *ApiConfig) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			httpresponses.RespondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		//r.Context() allows using the current context of the request
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			httpresponses.RespondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
