package readiness

import (
	"github.com/mike-kimani/fechronizo/pkg/httpresponses"
	"net/http"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	httpresponses.RespondWithJson(w, 200, struct{}{})
}
