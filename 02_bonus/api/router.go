package main

import (
	"github.com/gorilla/mux"
	"Users"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
 
type Routes []Route
 
func NewRouter() *mux.Router {
 
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Users.Home,
    },
    Route{
        "Heartbeat",
        "GET",
        "/pulse",
        heartbeat,
    },
    Route{
        "Users",
        "GET",
        "/users",
        Users.GetAllUsers,
    },
    Route{
        "Users",
        "GET",
        "/users/{email}",
        Users.GetByEmail,
    },
    Route{
		"Users by email",
		"DELETE",
		"/users/email/{email}",
		Users.DeleteUser,
    },
}