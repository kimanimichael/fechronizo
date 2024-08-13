package errors

import (
	"github.com/mike-kimani/fechronizo/pkg/jsonresponses"
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	jsonresponses.RespondWithError(w, 400, "Something went wrong")
}
