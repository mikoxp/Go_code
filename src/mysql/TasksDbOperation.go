package mysql

import (
	"fmt"
	"database/sql"
	"container/list"
)

var INSERT_FORMAT ="INSERT INTO `tasks`(`title`) VALUES ('%s')"
var SELECT="SELECT * FROM tasks"

type Task struct {
	ID   int    `json:"id"`
	TITLE string `json:"title"`
}

func Insert(con *sql.DB,title string)  {
	var s = fmt.Sprintf(INSERT_FORMAT,title)
	insert, err := con.Query(s)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func GetAll(con *sql.DB) *list.List  {
	values := list.New()

	results, err := con.Query(SELECT)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var task Task
		err = results.Scan(&task.ID, &task.TITLE)
		if err != nil {
			panic(err.Error())
		}
		values.PushFront(task)
	}
	defer results.Close()

	return values
}


