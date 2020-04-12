package postgres

import (
	"database/sql"
	"container/list"
)

const (
	SELECT_PEOPLE = "SELECT id, name, surname FROM public.people;"

)


type People struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
	SURNAME string `json:"name"`
}
func GetPeople(con *sql.DB)[]People  {
	values := list.New()

	results, err := con.Query(SELECT_PEOPLE)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var person People
		err = results.Scan(&person.ID, &person.NAME,&person.SURNAME)
		if err != nil {
			panic(err.Error())
		}
		values.PushFront(person)
	}
	defer results.Close()
	n := values.Len()
	tab:= make([]People, n)
	i := 0
	for temp := values.Front(); temp != nil; temp = temp.Next() {
		tab[i] = temp.Value.(People)
		i++
	}
	return tab
}
