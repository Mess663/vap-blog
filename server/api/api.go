package api

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(template string, mySqlUser string) http.HandlerFunc
}

type routes []route

var ApiRoutes = routes{
	route{
		"PostArticle",
		"POST",
		"/article",
		submitArticle,
	},
}