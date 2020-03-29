package pageRoute

import (
	"blog/modal"
	"fmt"
	"net/http"
	"path/filepath"
)

type route struct {
	Name        string
	Pattern     string
	Template string
	HandlerFunc func(template string, mySqlConf modal.MysqlConf) http.HandlerFunc
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
		"/article",
		StaticPath("article.html"),
		ArticleHandler,
	},
}

func StaticPath(url string) string {
	return fmt.Sprintf("%s/%s", webStaticPath, url)
}
