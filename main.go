package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/controllers/person"
	"github.com/mickelsonm/gotodo/controllers/team"
	"github.com/mickelsonm/gotodo/controllers/todo"
)

var (
	listenAddr = flag.Int("port", 9090, "http listen port")
)

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/person/{id}", person.GetPerson).Methods("GET")
	router.HandleFunc("/person/{id}", person.DeletePerson).Methods("DELETE")
	router.HandleFunc("/person", person.UpdatePerson).Methods("PUT")
	router.HandleFunc("/person", person.AddPerson).Methods("POST")
	router.HandleFunc("/person", person.GetAllPeople).Methods("GET")

	router.HandleFunc("/team/{id}", team.GetTeam).Methods("GET")
	router.HandleFunc("/team/{id}", team.DeleteTeam).Methods("DELETE")
	router.HandleFunc("/team", team.UpdateTeam).Methods("PUT")
	router.HandleFunc("/team", team.AddTeam).Methods("POST")
	router.HandleFunc("/team", team.GetAllTeams).Methods("GET")

	router.HandleFunc("/todo/{id}", todo.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", todo.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/todo", todo.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo", todo.AddTodo).Methods("POST")
	router.HandleFunc("/todo", todo.GetAllTodos).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Dude, it's an API.")
	})

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%d", *listenAddr))
}
