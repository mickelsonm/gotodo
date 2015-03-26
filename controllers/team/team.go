package team

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mickelsonm/gotodo/models/team"
	"gopkg.in/mgo.v2/bson"
)

func GetAllTeams(w http.ResponseWriter, req *http.Request) {
	teams, err := team.GetAllTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	toJSON(w, teams)
}

func GetTeam(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t team.Team

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

func AddTeam(rw http.ResponseWriter, req *http.Request) {
	var err error
	var t team.Team

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

func UpdateTeam(w http.ResponseWriter, req *http.Request) {
	var err error
	var t team.Team

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

func DeleteTeam(w http.ResponseWriter, req *http.Request) {
	var err error
	var t team.Team

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

func getID(rw http.ResponseWriter, req *http.Request, t *team.Team) error {
	if req != nil && t != nil {
		idStr := mux.Vars(req)["id"]
		if bson.IsObjectIdHex(idStr) {
			t.Id = bson.ObjectIdHex(idStr)
			return nil
		}
	}
	return errors.New("Error getting team ID")
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
