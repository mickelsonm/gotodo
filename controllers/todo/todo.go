package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/models/todo"
	"gopkg.in/mgo.v2/bson"
)

func GetTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	fmt.Println("GET")

	if err = getID(rw, req, t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = t.Get(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var data []byte
	if data, err = json.Marshal(&t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func AddTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	fmt.Println("ADD")

	t.Text = req.FormValue("text")
	fmt.Printf("text = %s\n", t.Text)

	if err = t.Add(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var data []byte
	if data, err = json.Marshal(&t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func UpdateTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	fmt.Println("UPDATE")

	if err = getID(rw, req, t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("UPDATE"))
}

func DeleteTodo(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t todo.Todo

	fmt.Println("DELETE")

	if err = getID(rw, req, t); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("DELETE"))
}

func getID(rw http.ResponseWriter, req *http.Request, t todo.Todo) error {
	if req != nil {
		idStr := mux.Vars(req)["id"]
		if bson.IsObjectIdHex(idStr) {
			t.Id = bson.ObjectIdHex(idStr)
			return nil
		}
	}
	return errors.New("Error getting Todo ID")
}
