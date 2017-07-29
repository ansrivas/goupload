package routes

import (
	"net/http"

	"github.com/ansrivas/goupload/internal"
	"github.com/go-chi/chi"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Resource handles requests on /
type Resource struct {
	templateList *internal.TemplateList
	logger       *logrus.Entry
	store        *Store
	// db db
	// templates map[string]*template.Template
}

// NewRouter is the entry point of routes.
func NewRouter(templateList *internal.TemplateList, log *logrus.Entry, conf *viper.Viper) (*chi.Mux, error) {
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

	poolsize := conf.GetInt("redis.poolsize")
	network := conf.GetString("redis.network")
	address := conf.GetString("redis.address")
	password := conf.GetString("redis.password")
	hashkey := conf.GetString("redis.hashkey")
	blockkey := conf.GetString("redis.blockkey")
	store, err := NewDefaultStore(poolsize, network, address, password, []byte(hashkey), []byte(blockkey))
	if err != nil {
		return nil, err
	}
	// Set default for 60s
	store.Redis.SetMaxAge(60)

	resources := Resource{
		templateList: templateList,
		logger:       log,
		store:        store,
	}
	FileServer(router, "/static", http.Dir(staticFilesPath))
	router.Mount("/", resources.RouteIndex())
	router.Mount("/protected", resources.RouteProtected())

	// router.Mount("/upload", UploadResource{}.Routes())
	return router, nil
}
