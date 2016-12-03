package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/8tomat8/GoRepost"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

func main() {
	var host = flag.String("host", "127.0.0.123", "Host.")
	var port = flag.String("port", "8181", "Port.")
	flag.Parse()

	fmt.Println("Creating router...")
	router := NewRouter()
	fmt.Println("Router created.")

	fmt.Println("Starting HTTP server at http://" + *host + ":" + *port)
	glog.Fatal(http.ListenAndServe(*host+":"+*port, router))
}

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
		"Create",
		"POST",
		"/tasks",
		GoRepost.TaskCreate,
	},
	Route{
		"Create",
		"GET",
		"/",
		GoRepost.Greating,
	},
	Route{
		"TasksList",
		"GET",
		"/tasks",
		GoRepost.TasksList,
	},
	Route{
		"TaskStatus",
		"GET",
		"/tasks/{taskId}",
		GoRepost.TaskStatus,
	},
}
