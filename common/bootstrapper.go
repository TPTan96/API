package common

import (
	"github.com/globalsign/mgo"
)

//StartUp sdjksjd
func StartUp() (*mgo.Collection, *mgo.Collection) {
	clcity, cluser := initmongoDB()
	return clcity, cluser
}
