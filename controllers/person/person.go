package person

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/models/person"
	"gopkg.in/mgo.v2/bson"
)

func GetAllPeople(w http.ResponseWriter, req *http.Request) {
	people, err := person.GetAllPeople()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	toJSON(w, people)
}

func GetPerson(rw http.ResponseWriter, req *http.Request) {
	var err error
	var p person.Person

	if err = getID(rw, req, &p); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.Get(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(rw, p)
}

func AddPerson(rw http.ResponseWriter, req *http.Request) {
	var err error
	var p person.Person

	if err = json.NewDecoder(req.Body).Decode(&p); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.Create(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(rw, p)
}

func UpdatePerson(w http.ResponseWriter, req *http.Request) {
	var err error
	var p person.Person

	if err = json.NewDecoder(req.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = p.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toJSON(w, p)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	var err error
	var p person.Person

	if err = getID(w, req, &p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = p.Delete(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	toJSON(w, p)
}

func getID(rw http.ResponseWriter, req *http.Request, p *person.Person) error {
	if req != nil && p != nil {
		idStr := mux.Vars(req)["id"]
		if bson.IsObjectIdHex(idStr) {
			p.Id = bson.ObjectIdHex(idStr)
			return nil
		}
	}
	return errors.New("Error getting person ID")
}

func toJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
