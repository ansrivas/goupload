package environment

import "github.com/ansrivas/goupload/internal"

//Env provides base for all the GLOBAL services.
type Env struct {
	TemplateList *internal.TemplateList
	// store sessions.Store
	// db db
	// templates map[string]*template.Template
}
