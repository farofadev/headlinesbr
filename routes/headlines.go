package routes

import (
	"net/http"

	"github.com/farofadev/headlinesbr/data"
	"github.com/julienschmidt/httprouter"
)

func HeadlinesHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data.ScrapeAndStoreHeadlines(data.Portals)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
