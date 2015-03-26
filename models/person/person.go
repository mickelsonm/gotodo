package person

import (
	"errors"

	"github.com/mickelsonm/gotodo/helpers/database"
	"github.com/mickelsonm/gotodo/helpers/email"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	peopleCollectionName = "People"
)

type People []Person
type Person struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	FirstName   string        `bson:"firstname" json:"firstName"`
	LastName    string        `bson:"lastname" json:"lastName"`
	Email       string        `bson:"email" json:"email"`
	PhoneNumber string        `bson:"phonenumber" json:"phoneNumber"`
}

func GetAllPeople() (people People, err error) {
	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return
	}
	defer sess.Close()

	err = sess.DB(database.MongoConnectionString().Database).C(peopleCollectionName).Find(bson.M{}).All(&people)

	return
}

func (p *Person) Get() error {
	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	qs := bson.M{}

	if bson.IsObjectIdHex(p.Id.Hex()) {
		qs["_id"] = p.Id
	}

	if p.FirstName != "" {
		qs["firstname"] = p.FirstName
	}
	if p.LastName != "" {
		qs["lastname"] = p.LastName
	}
	if p.Email != "" {
		qs["email"] = p.Email
	}
	if p.PhoneNumber != "" {
		qs["phonenumber"] = p.PhoneNumber
	}

	err = session.DB(database.MongoConnectionString().Database).C(peopleCollectionName).Find(qs).One(&p)

	return err
}

func (p *Person) Create() error {
	if p.FirstName == "" {
		return errors.New("First name is required")
	}
	if p.Email == "" {
		return errors.New("Email is required")
	}
	if !email.IsEmail(p.Email) {
		return errors.New("Email is invalid")
	}

	sess, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer sess.Close()

	col := sess.DB(database.MongoConnectionString().Database).C(peopleCollectionName)
	p.Id = bson.NewObjectId()
	err = col.Insert(p)
	return err
}

func (p *Person) Update() error {
	if p.Id.Hex() == "" {
		return errors.New("Invalid person ID")
	}
	if p.FirstName == "" {
		return errors.New("First name is required")
	}
	if p.Email == "" {
		return errors.New("Email is required")
	}
	if !email.IsEmail(p.Email) {
		return errors.New("Email is invalid")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	col := session.DB(database.MongoConnectionString().Database).C(peopleCollectionName)
	err = col.UpdateId(p.Id, p)
	return err
}

func (p *Person) Delete() error {
	if p.Id == "" {
		return errors.New("Invalid person ID")
	}

	session, err := mgo.DialWithInfo(database.MongoConnectionString())
	if err != nil {
		return err
	}
	defer session.Close()

	if err = p.Get(); err != nil {
		return err
	}

	err = session.DB(database.MongoConnectionString().Database).C(peopleCollectionName).RemoveId(p.Id)
	return err
}
