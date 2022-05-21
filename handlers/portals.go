package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/farofadev/headlinesbr/data"
	"github.com/farofadev/headlinesbr/payloads"
	"github.com/julienschmidt/httprouter"
)

func PortalsIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	payload := payloads.NewPayload(
		payloads.WithData(&data.Portals),
	)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(payload)
}
