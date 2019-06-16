package controllers

import (
	"API_MVC/data"
	"API_MVC/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/smartystreets/goconvey/convey"
)

func TestUserControllers(t *testing.T) {
	mockVal := &data.MockConnectdata{}
	Acc := models.Account{
		Username: "TanTruong",
		Password: "ahihi",
	}
	byteAcc, err := json.Marshal(Acc)
	if err != nil {
		t.Error(err)
	}
	AccWrongPass := models.Account{
		Username: "TanTruong",
		Password: "ahihihi",
	}
	byteAccWrongPass, err := json.Marshal(AccWrongPass)
	if err != nil {
		t.Error(err)
	}
	AccWrongUsername := models.Account{
		Username: "TanTruon",
		Password: "ahihi",
	}
	byteAccWrongUsername, err := json.Marshal(AccWrongUsername)
	if err != nil {
		t.Error(err)
	}

	mockVal.On("FindByName", Acc.Username).Return(Acc, nil)
	mockVal.On("CheckByName", Acc.Username).Return(true, nil)
	mockVal.On("CheckByName", AccWrongUsername.Username).Return(false, nil)
	mockVal.On("AddNewUser", AccWrongUsername).Return(nil)

	collecuser = mockVal

	convey.Convey("Create a user router", t, func() {
		router := mux.NewRouter()
		router.HandleFunc("/login", Login).Methods("GET")
		router.HandleFunc("/signup", AddNewUser).Methods("POST")
		//test GET /login
		convey.Convey("When the request GET /login", func() {
			req, err := http.NewRequest("GET", "/login", strings.NewReader(""))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
		convey.Convey("When the request GET /login with an exist account", func() {
			req, err := http.NewRequest("GET", "/login", bytes.NewReader(byteAcc))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  200", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
		convey.Convey("When the request GET /login with an exist account but wrong password", func() {
			req, err := http.NewRequest("GET", "/login", bytes.NewReader(byteAccWrongPass))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  406", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusNotAcceptable)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
		convey.Convey("When the request GET /login with an exist account but wrong username", func() {
			req, err := http.NewRequest("GET", "/login", bytes.NewReader(byteAccWrongUsername))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  406", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusNotAcceptable)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})

		//test POST /signup
		convey.Convey("When the request POST /signup", func() {
			req, err := http.NewRequest("POST", "/signup", strings.NewReader(""))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
		convey.Convey("When the request POST /signup add a user but already exist", func() {
			req, err := http.NewRequest("POST", "/signup", bytes.NewReader(byteAcc))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
		convey.Convey("When the request POST /signup add an unexist user", func() {
			req, err := http.NewRequest("POST", "/signup", bytes.NewReader(byteAccWrongUsername))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			convey.Convey("Then response should be  200", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				resp := w.Body.String()
				fmt.Println(resp)
				convey.So(resp, convey.ShouldNotEqual, ``)
			})
		})
	})

}
