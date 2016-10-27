package main

import (
    "log"
    "net/http"
    "github.com/user/cv/routes"
)

func main() {
    router := routes.NewRouter()
    
    // launch server
    log.Fatal(http.ListenAndServe(":8080", router))
}
