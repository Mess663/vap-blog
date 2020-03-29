package server

import (
	"blog/CONST"
	"blog/modal"
	"blog/server/api"
	"blog/server/pageRoute"
	"blog/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(isDev bool, psw string) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	mysqlConf := modal.MysqlConf{
		Password: psw,
	}

	if isDev {
		mysqlConf.Ip = CONST.DEV_MYSQL_HOST
		mysqlConf.User = CONST.DEV_MYSQL_USER
	} else {
		mysqlConf.Ip = CONST.PRO_MYSQL_HOST
		mysqlConf.User = CONST.PRO_MYSQL_USER
	}

	// 页面路由
	for _, route := range pageRoute.PageRoutes {
		handler := utils.Logger(route.HandlerFunc(route.Template, mysqlConf), route.Name)

		router.
			Methods("GET").
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// 静态资源
	router.PathPrefix("/static/").Handler(http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir(pageRoute.StaticPath("static/")))))

	// api
	for _, route := range api.ApiRoutes {
		handler := utils.Logger(route.HandlerFunc(mysqlConf), route.Name)
		router.
			Methods(route.Method).
			Path(fmt.Sprintf("/api%s", route.Pattern)).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
