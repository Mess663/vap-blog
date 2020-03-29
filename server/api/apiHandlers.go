package api

import (
	"blog/modal"
	"blog/utils"
	"database/sql"
	"encoding/json"
	"io"
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
	Status int `json:"status"`
	Id int64 `json:"id"`
}

func submitArticle(mySqlConf modal.MysqlConf) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 限制上传数据大小，以防恶意上传
		s, error := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
		utils.LogError(error, "submitArticle read body")
		error = request.Body.Close();
		utils.LogError(error, "submitArticle body close")

		var reqBody map[string]string
		json.Unmarshal(s, &reqBody)
		title := reqBody["title"]
		content := reqBody["content"]

		result, err := insertArticle(title, content, mySqlConf)
		checkErr(err)

		id, err:= result.LastInsertId()
		checkErr(err)

		var data Data
		if err == nil {
			data = Data{1, id}
		} else {
			data = Data{0, 0}
		}

		error = json.NewEncoder(writer).Encode(data)
		utils.LogError(error, "NewEncoder")
	}
}

func insertArticle(title string, content string, mySqlConf modal.MysqlConf) (sql.Result, error) {
	Article := modal.ArticleTable {
		Ip: mySqlConf.Ip,
		User: mySqlConf.User,
		Password: mySqlConf.Password,
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
