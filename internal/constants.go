package internal

import (
	"path/filepath"

	pongo "gopkg.in/flosch/pongo2.v3"
)

type paths struct {
	Templates  string
	CSS        string
	Javascript string
}

// Bunch of variables which are populated to be used with the program.
var (
	publicPage        string
	protectedPage     string
	loginPage         string
	SuccessUploadPage string
	path              paths
	templates         *TemplateList
)

// TemplateList is a collection of pongo templates which is used in the application.
type TemplateList struct {
	Login     *pongo.Template
	Protected *pongo.Template
	Public    *pongo.Template
}

//SetAssetsPath defines a baseDir for static assets.
func SetAssetsPath(baseDir string) *TemplateList {
	path = paths{
		Templates:  "templates/",
		CSS:        "assets/css/",
		Javascript: "assets/js",
	}
	loginPage = filepath.Join(baseDir, path.Templates, "login.html")
	publicPage = filepath.Join(baseDir, path.Templates, "public.html")
	protectedPage = filepath.Join(baseDir, path.Templates, "protected.html")

	return &TemplateList{
		Login:     pongo.Must(pongo.FromFile(loginPage)),
		Public:    pongo.Must(pongo.FromFile(publicPage)),
		Protected: pongo.Must(pongo.FromFile(protectedPage)),
	}
}
