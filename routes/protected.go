package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// GetProtected handles GET requests on / end point
func (idx Resource) GetProtected(w http.ResponseWriter, r *http.Request) {
	idx.logger.Info("Hitting the index page")
	idx.templateList.Protected.ExecuteWriter(nil, w)
}

// RouteProtected returns a chi.Mux for all the requests defined on `/`
func (idx Resource) RouteProtected() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetProtected)
	return router
}
