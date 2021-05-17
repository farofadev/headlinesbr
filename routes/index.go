package routes

import (
	"encoding/json"
	"net/http"

	"github.com/farofadev/headlinesbr/data"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts := data.FetchHeadlines()

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(posts)
}