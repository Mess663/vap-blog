package modal

type Article struct {
	Id int
	Title string
	Content string
	Time string
}

type MysqlConf struct {
	Ip string
	User string
	Password string
}