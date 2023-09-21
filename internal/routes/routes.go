package routes

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
  Description string
  Done        bool
}

var todolist = map[string][]Todo {
  "Todos": {
    {Description: "Learn Go", Done: false},
    {Description: "Finish Bioshock", Done: false},
    {Description: "Build shape", Done: true},
  },
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles("web/pages/index.html"))
  tmpl.Execute(w, todolist)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
  description := r.FormValue("todo-desc")
  newTodo := Todo{Description: description, Done: false}

  log.Print("HTMX request received")
  log.Printf("Form 'todo-desc': %s\n", description)

  tmpl := template.Must(template.ParseFiles("web/pages/index.html"))
  tmpl.ExecuteTemplate(w, "todo-list-item", newTodo)
}

func NewRouter() http.Handler {
  mux := http.NewServeMux()

  mux.HandleFunc("/", indexHandler)
  mux.HandleFunc("/add-todo", addTodoHandler)

  return mux
}
