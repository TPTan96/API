package common

import (
	"API_MVC/models"
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Connectiondata interface {
	FindAll() ([]models.Data, error)
	FindByID(id string) (models.Data, error)
	CountID(id string) int
	MaxID() int
	Insert(data models.Data) error
	Delete(id string) error
	Update(id string, k bson.M) error
	FindByName(name string) (models.Account, error)
	AddNewUser(User models.Account) error
	CheckByName(name string) (bool, error)
}

// Dao just dao
type Dao struct {
	Server   string
	Database string
}

const (
	//COLLECTION name of collection
	COLLECTIONCITY = "city"
	COLLECTIONUSER = "user"
)

var server Dao
var db *mgo.Database

func initmongoDB() (*mgo.Collection, *mgo.Collection) {
	server.Server = "localhost:27017"
	server.Database = "test"
	clcity, cluser := server.Conect()
	return clcity, cluser
}

func (m *Dao) Conect() (*mgo.Collection, *mgo.Collection) {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		fmt.Println("mongo conect err")
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	return db.C(COLLECTIONCITY), db.C(COLLECTIONUSER)
}
