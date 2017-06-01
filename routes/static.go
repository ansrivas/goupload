package routes

import (
	"net/http"

	"github.com/pressly/chi"
)

//StaticResource provides end point to serve static files.
type StaticResource struct {
	staticDir string
}

//Routes returns a chi.Mux for serving static files.
func (st StaticResource) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.FileServer("/", http.Dir(st.staticDir))
	return router
}
