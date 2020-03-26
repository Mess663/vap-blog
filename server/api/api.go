package api

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var ApiRoutes = routes{
	route{
		"Index",
		"POST",
		"/api/article",
		submitArticle,
	},
}