package routes

import (
	"encoding/json"
	"net/http"

	"github.com/farofadev/headlinesbr/data"
	"github.com/farofadev/headlinesbr/payloads"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts := data.FetchHeadlines()

	payload := payloads.NewPayload(
		payloads.WithAddMeta("total", len(*posts)),
		payloads.WithData(posts),
	)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(payload)
}
