package server

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func Serve() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/tasks", taskHandler)
	http.HandleFunc("/tasks/json", taskJsonHandler)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprender Go",
		Done: true,
	}

	t := template.Must(template.ParseFiles("server/task.html"))
	t.Execute(w, task)
}

func taskJsonHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprender Go",
		Done: true,
	}

	j, _ := json.Marshal(task)
	w.Write(j)
}
