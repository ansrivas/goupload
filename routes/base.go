package routes

import (
	"net/http"

	"github.com/pressly/chi"
)

//IndexRouter handles requests on /
type IndexRouter struct {
}

//GetIndex handles GET requests on / end point
func (idx IndexRouter) GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

//Routes returns a chi.Mux for all the requests defined on /
func (idx IndexRouter) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetIndex)
	// http.ListenAndServe(":3000", r)
	return router
}
