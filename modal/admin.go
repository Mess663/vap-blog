package modal

import (
	"blog/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type AdminTable struct {
	Ip string
	User string
	Password string
	Db *sql.DB
}

func (t *AdminTable) StartDb() (*sql.DB, error)  {
	dataSourceName :=fmt.Sprintf("%s:%s@tcp(%s)/%s", t.User, t.Password, t.Ip, "blog")
	DB, err := sql.Open("mysql", dataSourceName)
	t.Db = DB
	return DB, err
}

func (t *AdminTable) IsUserOk(uid string) (bool) {
	rows, err := t.Db.Query("SELECT * FROM `admin` where uid=?", uid)
	if err != nil {
		utils.LogError(err, "IsUserOk")
		return false
	}

	ok := rows.Next()
	fmt.Println("____是否ok：", ok)
	return ok
}