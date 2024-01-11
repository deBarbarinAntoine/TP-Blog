package TPBlog

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var mySession Session

// rootHandler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

// Index page handler.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/index" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	articles := randomArticles()
	data := struct {
		Base     BaseData
		Articles []Article
		Session  Session
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Home",
			StaticPath: "static/",
		},
		Articles: articles,
		Session:  mySession,
	}
	err := tmpl["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	categoryHandler
//
// fetch and show a list of all Article of the Article.Category indicated in the query params.
//
// query params: ?category=<category-name>
func categoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/category" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("category") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	category := r.URL.Query().Get("category")
	articles := selectCategory(category)
	if len(articles) == 0 {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base        BaseData
		MainArticle Article
		Category    []Article
		Session     Session
	}{
		Base: BaseData{
			Title:      category,
			StaticPath: "static/",
		},
		MainArticle: articles[0],
		Category:    articles[1:],
		Session:     mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["category"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	articleHandler
//
// fetch and show a specific Article which id number is indicated in the query params.
//
// Query params: ?article=<article-id>
func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/article" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("article") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article, ok := selectArticle(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base    BaseData
		Article Article
		Session Session
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
		Session: mySession,
	}
	fmt.Printf("log: data before formatArticle(): %#v\n", data) // testing
	data.Article.Content = formatArticle(article)
	fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["article"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	searchHandler
//
// fetch and show all Article which title matches the search indicated in the query params.
//
// Query params: ?q=<search>
func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/search" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if !r.URL.Query().Has("q") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	search := r.URL.Query().Get("q")
	articles := searchArticle(search)
	var message string
	if len(articles) == 0 {
		message = "<div class=\"message\">There is no article matching your research!</div>"
	}
	data := struct {
		Base     BaseData
		Articles []Article
		Search   string
		Message  string
		Session  Session
	}{
		Base: BaseData{
			Title:      "Research",
			StaticPath: "static/",
		},
		Articles: articles,
		Search:   search,
		Message:  message,
		Session:  mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["search"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	loginHandler
//
// takes the User info to send it to loginTreatmentHandler via Post Method.
//
// Optional query params: ?status=<error>	(error: "error" or "restricted")
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/login" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	var message string
	switch r.URL.Query().Get("status") {
	case "error":
		message = "<div class=\"message\">Wrong username or password!</div>"
	case "restricted":
		message = "<div class=\"message\">You need to sign in to access to this resource!</div>"
	}
	data := struct {
		Base    BaseData
		Message string
		Session Session
	}{
		Base: BaseData{
			Title:      "Login - Sport Pulse",
			StaticPath: "static/",
		},
		Message: message,
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["login"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	loginTreatmentHandler
//
// checks the form values sent by loginHandler to open the session and redirect to adminHandler
// or redirect to loginHandler with query params: ?status=error
func loginTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/login/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if login(r.FormValue("Username"), r.FormValue("Password")) {
		fmt.Println("log: loginTreatment() correct login: welcome ", r.FormValue("Username"), "!")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		fmt.Println("log: loginTreatment() incorrect login!")
		http.Redirect(w, r, "/login?status=error", http.StatusSeeOther)
	}
}

//	logoutHandler
//
// close and clear the Session opened.
// It also clears the cache so that the Session can be closed.
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/logout" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	fmt.Println("log: logout() passing through!")
	mySession.Close()
	fmt.Printf("log logout() session cleared: %#v\n", mySession)
	http.Redirect(w, r, "/index", http.StatusMovedPermanently)
}

//	createUserHandler
//
// takes the new User info to send it to createUserTreatmentHandler via Post Method.
//
// Optional query params: ?pass=error or ?user=error
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/createuser" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	var message string
	if pass == "error" {
		message = "<div class=\"message\">The password must contain at least 5 characters!</div>"
	}
	if user == "error" {
		message = "<div class=\"message\">Username already used!</div>"
	}
	data := struct {
		Base    BaseData
		Message string
		Session Session
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Sign Up",
			StaticPath: "static/",
		},
		Message: message,
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["createuser"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	createUserTreatmentHandler
//
// checks the form values sent by createUserHandler and calls User.addUser to sign up the new User.
//
// In case of invalid values, it redirects to createUserHandler with ?pass=error or ?user=error query params.
func createUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/createuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	confirmPassword := r.FormValue("Password-check")
	var user = User{
		Name:     r.FormValue("Username"),
		Password: r.FormValue("Password"),
	}
	if checkUsername(user.Name) {
		if len(user.Password) > 5 && user.Password == confirmPassword {
			user.addUser()
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/createuser?pass=error", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/createuser?user=error", http.StatusSeeOther)
	}
}

//	modifyUserHandler
//
// takes the User new info to send it to modifyUserTreatmentHandler via Post Method.
//
// Optional query params: ?status=error
func modifyUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyuser" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	var message string
	if r.URL.Query().Get("status") == "error" {
		message = "<div class=\"message\">Invalid data!</div>"
	}
	data := struct {
		Base    BaseData
		Message string
		Session Session
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Personal data",
			StaticPath: "static/",
		},
		Message: message,
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["modifyuser"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	modifyUserTreatmentHandler
//
// checks the User new info and runs User.modifyUser with the new info.
//
// If new info is invalid, it redirects to modifyUserHandler with ?status=error query params.
func modifyUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	if r.Method == http.MethodPost {
		fmt.Println("log: modifyUserTreatment() update user")
		username := r.FormValue("username")
		newPassword := r.FormValue("newPassword")
		if (checkUsername(username) || username == mySession.MyUser.Name) && mySession.MyUser.Password == r.FormValue("password") && len(newPassword) > 5 {
			fmt.Println("log: modifyUserTreatment() Previous name: ", mySession.MyUser.Name)
			fmt.Println("log: modifyUserTreatment() Previous password: ", mySession.MyUser.Password)
			fmt.Println()
			newUser := User{Name: username, Password: newPassword}
			err := mySession.MyUser.modifyUser(newUser)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("log: modifyUserTreatment() New name: ", mySession.MyUser.Name)
			fmt.Println("log: modifyUserTreatment() New password: ", mySession.MyUser.Password)
			fmt.Println()
			http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
			return
		} else {
			fmt.Println("log: modifyUserTreatment() error: Invalid data!")
			http.Redirect(w, r, "/modifyuser?status=error", http.StatusSeeOther)
			return
		}
	} else {
		http.Redirect(w, r, "/modifyuser", http.StatusSeeOther)
		return
	}
}

//	adminHandler
//
// shows all Article and permits access to addArticleHandler, modifyArticleHandler and deleteArticleHandler
// for each Article shown.
func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/admin" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	articles, err := retrieveArticles()
	if err != nil {
		log.Fatal("log: retrieveArticles() error!\n", err)
	}
	adminGuard(w, r)
	data := struct {
		Base     BaseData
		Articles []Article
		Session  Session
	}{
		Base: BaseData{
			Title:      "Dashboard - Sport Pulse",
			StaticPath: "static/",
		},
		Articles: articles,
		Session:  mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["admin"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	addArticleHandler
//
// takes the new Article info to send it to addArticleTreatmentHandler via Post Method.
func addArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/addarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	data := struct {
		Base       BaseData
		Categories []string
		Article    Article
		Session    Session
	}{
		Base: BaseData{
			Title:      "New article - Sport Pulse",
			StaticPath: "static/",
		},
		Categories: []string{"Formule 1", "Esport", "Football"},
		Article: Article{
			Id:           getIdNewArticle(),
			Category:     "",
			Title:        "",
			Author:       mySession.MyUser.Name,
			Date:         time.Now().Format("02/01/2006"),
			BigImg:       "",
			SmallImg:     "",
			Introduction: "",
			Content:      "",
		},
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["addarticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	addArticleTreatmentHandler
//
// runs addArticle to save the new Article in the json data file (articles.json).
func addArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/addarticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	newCtn := Article{
		Id:       getIdNewArticle(),
		Category: r.FormValue("category"),
		Title:    r.FormValue("title"),
		Author:   mySession.MyUser.Name,
		Date:     time.Now().Format("02/01/2006"),
		Content:  r.FormValue("content"),
	}
	addArticle(newCtn)
	fmt.Printf("log: newCtn: %#v\n", newCtn) // testing
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

//	modifyArticleHandler
//
// select and shows the Article which id is indicated in the query params
// and takes the Article new info to send it to addArticleTreatmentHandler via Post Method.
//
// Query params: ?article=<article-id>
func modifyArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	if !r.URL.Query().Has("article") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article, ok := selectArticle(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base    BaseData
		Article Article
		Session Session
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["modifyarticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	modifyArticleTreatmentHandler
//
// checks if the id number in the form is the same as the one in the query params (to avoid some problems)
// and runs modifyArticle to update the Article with the new content.
//
// Query params: ?article=<article-id>
func modifyArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyarticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	if !r.URL.Query().Has("article") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	idInQuery, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	idInForm, err2 := strconv.Atoi(r.FormValue("id"))
	if err2 != nil {
		log.Fatal("log: modifyArticleTreatment() Atoi error!\n", err)
	}
	// INFO: Checking if the Id wasn't forcefully changed in the Form, or in the action attribute (a little security)
	if idInQuery != idInForm {
		http.Redirect(w, r, "/admin?status=error", http.StatusSeeOther)
	} else {
		article, ok := selectArticle(idInForm)
		if !ok {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		newCtn := Article{
			Id:       idInForm,
			Category: r.FormValue("category"),
			Title:    r.FormValue("title"),
			Author:   article.Author,
			Date:     time.Now().Format("02/01/2006"),
			Content:  r.FormValue("content"),
		}
		modifyArticle(newCtn)
		fmt.Printf("log: updatedCtn: %#v\n", newCtn) // testing
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

//	deleteArticleHandler
//
// select the Article which id is indicated in the query params and shows it in the form
// to send it to deleteArticleTreatmentHandler via Post Method.
//
// Query params: ?article=<article-id>
func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/deletearticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	if !r.URL.Query().Has("article") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article, ok := selectArticle(id)
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base    BaseData
		Article Article
		Message string
		Session Session
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
		Message: "<div class=\"message\">Do you really want to delete that article ?</div>",
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err = tmpl["deletearticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	deleteArticleTreatmentHandler
//
// runs deleteArticle to remove the Article which id is indicated in the query params.
// It then redirects to adminHandler.
//
// Query params: ?article=<article-id>
func deleteArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/deletearticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	adminGuard(w, r)
	if !r.URL.Query().Has("article") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal("log: modifyArticleTreatment() Atoi error!\n", err)
	}
	deleteArticle(id)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

//	aboutHandler
//
// shows the website map and info and all Terms and Conditions.
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/about" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base    BaseData
		Session Session
	}{
		Base: BaseData{
			Title:      "Sport Pulse - About",
			StaticPath: "static/",
		},
		Session: mySession,
	}
	fmt.Printf("log: data: %#v\n", data) // testing
	err := tmpl["about"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

//	errorHandler
//
// shows the custom error 404 page.
// It is called whenever the url is unknown or incorrect.
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Printf("log: status: %#v\n", status) // testing
	if status == http.StatusNotFound {
		data := struct {
			Base    BaseData
			Session Session
		}{
			Base: BaseData{
				Title:      "Sport Pulse - 404 Not Found",
				StaticPath: "static/",
			},
			Session: mySession,
		}
		fmt.Printf("log: data: %#v\n", data) // testing
		err := tmpl["error404"].ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//	adminGuard
//
// checks if Session.IsOpen to access the restricted area of the website.
// Else, it redirects to loginHandler with ?status=restricted query params.
func adminGuard(w http.ResponseWriter, r *http.Request) {
	if mySession.IsOpen {
		return
	} else {
		http.Redirect(w, r, "/login?status=restricted", http.StatusSeeOther)
	}
}
