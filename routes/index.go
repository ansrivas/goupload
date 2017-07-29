package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// GetIndex handles GET requests on / end point
func (idx Resource) GetIndex(w http.ResponseWriter, r *http.Request) {
	idx.logger.Info("Hitting the index page")

	sess, _ := idx.store.GetSession(r)

	idx.logger.Infof("Returning user is %v", sess.Values["user"])
	err := idx.templateList.Public.ExecuteWriter(nil, w)
	if err != nil {
		idx.logger.Info("Failed to serve the index page")
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

// PostIndex handles POST requests on / end point
func (idx Resource) PostIndex(w http.ResponseWriter, r *http.Request) {
	redirectTarget := "/"
	emailUser := r.FormValue("emailuser")
	password := r.FormValue("passwordfield")

	// If the session is there, direct to upload page
	// sess, err := session.GetSessionStore(w, r)
	// if err != nil {
	// 	HandleSessionError(w, err)
	// 	return
	// }
	idx.logger.Info("Post on this resource.")
	idx.logger.Infoln(emailUser, password)
	if (emailUser == "admin@gmail.com") && (password == "admin") {

		idx.logger.Info("Successfully logged in there , now what?")
		sess, _ := idx.store.GetSession(r)
		sess.Values["user"] = emailUser
		if err := sess.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		redirectTarget = "/protected"
		http.Redirect(w, r, redirectTarget, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, redirectTarget, http.StatusFound)
}

// RouteIndex returns a chi.Mux for all the requests defined on `/`
func (idx Resource) RouteIndex() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", idx.GetIndex)
	router.Post("/", idx.PostIndex)
	return router
}
