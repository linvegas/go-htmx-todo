package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/linvegas/go-htmx-todo/internal/routes"
)

func main() {
  router := routes.NewRouter()

  port := 4242
  addr := fmt.Sprintf(":%d", port)

  fmt.Printf("Listening on: http://localhost%s\n", addr)

  err := http.ListenAndServe(addr, router)

  if err != nil {
    log.Fatal(err)
  }
}
