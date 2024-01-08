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
	http.HandleFunc("/addarticle", addArticleHandler)
	http.HandleFunc("/addarticle/treatment", addArticleTreatmentHandler)
	http.HandleFunc("/modifyarticle", modifyArticleHandler) // use query params: ?id=<article-id>
	http.HandleFunc("/modifyarticle/treatment", modifyArticleTreatmentHandler)
	http.HandleFunc("/about", aboutHandler)
}
