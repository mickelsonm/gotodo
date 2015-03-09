package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/controllers/todo"
)

var (
	listenAddr = flag.Int("port", 9090, "http listen port")
)

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/todo/{id}", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", todo.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", todo.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo", todo.AddTodo).Methods("POST")
	router.HandleFunc("/todo", todo.GetAllTodos).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Dude, it's an API.")
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%d", *listenAddr))
}
