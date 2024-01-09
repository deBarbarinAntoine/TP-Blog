package TPBlog

import (
	"fmt"
	"log"
	"net/http"
)

var mySession Session

// Root handler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.RedirectHandler("/index", http.StatusSeeOther) // see if it works, otherwise, change to http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// Index page handler.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/index" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["index"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/category" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["category"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/article" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["article"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/search" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["search"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/login" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["login"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loginTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/login/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if login(r.FormValue("username"), r.FormValue("password")) {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/createuser" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["createuser"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/createuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// creating user
}

func modifyUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/modifyuser" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["modifyuser"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func modifyUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/modifyuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// modifying user
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/admin" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["admin"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func addArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/addarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["addarticle"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func addArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/addarticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// addArticle
}

func modifyArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/modifyarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["modifyarticle"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func modifyArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/modifyarticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// modifyArticle
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/deletearticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["deletearticle"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/deletearticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	// deleteArticle
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
	if r.URL.Path != "/about" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err := tmpl["about"].ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Printf("log: status: %#v\n", status) // for testing purposes
	if status == http.StatusNotFound {
		err := tmpl["error404"].ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
