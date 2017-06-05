package routes

import (
	"net/http"

	"github.com/pressly/chi"

	pongo "gopkg.in/flosch/pongo2.v3"
)

//GetIndex handles GET requests on / end point
func (idx Resource) GetIndex(w http.ResponseWriter, r *http.Request) {
	idx.logger.Info("Hitting the index page")
	idx.templateList.Public.ExecuteWriter(pongo.Context{}, w)
	w.Write([]byte("welcome"))

}

//PostIndex handles POST requests on / end point
func (idx Resource) PostIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

//Routes returns a chi.Mux for all the requests defined on /
func (idx Resource) RouteIndex() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetIndex)
	router.Post("/", idx.PostIndex)
	return router
}
