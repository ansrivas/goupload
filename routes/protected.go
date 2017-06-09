package routes

import (
	"net/http"

	"github.com/pressly/chi"

	pongo "gopkg.in/flosch/pongo2.v3"
)

// GetIndex handles GET requests on / end point
func (idx Resource) GetProtected(w http.ResponseWriter, r *http.Request) {
	idx.logger.Info("Hitting the index page")
	idx.templateList.Protected.ExecuteWriter(pongo.Context{}, w)
}

// RouteIndex returns a chi.Mux for all the requests defined on `/`
func (idx Resource) RouteProtected() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetProtected)
	return router
}
