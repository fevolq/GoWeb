package dao

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"go-web/config"
)

var MysqlDb *sql.DB

func MysqlConn() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.MysqlUser, config.MysqlPwd,
		config.MysqlHost, config.MysqlPost, config.MysqlDb)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if db.Ping() != nil {
		panic("mysql连接异常")
	}
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)

	MysqlDb = db
}
