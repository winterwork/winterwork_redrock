package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// GetDb 获取数据库
func GetDb() *sql.DB {
	return Db
}

// OpenDb 连接数据库
func OpenDb() {
	db, err := sql.Open("mysql", "root:John888456123.@tcp(121.41.120.238:3306)/douban")
	CheckErr(err)
	Db = db
	return
}
