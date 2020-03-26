package server

import (
	"blog/server/api"
	"blog/server/pageRoute"
	"fmt"
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
			Path(fmt.Sprintf("/api%s", route.Pattern)).
			Name(route.Name).
			Handler(route.HandlerFunc("207.148.99.103:3306"))
	}

	// test api
	for _, route := range api.ApiRoutes {
		router.
			Methods(route.Method).
			Path(fmt.Sprintf("/test/api%s", route.Pattern)).
			Name(route.Name).
			Handler(route.HandlerFunc("0.0.0.0:3306"))
	}

	return router
}