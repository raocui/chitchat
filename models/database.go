package models

//连接数据库
import (
	"chitchat/conf"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = Connect(conf.Config.Mysql.User, conf.Config.Mysql.Password, conf.Config.Mysql.Host, conf.Config.Mysql.Port, conf.Config.Mysql.DbName)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Connect(user, password, host, port, dbname string) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	return sql.Open("mysql", connStr)

}
