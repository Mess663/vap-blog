package pageRoute

import (
	"fmt"
	tem "html/template"
	"net/http"
)

func commonHandler(template string) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(template)
		t1, err := tem.ParseFiles(template)
		if err != nil {
			panic(err)
		}
		t1.Execute(w, "hello world")
	}
}

