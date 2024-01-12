package TPBlog

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var tmpl = make(map[string]*template.Template)

var (
	_, b, _, _ = runtime.Caller(0)
	path       = filepath.Dir(b) + "/"
)

// Establishing the fileserver to make assets directory available for the client.
func fileServer() {
	fs := http.FileServer(http.Dir(path + "assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// runServer runs the server in a goroutine and open the browser at the correct URL and then loops waiting for the "stop" input to end all goroutines.
func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 5)
	cmd := exec.Command("explorer", url)
	cmd.Run()
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}

// Run is the public function that executes all necessary functions to run the server and website.
func Run() {
	tmplPath := path + "templates/"
	tmpl["index"] = template.Must(template.ParseFiles(tmplPath+"index.gohtml", tmplPath+"base.gohtml"))
	tmpl["category"] = template.Must(template.ParseFiles(tmplPath+"category.gohtml", tmplPath+"base.gohtml"))
	tmpl["article"] = template.Must(template.ParseFiles(tmplPath+"article.gohtml", tmplPath+"base.gohtml"))
	tmpl["search"] = template.Must(template.ParseFiles(tmplPath+"search.gohtml", tmplPath+"base.gohtml"))
	tmpl["login"] = template.Must(template.ParseFiles(tmplPath+"login.gohtml", tmplPath+"base.gohtml"))
	tmpl["createuser"] = template.Must(template.ParseFiles(tmplPath+"createuser.gohtml", tmplPath+"base.gohtml"))
	tmpl["modifyuser"] = template.Must(template.ParseFiles(tmplPath+"modifyuser.gohtml", tmplPath+"base.gohtml"))
	tmpl["admin"] = template.Must(template.ParseFiles(tmplPath+"admin.gohtml", tmplPath+"base.gohtml"))
	tmpl["addarticle"] = template.Must(template.ParseFiles(tmplPath+"addarticle.gohtml", tmplPath+"base.gohtml"))
	tmpl["modifyarticle"] = template.Must(template.ParseFiles(tmplPath+"modifyarticle.gohtml", tmplPath+"base.gohtml"))
	tmpl["deletearticle"] = template.Must(template.ParseFiles(tmplPath+"deletearticle.gohtml", tmplPath+"base.gohtml"))
	tmpl["about"] = template.Must(template.ParseFiles(tmplPath+"about.gohtml", tmplPath+"base.gohtml"))
	tmpl["error404"] = template.Must(template.ParseFiles(tmplPath+"error404.gohtml", tmplPath+"base.gohtml"))
	routes()
	fileServer()
	runServer()
}
