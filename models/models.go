package models

import (
	"github.com/globalsign/mgo/bson"
)

type (
	Data struct {
		ID       string    `bson:"_id" json:"_id"`
		Name     string    `bson:"city" json:"city"`
		Location []float64 `bson:"loc" json:"loc"`
		Pop      int       `bson:"pop" json:"pop"`
		State    string    `bson:"state" json:"state"`
	}
	Account struct {
		ID       bson.ObjectId `bson:"_id" json:"_id"`
		Username string        `bson:"Name" json "Name"`
		Password string        `bson:"Password" json:"Password"`
	}
)
