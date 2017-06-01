package routes

import (
	"net/http"

	"github.com/pressly/chi"
)

//IndexResource handles requests on /
type IndexResource struct {
}

//GetIndex handles GET requests on / end point
func (idx IndexResource) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

//PostIndex handles POST requests on / end point
func (idx IndexResource) PostIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

//Routes returns a chi.Mux for all the requests defined on /
func (idx IndexResource) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetIndex)
	router.Post("/", idx.PostIndex)
	return router
}
