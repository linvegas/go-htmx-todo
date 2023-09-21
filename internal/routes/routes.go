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

func indexHandler(w http.ResponseWriter, r *http.Request) {
  tmpl := template.Must(template.ParseFiles(
    "web/pages/index.html",
    ))

  todolist := map[string][]Todo {
    "Todos": {
      {Description: "Learn Go", Done: false},
      {Description: "Finish Bioshock", Done: false},
      {Description: "Build shape", Done: true},
    },
  }

  tmpl.Execute(w, todolist)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
  description := r.FormValue("todo-desc")
  log.Print("HTMX request received")
  log.Printf("Form 'todo-desc': %s\n", description)
  tmpl := template.Must(template.ParseFiles("web/pages/index.html"))
  tmpl.ExecuteTemplate(w, "todo-list-item", Todo{Description: description, Done: false})
}

func NewRouter() http.Handler {
  mux := http.NewServeMux()

  mux.HandleFunc("/", indexHandler)
  mux.HandleFunc("/add-todo", addTodoHandler)
  // mux.HandleFunc("/api/data", apiDataHandler)

  return mux
}
