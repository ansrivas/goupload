package routes

import (
	"net/http"

	"github.com/ansrivas/goupload/environment"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/spf13/viper"
)

//BaseRoutes is the entry point of routes.
func BaseRoutes(env *environment.Env, conf *viper.Viper) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
				next.ServeHTTP(w, r)
			})
		},
		middleware.Recoverer,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.StripSlashes,
	)

	staticResource := StaticResource{staticDir: conf.GetString("app.static_dir_path")}
	router.Mount("/static", staticResource.Routes())

	router.Mount("/", IndexResource{}.Routes())

	// router.Mount("/upload", UploadResource{}.Routes())
	return router
}
