package main

import (
	"blog/server"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	// 判断系统输入 4
	var isDev bool
	var psw string

	if len(os.Args) > 1 {
		isDev = false
		psw = os.Args[1]
	} else {
		isDev = true
		psw = "123"
	}

	router := server.NewRouter(isDev, psw)

	log.Fatal(http.ListenAndServe(":8080", router))
}
