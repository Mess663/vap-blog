package api

import (
	"blog/modal"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(mySqlConf modal.MysqlConf) http.HandlerFunc
}

type routes []route

var ApiRoutes = routes{
	route{
		"PostArticle",
		"POST",
		"/article",
		submitArticle,
	},
	route{
		"PostArticle",
		"post",
		"/isAdmin",
		checkIsAdmin,
	},
}