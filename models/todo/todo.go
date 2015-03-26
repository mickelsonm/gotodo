package todo

import (
	"errors"

	"github.com/mickelsonm/gotodo/helpers/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	todosCollectionName = "Todos"
)

type Todos []Todo
type Todo struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Text      string        `bson:"text" json:"text"`
	Assigned  bool          `bson:"assigned" json:"assigned"`
	Completed bool          `bson:"completed" json:"completed"`
}

func GetAllTodos() (todos Todos, err error) {
	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return
	}
	defer sess.Close()

	err = sess.DB(database.MongoConnectionString().Database).C(todosCollectionName).Find(bson.M{}).All(&todos)

	return
}

func (t *Todo) Get() error {
	if t.Id.Hex() == "" {
		return errors.New("Invalid Todo ID")
	}

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	err = sess.DB(database.MongoConnectionString().Database).C(todosCollectionName).Find(bson.M{
		"_id": t.Id,
	}).One(&t)

	return err
}

func (t *Todo) Create() error {
	if t.Text == "" {
		return errors.New("Todo must have text!")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	col := session.DB(database.MongoConnectionString().Database).C(todosCollectionName)
	t.Id = bson.NewObjectId()
	err = col.Insert(t)
	return err
}

func (t *Todo) Update() error {
	if t.Id.Hex() == "" {
		return errors.New("Invalid Todo ID")
	}
	if t.Text == "" {
		return errors.New("Todo must have text!")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	col := session.DB(database.MongoConnectionString().Database).C(todosCollectionName)
	err = col.UpdateId(t.Id, t)
	return err
}

func (t *Todo) Delete() error {
	if t.Id == "" {
		return errors.New("Invalid Todo ID")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	if err = t.Get(); err != nil {
		return err
	}

	err = session.DB(database.MongoConnectionString().Database).C(todosCollectionName).RemoveId(t.Id)
	return err
}
