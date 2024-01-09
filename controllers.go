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
	category := r.URL.Query().Get("category")
	data := struct {
		Base     BaseData
		category string
	}{
		Base: BaseData{
			Title:      category,
			StaticPath: "static/",
			Line:       "",
		},
		category: category,
	}
	err := tmpl["category"].ExecuteTemplate(w, "base", data)
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
	article := r.URL.Query().Get("article")
	data := struct {
		Base    BaseData
		Article string
	}{
		Base: BaseData{
			Title:      article,
			StaticPath: "static/",
			Line:       "",
		},
		Article: article,
	}
	err := tmpl["article"].ExecuteTemplate(w, "base", data)
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
	articles := searchArticle(r.URL.Query().Get("q"))
	data := struct {
		Base     BaseData
		Articles []Article
	}{
		Base: BaseData{
			Title:      "Research",
			StaticPath: "static/",
			Line:       "",
		},
		Articles: articles,
	}
	err := tmpl["search"].ExecuteTemplate(w, "base", data)
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
		fmt.Println("log: loginTreatment() correct login: welcome ", r.FormValue("username"), "!")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		fmt.Println("log: loginTreatment() incorrect login!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("log: UrlPath: %#v\n", r.URL.Path) // for testing purposes
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
	var user = User{
		Name:     r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	if checkUsername(user.Name) {
		if len(user.Password) > 5 {
			user.addUser()
			http.Redirect(w, r, "/user/login", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/user/create", http.StatusMovedPermanently)
		}
	} else {
		http.Redirect(w, r, "/user/create", http.StatusMovedPermanently)
	}
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
			//mySession.MyGameData.Message = "Données invalides !"
			//mySession.MyGameData.MessageClass = "message red"
			fmt.Println("log: modifyUserTreatment() error: Invalid data!")
			http.Redirect(w, r, "/modifyuser", http.StatusSeeOther)
			return
		}
	} else {
		http.Redirect(w, r, "/modifyuser", http.StatusMovedPermanently)
		return
	}
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
	articles, err := RetrieveArticles()
	if err != nil {
		log.Fatal("log: RetrieveArticles() error!\n", err)
	}
	id := len(articles) + 1
	newCtn := Article{
		Id:       id,
		Category: r.FormValue("category"),
		Title:    r.FormValue("title"),
		Author:   mySession.MyUser.Name,
		Date:     fmt.Sprint(time.DateOnly),
		Content:  r.FormValue("content"),
	}
	addArticle(newCtn)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal("log: modifyArticleTreatment() Atoi error!\n", err)
	}
	newCtn := Article{
		Id:       id,
		Category: r.FormValue("category"),
		Title:    r.FormValue("title"),
		Author:   r.FormValue("author"),
		Date:     fmt.Sprint(time.DateOnly),
		Content:  r.FormValue("content"),
	}
	modifyArticle(newCtn)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
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
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal("log: modifyArticleTreatment() Atoi error!\n", err)
	}
	deleteArticle(id)
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
