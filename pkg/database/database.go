package database

import (
	"database/sql"
	"goblog/pkg/logger"
	"time"

	"github.com/go-sql-driver/mysql"
)

// DB 数据库对象
var DB *sql.DB

// Initialize 初始化数据库
func Initialize() {
	initDB()
	createTable()
}

func initDB() {
	var err error

	config := mysql.Config{
		User:                 "litianyu",
		Passwd:               "lty123456",
		Addr:                 "81.70.8.101:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	// 准备数据库链接池, DSN 全称Data Source Name，表示数据源信息
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲连接数, <= 0时表示无限大，默认值为2
	DB.SetMaxIdleConns(20)
	// 设置每个链接过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)

	// 尝试链接， 失败后会报错
	err = DB.Ping()
	logger.LogError(err)
}

// CreateTable 创建主要的表单
func createTable() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	); `

	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
