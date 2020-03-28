package pageRoute

import (
	"blog/modal"
	"blog/utils"
	"fmt"
	tem "html/template"
	"net/http"
	"net/url"
)

func commonHandler(template string, _ string, _ string) http.HandlerFunc  {
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

func IndexHandler(template string, mySqlIp string, mySqlUser string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)

		Article := modal.ArticleTable {
			Host: mySqlIp,
			User: mySqlUser,
			Password: "18675270821",
		}
		db, err := Article.StartDb()
		utils.LogError(err, "connect mysql")
		defer  db.Close()

		articles, err := Article.GetArticles(10)
		utils.LogError(err, "get articles")

		t1.Execute(w, articles)
	}
}

func ArticleHandler(template string, mySqlIp string, mySqlUser string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)

		//vars := mux.Vars(r)
		//id := vars["id"]
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		utils.LogError(err, "url.ParseQuery")
		id := queryForm["id"][0]

		Article := modal.ArticleTable {
			Host: mySqlIp,
			User: mySqlUser,
			Password: "18675270821",
		}
		db, err := Article.StartDb()
		utils.LogError(err, "connect mysql")
		defer  db.Close()

		article, err := Article.GetAticle(id)
		utils.LogError(err, "get articles")

		t1.Execute(w, article)
	}
}


