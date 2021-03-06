package todo

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/models/todo"
	"gopkg.in/mgo.v2/bson"
)

func GetAllTodos(rw http.ResponseWriter, req *http.Request) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	toJSON(rw, todos)
}

func GetTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	if err = getID(rw, req, &t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = t.Get(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(rw, t)
}

func AddTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	if err = json.NewDecoder(req.Body).Decode(&t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = t.Create(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(rw, t)
}

func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	if err = json.NewDecoder(req.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = t.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(w, t)
}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	if err = getID(w, req, &t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = t.Delete(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	toJSON(w, t)
}

func getID(rw http.ResponseWriter, req *http.Request, t *todo.Todo) error {
	if req != nil && t != nil {
		idStr := mux.Vars(req)["id"]
		if bson.IsObjectIdHex(idStr) {
			t.Id = bson.ObjectIdHex(idStr)
			return nil
		}
	}
	return errors.New("Error getting todo ID")
}

func toJSON(rw http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}
