package routes

import (
    "net/http"
    "cv-server-rest-go/handlers"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Candidate",
        "GET",
        "/candidate/{key}",
        handlers.GetCandidate,
    },
    Route{
        "CreateCandidate",
        "POST",
        "/candidate",
        handlers.CreateCandidate,
    },
    Route{
        "DeleteCandidate",
        "DELETE",
        "/candidate/{id}",
        handlers.DeleteCandidate,
    },
}
