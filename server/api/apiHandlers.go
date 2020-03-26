package api

import (
	"blog/modal"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
type Data struct {
	status int
}

func submitArticle(w http.ResponseWriter, r *http.Request)  {
	s, _ := ioutil.ReadAll(r.Body)
	var reqBody map[string]string
	json.Unmarshal(s, &reqBody)
	title := reqBody["title"]
	content := reqBody["content"]

	_, err := insertArticle(title, content)
	checkErr(err)

	var data Data
	if err == nil {
		data = Data{0}
	} else {
		data = Data{1}
	}

	json.NewEncoder(w).Encode(data)
}

func insertArticle(title string, content string) (sql.Result, error) {
	Article := modal.ArticleTable {
		Host: "207.148.99.103:3306",
		User: "vaporSpace",
		Password: "18675270821",
	}
	db, err := Article.StartDb()
	checkErr(err)
	defer  db.Close()

	return Article.InsertItem(title, content)
}

func checkErr(err error) {
	if err != nil {
		log.Printf("mySQL error: %v", err)
	}
}
