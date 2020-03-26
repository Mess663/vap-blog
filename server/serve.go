package server

import (
	"blog/server/api"
	"blog/server/pageRoute"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// 页面路由
	for _, route := range pageRoute.PageRoutes {
		router.
			Methods("GET").
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc(route.Template))
	}

	// 静态资源
	router.PathPrefix("/static/").Handler(http.StripPrefix(
	"/static/",
		http.FileServer(http.Dir(pageRoute.StaticPath("static/")))))

	// api
	for _, route := range api.ApiRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}