package server

import (
	"fmt"
	"html/template"
	"net/http"
)

type ServerHandler struct {
	Url string
}

func (s ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(s.Url)
	t1, err := template.ParseFiles(s.Url)
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

type StaticFileHandeler struct {
}

func (th *StaticFileHandeler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("./web/dist/static"))).ServeHTTP(w, r)
}