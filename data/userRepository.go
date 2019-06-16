package data

import (
	"API_MVC/models"

	"github.com/globalsign/mgo/bson"
)

func (colec Col) FindByName(name string) (models.Account, error) {
	var User models.Account
	err := colec.Cl.Find(bson.M{"Name": name}).One(&User)
	return User, err
}

func (colec Col) AddNewUser(User models.Account) error {
	obj_id := bson.NewObjectId()
	User.ID = obj_id
	err := colec.Cl.Insert(User)
	return err
}

func (colec Col) CheckByName(name string) (bool, error) {
	v, err := colec.Cl.Find(bson.M{"Name": name}).Count()
	if v == 0 {
		return false, err
	} else {
		return true, err
	}

}
