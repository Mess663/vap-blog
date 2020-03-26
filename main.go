package main

import (
	"blog/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	// 判断系统输入
	//fmt.Println(os.Args[1], os.Args[2])

	router := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

	//mux := http.NewServeMux()
	//mux.Handle("/api", &server.ApiPostArticle{})
	//mux.Handle("/static/", &server.StaticFileHandeler{})
	//mux.Handle("/article", &server.ServerHandler{"./web/dist/article.html"})
	//mux.Handle("/admin", &server.ServerHandler{"./web/dist/admin.html"})
	//mux.Handle("/", &server.ServerHandler{"./web/dist/index.html"})
	//
	//
	//err := http.ListenAndServe(":8080", mux)
	//
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//Article := modal.ArticleTable {
	//	Host: "207.148.99.103:3306",
	//	User: "vaporSpace",
	//	Password: "18675270821",
	//}
	//db, err := Article.StartDb()
	//checkErr(err)
	//defer  db.Close()
	//
	//_, err = Article.DeleteItem(3)
	//checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Printf("mySQL error: %v", err)
	}
}
