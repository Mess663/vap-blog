package server

import (
	"blog/CONST"
	"blog/server/api"
	"blog/server/pageRoute"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	host := CONST.DEV_MYSQL_HOST
	user := CONST.DEV_MYSQL_USER

	// 页面路由
	for _, route := range pageRoute.PageRoutes {
		router.
			Methods("GET").
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc(route.Template, host, user))
	}

	// 静态资源
	router.PathPrefix("/static/").Handler(http.StripPrefix(
	"/static/",
		http.FileServer(http.Dir(pageRoute.StaticPath("static/")))))

	// api
	for _, route := range api.ApiRoutes {
		fmt.Println(fmt.Sprintf("/test/api%s", route.Pattern))
		// production
		router.
			Methods(route.Method).
			Path(fmt.Sprintf("/api%s", route.Pattern)).
			Name(route.Name).
			Handler(route.HandlerFunc(host, user))
	}

	return router
}