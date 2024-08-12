package errors

import (
	"github.com/mike-kimani/fechronizo/pkg/jsonresponses"
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	jsonresponses.RespondWithError(w, 400, "Something went wrong")
}
