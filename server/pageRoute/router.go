package pageRoute

import (
	"fmt"
	"net/http"
	"path/filepath"
)

type route struct {
	Name        string
	Pattern     string
	Template string
	HandlerFunc func(template string, mySqlIp string) http.HandlerFunc
}

var webStaticPath, _ = filepath.Abs("web/dist")

type routes []route

var PageRoutes = routes{
	route{
		"Index",
		"/",
		StaticPath("index.html"),
		IndexHandler,
	},
	route{
		"Admin",
		"/admin",
		StaticPath("admin.html"),
		commonHandler,
	},
	route{
		"Article",
		"/article/{todoId}",
		StaticPath("article.html"),
		commonHandler,
	},
}

func StaticPath(url string) string {
	return fmt.Sprintf("%s/%s", webStaticPath, url)
}
