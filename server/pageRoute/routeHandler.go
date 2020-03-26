package pageRoute

import (
	"blog/modal"
	"blog/utils"
	"fmt"
	tem "html/template"
	"net/http"
)

func commonHandler(template string, _ string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(template)
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)
		t1.Execute(w, "hello world")
	}
}

type IndexParams struct {
	Articles []modal.Article
}

func IndexHandler(template string, mySqlIp string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)

		Article := modal.ArticleTable {
			Host: mySqlIp,
			User: "vaporSpace",
			Password: "18675270821",
		}
		db, err := Article.StartDb()
		utils.LogError(err, "connect mysql")
		defer  db.Close()

		articles, err := Article.GetArticles(10)
		utils.LogError(err, "get articles")

		t1.Execute(w, IndexParams{articles})
	}
}


