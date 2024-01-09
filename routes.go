package TPBlog

import "net/http"

// routes initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/category", categoryHandler) // use query params: ?category=<category-name>
	http.HandleFunc("/article", articleHandler)   // use query params: ?id=<article-id>
	http.HandleFunc("/search", searchHandler)     // use query params: ?q=<search>
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/login/treatment", loginTreatmentHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/createuser", createUserHandler)
	http.HandleFunc("/createuser/treatment", createUserTreatmentHandler)
	http.HandleFunc("/modifyuser", modifyUserHandler)
	http.HandleFunc("/modifyuser/treatment", modifyUserTreatmentHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/addarticle", addArticleHandler)
	http.HandleFunc("/addarticle/treatment", addArticleTreatmentHandler)
	http.HandleFunc("/modifyarticle", modifyArticleHandler) // use query params: ?id=<article-id>
	http.HandleFunc("/modifyarticle/treatment", modifyArticleTreatmentHandler)
	http.HandleFunc("/deletearticle", deleteArticleHandler) // use query params: ?id=<article-id>
	http.HandleFunc("/deletearticle/treatment", deleteArticleTreatmentHandler)
	http.HandleFunc("/about", aboutHandler)
}
