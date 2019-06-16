package controllers

import (
	"API_MVC/common"
	"API_MVC/models"
	"encoding/json"
	"io"
	"net/http"
)

var collecuser common.Connectiondata //interface type

//GetUserSession get session
func GetUserSession(typedata common.Connectiondata) {
	collecuser = typedata
	return
}
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Acc models.Account
	err := json.NewDecoder(r.Body).Decode(&Acc)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(`Please fill Username and Password`)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := collecuser.CheckByName(Acc.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode("The user is unexist")
	} else {
		User, err := collecuser.FindByName(Acc.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if Acc.Username == User.Username && Acc.Password == User.Password {
			TokenString, err := common.GenerateJWT(Acc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			} else {
				json.NewEncoder(w).Encode(TokenString)
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode("Worng Password")

		}
	}

}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Acc models.Account
	err := json.NewDecoder(r.Body).Decode(&Acc)
	switch {
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please fill Username and Password")
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := collecuser.CheckByName(Acc.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username exist")
	} else {
		err = collecuser.AddNewUser(Acc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode("Sign up success")
	}

}
