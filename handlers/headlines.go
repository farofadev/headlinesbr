package handlers

import (
	"net/http"
	"strconv"

	"github.com/goccy/go-json"

	"github.com/farofadev/headlinesbr/payloads"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/farofadev/headlinesbr/model"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	findOptions := options.Find()

	page, perPage := model.ParsePageAndPerPage(
		r.URL.Query().Get("page"),
		r.URL.Query().Get("per_page"),
	)

	if perPage > 500 {
		perPage = 500
	}

	filters := bson.M{}
	var portalId int64

	if str := r.URL.Query().Get("portal_id"); str != "" {
		portalId, _ = strconv.ParseInt(str, 10, 64)

		filters["portal_id"] = portalId
	}

	findOptions.SetSort(bson.M{"created_at": -1})
	findOptions.SetSkip(model.PageOffset(page, perPage))
	findOptions.SetLimit(perPage)

	posts := model.FetchHeadlines(filters, findOptions)

	payload := payloads.NewPayload(
		payloads.AddMeta("page", page),
		payloads.AddMeta("per_page", perPage),
		payloads.AddExtraWhen(
			portalId > 0,
			"portal",
			func() interface{} { return model.FindPortalById(&model.Portals, uint(portalId)) },
		),
		payloads.WithData(posts),
	)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(payload)
}

func HeadlinesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	model.ScrapeAndStoreHeadlines(model.Portals)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func PortalsIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	payload := payloads.NewPayload(
		payloads.WithData(&model.Portals),
	)

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	enc.Encode(payload)
}
