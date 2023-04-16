package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
}

func Setup() {

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "node_test",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)

	db.SetMaxIdleConns(10)

	DB = db

	fmt.Println("database connection pool established")
}
