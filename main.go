package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServe := &staticFileHandeler{"TextHandler !"}
	mux.Handle("/static/", fileServe)
	mux.Handle("/article", &articleHandler{})
	mux.Handle("/admin", &adminHandler{})
	mux.Handle("/", &indexHandler{})


	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type staticFileHandeler struct {
	responseText string
}

func (th *staticFileHandeler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("./web/dist/static"))).ServeHTTP(w, r)
}

type indexHandler struct {}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("./web/dist/index.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

type adminHandler struct {}

func (ih *adminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("./web/dist/admin.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

type articleHandler struct {}

func (ih *articleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("./web/dist/article.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}