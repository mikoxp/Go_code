package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	host_port = 5432
	username = "postgres"
	password = "postgres"
	database_name = "postgres"
)

func CreateCon() *sql.DB {

	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host_port, hostname, username, password, database_name)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Postgres db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
