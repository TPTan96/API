package controllers

import (
	"API_MVC/common"
	"API_MVC/models"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var colleccity common.Connectiondata //interface type
//GetCitySession get session
func GetCitySession(typedata common.Connectiondata) {
	colleccity = typedata
	return
}

// GetAllDatas  get all data from mongoDB
func GetAllDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := colleccity.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// GetAData  get a data from mongoDB
func GetAData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	data, err := colleccity.FindByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

// CreateAData  create a data from mongoDB
func CreateAData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var city models.Data
	err := json.NewDecoder(r.Body).Decode(&city)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please type down data")
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	n := colleccity.MaxID()
	city.ID = strconv.Itoa(n + 1)
	err = colleccity.Insert(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(city)

}

// DeleteAData   delete a data from mongoDB
func DeleteAData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	err := colleccity.Delete(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(bson.M{"remove": "OK"})
}

// UpdateAData  /////////////////
func UpdateAData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if colleccity.CountID(params["id"]) != 0 {
		var k bson.M
		err := json.NewDecoder(r.Body).Decode(&k)
		switch {
		case err == io.EOF:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Please type down data")
			return
		case err != nil:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = colleccity.Update(params["id"], k)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := colleccity.FindByID(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Don't exist Id")
	}
}
