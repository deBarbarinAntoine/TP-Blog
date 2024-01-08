package TPBlog

import (
	"log"
	"net/http"
)

var mySession Session

// Root handler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl["index"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Index page handler.
func albumJulHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl["albumJul"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Creating user page handler.
func trackSdmHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl["trackSdm"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}
