package main

import (
  "log"
  "net/http"

	"github.com/gorilla/handlers"

  "cv-server-rest-go/routes"
)

func main() {
  router := routes.NewRouter()
  

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

  // launch server
  log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
