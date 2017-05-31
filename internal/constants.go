package internal

import (
	"path/filepath"
)

type paths struct {
	Templates  string
	CSS        string
	Javascript string
}

// Bunch of variables which are populated to be used with the program.
var (
	LandingPage       string
	UploadPage        string
	LoginPage         string
	SuccessUploadPage string
	path              paths
)

func init() {
	path = paths{
		Templates:  "static/templates/",
		CSS:        "static/assets/css/",
		Javascript: "static/assets/js",
	}
	LandingPage = filepath.Join(path.Templates, "landing.html")
	UploadPage = filepath.Join(path.Templates, "upload.html")
	LoginPage = filepath.Join(path.Templates, "login.html")
	SuccessUploadPage = filepath.Join(path.Templates, "successUpload.html")

}

//StaticCSS gets the path of static css files
func StaticCSS(css string) string {
	return filepath.Join(path.CSS, css)
}

//StaticJS gets the path of static JS files
func StaticJS(js string) string {
	return filepath.Join(path.Javascript, js)
}
