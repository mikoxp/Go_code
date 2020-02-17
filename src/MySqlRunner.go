package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	db "Go_code/src/mysql"
)

func main() {
	fmt.Println(" Go MySQL")
	var con= db.CreateCon()
	db.Insert(con,"Test")
	tasks:=db.GetAll(con)
	fmt.Println("-------TASKS-----------")
	fmt.Println(tasks.Len())
	for temp := tasks.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value.(db.Task).ID)

	}
}
