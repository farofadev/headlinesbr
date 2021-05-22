package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/farofadev/headlinesbr/data"
	"github.com/farofadev/headlinesbr/payloads"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	findOptions := options.Find()

	page, perPage := data.ParsePageAndPerPage(
		r.URL.Query().Get("page"),
		r.URL.Query().Get("per_page"),
	)

	if perPage > 500 {
		perPage = 500
	}

	filters := bson.M{}
	portalId := int64(0)

	if sPortalId := r.URL.Query().Get("portal_id"); sPortalId != "" {
		portalId, _ = strconv.ParseInt(sPortalId, 10, 64)

		filters["portal_id"] = portalId
	}

	findOptions.SetSort(bson.M{"created_at": -1})
	findOptions.SetSkip(data.PageOffset(page, perPage))
	findOptions.SetLimit(perPage)

	posts := data.FetchHeadlines(filters, findOptions)

	payload := payloads.NewPayload(
		payloads.AddMeta("page", page),
		payloads.AddMeta("per_page", perPage),
		payloads.AddExtraWhen(
			portalId > 0,
			"portal",
			func() interface{} { return data.FindPortalById(&data.Portals, uint(portalId)) },
		),
		payloads.WithData(posts),
	)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(payload)
}
