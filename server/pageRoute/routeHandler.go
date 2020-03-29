package pageRoute

import (
	"blog/modal"
	"blog/utils"
	"fmt"
	tem "html/template"
	"net/http"
	"net/url"
)

func commonHandler(template string, _ modal.MysqlConf) http.HandlerFunc  {
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

func IndexHandler(template string, mySqlConf modal.MysqlConf) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)

		Article := modal.ArticleTable {
			Ip: mySqlConf.Ip,
			User: mySqlConf.User,
			Password: mySqlConf.Password,
		}
		db, err := Article.StartDb()
		utils.LogError(err, "connect mysql")
		defer  db.Close()

		articles, err := Article.GetArticles()
		utils.LogError(err, "get articles")

		t1.Execute(w, articles)
	}
}

func ArticleHandler(template string, mySqlConf modal.MysqlConf) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		t1, err := tem.ParseFiles(template)
		utils.LogError(err, template)

		//vars := mux.Vars(r)
		//id := vars["id"]
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		utils.LogError(err, "url.ParseQuery")
		id := queryForm["id"][0]

		Article := modal.ArticleTable {
			Ip: mySqlConf.Ip,
			User: mySqlConf.User,
			Password: mySqlConf.Password,
		}
		db, err := Article.StartDb()
		utils.LogError(err, "connect mysql")
		defer  db.Close()

		article, err := Article.GetAticle(id)
		utils.LogError(err, "get articles")

		t1.Execute(w, article)
	}
}


