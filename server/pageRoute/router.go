package pageRoute

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type route struct {
	Name        string
	Pattern     string
	Template string
	HandlerFunc func(template string) http.HandlerFunc
}

var webStaticPath, _ = filepath.Abs("web/dist")

type routes []route

var PageRoutes = routes{
	route{
		"Index",
		"/",
		StaticPath("index.html"),
		commonHandler,
	},
	route{
		"Admin",
		"/admin",
		StaticPath("admin.html"),
		commonHandler,
	},
	route{
		"Article",
		"/article/{todoId}",
		StaticPath("article.html"),
		commonHandler,
	},
	route{
		"Article",
		"/article/{todoId}",
		StaticPath("article.html"),
		commonHandler,
	},
}

func StaticPath(url string) string {
	return fmt.Sprintf("%s/%s", webStaticPath, url)
}

type ServerHandler struct {
	Url string
}

func (s ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(s.Url)
	t1, err := template.ParseFiles(s.Url)
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

type StaticFileHandeler struct {
}

func (th *StaticFileHandeler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("./web/dist/static"))).ServeHTTP(w, r)
}

type ApiPostArticle struct {
}

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
type B struct {
	a string
}

func (th *ApiPostArticle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//s, _ := ioutil.ReadAll(r.Body)

	b := B{}
	myJsonString := `{"some":"json"}`
	err := json.Unmarshal([]byte(myJsonString), &b)
	if err != nil {
		log.Printf("解析json错误：%v", err)
	}
	fmt.Println(b)

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}