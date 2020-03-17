package server

import (
	"html/template"
	"net/http"
)

type ServerHandler struct {
	Router string
	Url string
	Handler func(http.ResponseWriter, *http.Request)
}

func (ServerHandler) serveHTTP(w http.ResponseWriter, r *http.Request)  {
	t1, err := template.ParseFiles(ServerHandler.Url)
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}


func InitRouter(routers []string)  {
	for r := range routers {
		s := ServerHandler{r}
	}
}