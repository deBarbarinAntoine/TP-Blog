package TPBlog

import "net/http"

//	routes
//
// initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/category", categoryHandler) // use query params: ?category=<category-name>
	http.HandleFunc("/article", articleHandler)   // use query params: ?article=<article-id>
	http.HandleFunc("/search", searchHandler)     // use query params: ?q=<search>
	http.HandleFunc("/login", loginHandler)       // use query params: ?status=<error>
	http.HandleFunc("/login/treatment", loginTreatmentHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/createuser", createUserHandler) // use query params: ?<input>=error (input: "pass" or "user")
	http.HandleFunc("/createuser/treatment", createUserTreatmentHandler)
	http.HandleFunc("/modifyuser", modifyUserHandler) // use query params: ?status=error
	http.HandleFunc("/modifyuser/treatment", modifyUserTreatmentHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/addarticle", addArticleHandler)
	http.HandleFunc("/addarticle/treatment", addArticleTreatmentHandler)
	http.HandleFunc("/modifyarticle", modifyArticleHandler)                    // use query params: ?article=<article-id>
	http.HandleFunc("/modifyarticle/treatment", modifyArticleTreatmentHandler) // use query params: ?article=<article-id> to secure data
	http.HandleFunc("/deletearticle", deleteArticleHandler)                    // use query params: ?article=<article-id>
	http.HandleFunc("/deletearticle/treatment", deleteArticleTreatmentHandler) // use query params: ?article=<article-id>
	http.HandleFunc("/about", aboutHandler)
}
