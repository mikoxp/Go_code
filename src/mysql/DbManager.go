package mysql

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
