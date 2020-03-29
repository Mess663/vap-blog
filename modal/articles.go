package modal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type ArticleTable struct {
	Host string
	User string
	Password string
	Db *sql.DB
}

func (t *ArticleTable) StartDb() (*sql.DB, error)  {
	dataSourceName :=fmt.Sprintf("%s:%s@tcp(%s)/%s", t.User, t.Password, t.Host, "blog")
	DB, err := sql.Open("mysql", dataSourceName)
	t.Db = DB
	return DB, err
}

func (t *ArticleTable) InsertItem(title string, content string) (sql.Result, error) {
	stmt, err := t.Db.Prepare("INSERT INTO articles(title, content) values(?, ?)")
	if err != nil {
		return nil, err
	}

	return stmt.Exec(title, content)
}

func (t *ArticleTable) UpdateItem(id int, title string, content string) (sql.Result, error) {
	return t.Db.Exec("UPDATE articles SET title=?, content=? where id=?", title, content, id)
}

func (t *ArticleTable) GetArticles(limit int) ([]Article, error) {
	var articleList []Article
	rows, err :=  t.Db.Query("SELECT * FROM articles LIMIT ?", limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var title string
		var content string
		var timeStamp string
		err = rows.Scan(&id, &content,&title,  &timeStamp)
		if err != nil {
			return nil, err
		}

		article := Article{
			id,
			title,
			content,
			strings.Split(timeStamp, " ")[0],
		}

		articleList = append(articleList, article)
	}

	return articleList, nil
}

func (t *ArticleTable) GetAticle(id string) (Article, error) {
	rows, err := t.Db.Query("SELECT * FROM articles where id=?", id)
	if err != nil {
		return Article{}, err
	}

	var a Article
	for rows.Next() {
		var id int
		var title string
		var content string
		var timeStamp string
		err = rows.Scan(&id, &content,&title,  &timeStamp)
		if err != nil {
			return Article{}, err
		}

		a.Id = id
		a.Title = title
		a.Content = content
		a.Time = strings.Split(timeStamp, " ")[0]
	}

	return a, nil
}

func (t *ArticleTable) DeleteItem(id int) (sql.Result, error) {
	return t.Db.Exec("DELETE from articles where id=?", id)
}
