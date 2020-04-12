package main

import (
	"fmt"
	_ "github.com/lib/pq"
	db "Go_code/src/postgres"
)


func main() {
	fmt.Println("Postgres db connect")
	var con= db.CreateCon()

	people:=db.GetPeople(con)
	fmt.Println("-------PEOPLE-----------")
	for _, v := range people {
		fmt.Printf("id: %d, name: %s, surname: %s\n",v.ID,v.NAME,v.SURNAME)
	}
	defer con.Close()


}
