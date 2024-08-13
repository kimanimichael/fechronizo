package readiness

import "net/http"
import "github.com/mike-kimani/fechronizo/pkg/jsonresponses"

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	jsonresponses.RespondWithJson(w, 200, struct{}{})
}
