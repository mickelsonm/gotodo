package todo

import (
	"errors"

	"github.com/mickelsonm/gotodo/helpers/database"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	Id        bson.ObjectId `string:"id"`
	Text      string        `string:"task"`
	Completed bool          `string:"completed"`
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

	err = sess.DB(database.DATABASE_NAME).C("Todos").Find(bson.M{
		"_id": t.Id,
	}).One(&t)

	return err
}

func (t *Todo) Add() error {
	if t.Text == "" {
		return errors.New("Todo must have text!")
	}
	t.Id = bson.NewObjectId()

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	col := sess.DB(database.DATABASE_NAME).C("Todos")

	if _, err = col.UpsertId(t.Id, t); err != nil {
		return err
	}

	idx := mgo.Index{
		Key:        []string{"text", "completed"},
		Background: true,
		Sparse:     false,
		DropDups:   true,
	}
	col.EnsureIndex(idx)

	return err
}

func (t *Todo) Update() error {
	if t.Id.Hex() == "" {
		return errors.New("Invalid Todo ID")
	}
	if t.Text == "" {
		return errors.New("Todo must have text!")
	}

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	var change = mgo.Change{
		ReturnNew: true,
		Update: bson.M{
			"$set": bson.M{
				"text":      t.Text,
				"completed": t.Completed,
			},
		},
	}

	_, err = sess.DB(database.DATABASE_NAME).C("Todos").Find(bson.M{
		"_id": t.Id,
	}).Apply(change, t)

	return err
}

func (t *Todo) Delete() error {
	if t.Id == "" {
		return errors.New("Invalid Todo ID")
	}

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	if err = t.Get(); err != nil {
		return err
	}

	return sess.DB(database.DATABASE_NAME).C("Todos").RemoveId(t.Id)
}
