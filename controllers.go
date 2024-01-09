package TPBlog

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var mySession Session

// Root handler redirects to index handler.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.RedirectHandler("/index", http.StatusSeeOther) // see if it works, otherwise, change to http.Redirect(w, r, "/index", http.StatusSeeOther)
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
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Home",
			StaticPath: "static/",
		},
		Articles: articles,
	}
	err := tmpl["index"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/category" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	category := r.URL.Query().Get("category")
	articles := selectCategory(category)
	data := struct {
		Base     BaseData
		Category []Article
	}{
		Base: BaseData{
			Title:      category,
			StaticPath: "static/",
		},
		Category: articles,
	}
	err := tmpl["category"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/article" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article := selectArticle(id)
	data := struct {
		Base    BaseData
		Article Article
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
	}
	data.Article.Content = formatArticle(article)
	err = tmpl["article"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/search" {
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
	}{
		Base: BaseData{
			Title:      "Research",
			StaticPath: "static/",
		},
		Articles: articles,
		Search:   search,
		Message:  message,
	}
	err := tmpl["search"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/login" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	var message string
	if r.URL.Query().Get("status") == "error" {
		message = "<div class=\"message\">Wrong username or password!</div>"
	}
	data := struct {
		Base    BaseData
		Message string
	}{
		Base: BaseData{
			Title:      "Login - Sport Pulse",
			StaticPath: "static/",
		},
		Message: message,
	}
	err := tmpl["login"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func loginTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/login/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if login(r.FormValue("username"), r.FormValue("password")) {
		fmt.Println("log: loginTreatment() correct login: welcome ", r.FormValue("username"), "!")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		fmt.Println("log: loginTreatment() incorrect login!")
		http.Redirect(w, r, "/login?status=error", http.StatusSeeOther)
	}
}

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
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Sign Up",
			StaticPath: "static/",
		},
		Message: message,
	}
	err := tmpl["createuser"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/createuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	var user = User{
		Name:     r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	if checkUsername(user.Name) {
		if len(user.Password) > 5 {
			user.addUser()
			http.Redirect(w, r, "/user/login", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/user/create?pass=error", http.StatusMovedPermanently)
		}
	} else {
		http.Redirect(w, r, "/user/create?user=error", http.StatusMovedPermanently)
	}
}

func modifyUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyuser" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	status := r.URL.Query().Get("error")
	var message string
	if status == "error" {
		message = "<div class=\"message\">Invalid data!</div>"
	}
	data := struct {
		Base    BaseData
		Message string
	}{
		Base: BaseData{
			Title:      "Sport Pulse - Personal data",
			StaticPath: "static/",
		},
		Message: message,
	}
	err := tmpl["modifyuser"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func modifyUserTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyuser/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
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
		http.Redirect(w, r, "/modifyuser", http.StatusMovedPermanently)
		return
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/admin" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
		User User
	}{
		Base: BaseData{
			Title:      "Dashboard - Sport Pulse",
			StaticPath: "static/",
		},
		User: mySession.MyUser,
	}
	err := tmpl["admin"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func addArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/addarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base       BaseData
		User       User
		Categories []string
		Article    Article
	}{
		Base: BaseData{
			Title:      "New article - Sport Pulse",
			StaticPath: "static/",
		},
		User:       mySession.MyUser,
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
	}
	err := tmpl["addarticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func addArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/addarticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	newCtn := Article{
		Id:       getIdNewArticle(),
		Category: r.FormValue("category"),
		Title:    r.FormValue("title"),
		Author:   mySession.MyUser.Name,
		Date:     time.Now().Format("02/01/2006"),
		Content:  r.FormValue("content"),
	}
	addArticle(newCtn)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func modifyArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyarticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article := selectArticle(id)
	data := struct {
		Base    BaseData
		Article Article
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
	}
	err = tmpl["modifyarticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func modifyArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/modifyarticle/treatment" {
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
		article := selectArticle(idInForm)
		newCtn := Article{
			Id:       idInForm,
			Category: r.FormValue("category"),
			Title:    r.FormValue("title"),
			Author:   article.Author,
			Date:     time.Now().Format("02/01/2006"),
			Content:  r.FormValue("content"),
		}
		modifyArticle(newCtn)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

func deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/deletearticle" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("article"))
	if err != nil {
		log.Fatal("log: articleHandler() strconv.Atoi error!\n", err)
	}
	article := selectArticle(id)
	data := struct {
		Base    BaseData
		Article Article
		Message string
	}{
		Base: BaseData{
			Title:      article.Title,
			StaticPath: "static/",
		},
		Article: article,
		Message: "<div class=\"message\">Do you really want to delete that article ?</div>",
	}
	err = tmpl["deletearticle"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteArticleTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/deletearticle/treatment" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal("log: modifyArticleTreatment() Atoi error!\n", err)
	}
	deleteArticle(id)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // testing
	if r.URL.Path != "/about" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	data := struct {
		Base BaseData
	}{
		Base: BaseData{
			Title:      "Sport Pulse - About",
			StaticPath: "static/",
		},
	}
	err := tmpl["about"].ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	fmt.Printf("log: status: %#v\n", status) // testing
	if status == http.StatusNotFound {
		data := struct {
			Base BaseData
		}{
			Base: BaseData{
				Title:      "Sport Pulse - 404 Not Found",
				StaticPath: "static/",
			},
		}
		err := tmpl["error404"].ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Fatal(err)
		}
	}
}
