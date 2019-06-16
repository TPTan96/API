package data

import (
	"API_MVC/models"
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Col struct {
	Cl *mgo.Collection
}

func (colec Col) FindAll() ([]models.Data, error) {
	var data []models.Data
	err := colec.Cl.Find(nil).All(&data)
	return data, err
}

func (colec Col) FindByID(id string) (models.Data, error) {
	var data models.Data
	err := colec.Cl.FindId(id).One(&data)
	return data, err
}
func (colec Col) CountID(id string) int {
	n, _ := colec.Cl.FindId(id).Count()
	return n
}
func (colec Col) MaxID() int {
	var data models.Data
	_ = colec.Cl.Find(nil).Sort("-_id").One(&data)
	n, _ := strconv.Atoi(data.ID)
	return n
}
func (colec Col) Insert(data models.Data) error {
	err := colec.Cl.Insert(&data)
	return err
}

func (colec Col) Delete(id string) error {
	err := colec.Cl.Remove(bson.M{"_id": id})
	return err
}

func (colec Col) Update(id string, k bson.M) error {
	err := colec.Cl.UpdateId(id, bson.M{"$set": k})
	return err
}
