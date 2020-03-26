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

	log.Fatal(http.ListenAndServe(":8080", router))
}
