package team

import (
	"errors"

	"github.com/mickelsonm/gotodo/helpers/database"
	"github.com/mickelsonm/gotodo/models/person"
	"github.com/mickelsonm/gotodo/models/todo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	teamsCollectionName = "Teams"
)

type Teams []Team
type Team struct {
	Id      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Members person.People `bson:"members" json:"members"`
	Todos   todo.Todos    `bson:"todos" json:"todos"`
}

func GetAllTeams() (teams Teams, err error) {
	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return
	}
	defer sess.Close()

	err = sess.DB(database.MongoConnectionString().Database).C(teamsCollectionName).Find(bson.M{}).All(&teams)

	return
}

func (t *Team) Get() error {
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	qs := bson.M{}

	if bson.IsObjectIdHex(t.Id.Hex()) {
		qs["_id"] = t.Id
	}

	if t.Name != "" {
		qs["name"] = t.Name
	}

	err = session.DB(database.MongoConnectionString().Database).C(teamsCollectionName).Find(qs).One(&t)

	return err
}

func (t *Team) Create() error {
	if t.Name == "" {
		return errors.New("Team must have a name")
	}
	if len(t.Members) < 1 {
		return errors.New("Team must have at least one member")
	}

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	col := sess.DB(database.MongoConnectionString().Database).C(teamsCollectionName)
	t.Id = bson.NewObjectId()
	err = col.Insert(t)
	return err
}

func (t *Team) Update() error {
	if t.Id.Hex() == "" {
		return errors.New("Invalid team ID")
	}
	if t.Name == "" {
		return errors.New("Team must have a name")
	}
	if len(t.Members) < 1 {
		return errors.New("Team must have at least one member")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	col := session.DB(database.MongoConnectionString().Database).C(teamsCollectionName)
	err = col.UpdateId(t.Id, t)
	return err
}

func (t *Team) Delete() error {
	if t.Id == "" {
		return errors.New("Invalid team ID")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	if err = t.Get(); err != nil {
		return err
	}

	err = session.DB(database.MongoConnectionString().Database).C(teamsCollectionName).RemoveId(t.Id)
	return err
}
