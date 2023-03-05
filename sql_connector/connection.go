package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ecommerce")
	if err != nil {
		return err
	}

	// Set the maximum number of connections in the pool
	db.SetMaxOpenConns(10)

	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(5)

	// Ensure the connection is valid
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database connection established")

	DB = db

	return nil
}
