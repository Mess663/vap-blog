package api

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(template string) http.HandlerFunc
}

type routes []route

var ApiRoutes = routes{
	route{
		"Index",
		"POST",
		"/article",
		submitArticle,
	},
}