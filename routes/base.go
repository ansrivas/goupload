package routes

import (
	"net/http"

	"github.com/ansrivas/goupload/internal"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Resource handles requests on /
type Resource struct {
	templateList *internal.TemplateList
	logger       *logrus.Entry
	// store sessions.Store
	// db db
	// templates map[string]*template.Template
}

// NewRouter is the entry point of routes.
func NewRouter(templateList *internal.TemplateList, log *logrus.Entry, conf *viper.Viper) *chi.Mux {
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

	staticFilesPath := conf.GetString("app.static_dir_path")

	logrus.Infoln("Path to static files", staticFilesPath)

	resources := Resource{
		templateList: templateList,
		logger:       log,
	}
	router.FileServer("/static", http.Dir(staticFilesPath))
	router.Mount("/", resources.RouteIndex())
	router.Mount("/protected", resources.RouteProtected())

	// router.Mount("/upload", UploadResource{}.Routes())
	return router
}
