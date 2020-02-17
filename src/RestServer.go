package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
	"encoding/json"
	db "Go_code/src/mysql"
)

func main() {
	fmt.Println("Serwer running in http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/task", GetTasks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var con = db.CreateCon()
	tasks := db.GetAll(con)
	n := tasks.Len()
	result := make([]db.Task, n)
	i := 0
	for temp := tasks.Front(); temp != nil; temp = temp.Next() {
		result[i] = temp.Value.(db.Task)
		i++
	}
	json.NewEncoder(w).Encode(result)
}
