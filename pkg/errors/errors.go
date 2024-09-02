package errors

import (
	"github.com/mike-kimani/fechronizo/pkg/httpresponses"
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	httpresponses.RespondWithError(w, 400, "Something went wrong")
}
