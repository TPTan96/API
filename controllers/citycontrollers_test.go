package controllers

import (
	"API_MVC/data"
	"API_MVC/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/globalsign/mgo/bson"

	"github.com/gorilla/mux"
	"github.com/smartystreets/goconvey/convey"
)

func TestCityControllers(t *testing.T) {
	mockVal := &data.MockConnectdata{}
	var datas, respdatas []models.Data
	var Adata, Datanil, respdata models.Data
	var k bson.M

	var jsondatas = []byte(`[
		{
			"_id": "01001",
			"city": "AGAWAM",
			"loc": [
				-72.622739,
				42.070206
			],
			"pop": 1996,
			"state": "MA"
		},
		{
			"_id": "01008",
			"city": "BLANDFORD",
			"loc": [
				-72.936114,
				42.182949
			],
			"pop": 1240,
			"state": "MA"
		}
		]`)
	err := json.Unmarshal(jsondatas, &datas)
	if err != nil {
		t.Error("error:", err)
	}
	var adatabyte = []byte(`{
			"_id": "1996",
			"city": "AGAWAM",
			"loc": [
				-72.622739,
				42.070206
			],
			"pop": 1996,
			"state": "MA"
		}`)

	err = json.Unmarshal(adatabyte, &Adata)
	//err = json.NewDecoder(bytes.NewReader(adatabyte)).Decode(&Adata)
	if err != nil {
		t.Error("error:", err)
	}

	var dataupdate = []byte(`{"state": "MA"}`)
	k = bson.M{"state": "MA"}
	if err != nil {
		t.Error("error:", err)
	}
	mockVal.On("FindAll").Return(datas, nil)
	mockVal.On("FindByID", "1996").Return(Adata, nil)
	mockVal.On("FindByID", "123").Return(Datanil, fmt.Errorf("not found"))
	mockVal.On("Insert", Adata).Return(nil)
	mockVal.On("MaxID").Return(1995)
	mockVal.On("Delete", "1996").Return(nil)
	mockVal.On("CountID", "1996").Return(1)
	mockVal.On("CountID", "123").Return(0)
	mockVal.On("Update", "1996", k).Return(nil)
	colleccity = mockVal

	convey.Convey("create a router", t, func() {
		router := mux.NewRouter()
		router.HandleFunc("/city", GetAllDatas).Methods("GET")
		router.HandleFunc("/city/{id}", GetAData).Methods("GET")
		router.HandleFunc("/city", CreateAData).Methods("POST")
		router.HandleFunc("/city/{id}", DeleteAData).Methods("DELETE")
		router.HandleFunc("/city/{id}", UpdateAData).Methods("PATCH")
		// testing /city/ GET
		convey.Convey("when the request GET /city is hanled  by Router", func() {
			req, err := http.NewRequest("GET", "/city", nil)
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be a 200", func() {
				convey.So(w.Code, convey.ShouldEqual, 200)
				err = json.Unmarshal(w.Body.Bytes(), &respdatas)
				if err != nil {
					t.Error(err)
				}
				Equal := reflect.DeepEqual(respdatas, datas)
				convey.So(Equal, convey.ShouldEqual, true)
			})
		})
		// testing /city/{id} GET
		convey.Convey("when the request GET /city/{id} is hanled  by Router", func() {
			req, err := http.NewRequest("GET", "/city/1996", nil)
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  200", func() {
				convey.So(w.Code, convey.ShouldEqual, 200)
				err = json.Unmarshal(w.Body.Bytes(), &respdata)
				if err != nil {
					t.Error(err)
				}
				Equal := reflect.DeepEqual(respdata, Adata)
				convey.So(Equal, convey.ShouldEqual, true)

			})
		})
		convey.Convey("when the request GET /city/{id} is hanled  by Router but the id doesn't exist", func() {
			req, err := http.NewRequest("GET", "/city/123", nil)
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)

			})
		})
		//testing /city create a data  "POST"
		convey.Convey("when the request POST /city is hanled  by Router", func() {
			req, err := http.NewRequest("POST", "/city", bytes.NewReader(adatabyte))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be 200", func() {
				convey.So(w.Code, convey.ShouldEqual, 200)
				err = json.Unmarshal(w.Body.Bytes(), &respdata)
				if err != nil {
					t.Error(err)
				}
				Equal := reflect.DeepEqual(respdata, Adata)
				convey.So(Equal, convey.ShouldEqual, true)
			})
		})
		convey.Convey("when the request POST /city is hanled by Router with empty body", func() {
			req, err := http.NewRequest("POST", "/city", strings.NewReader(""))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be 400", func() {
				convey.So(w.Code, convey.ShouldEqual, 400)
			})
		})
		//testing /city/id delete a data  "DELETE"
		convey.Convey("when the request DELETE /city/1996 is hanled  by Router", func() {
			req, err := http.NewRequest("DELETE", "/city/1996", nil)
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  200", func() {
				convey.So(w.Code, convey.ShouldEqual, 200)
			})
		})
		//testing /city/id updatea data  "PATCH"
		convey.Convey("when the request PATCH /city/1996 is hanled  by Router", func() {
			req, err := http.NewRequest("PATCH", "/city/1996", bytes.NewReader(dataupdate))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  200", func() {
				convey.So(w.Code, convey.ShouldEqual, 200)
				err = json.Unmarshal(w.Body.Bytes(), &respdata)
				if err != nil {
					t.Error(err)
				}
				Equal := reflect.DeepEqual(respdata, Adata)
				convey.So(Equal, convey.ShouldEqual, true)

			})
		})
		convey.Convey("when the request PATCH /city/123 is hanled  by Router but city unexist", func() {
			req, err := http.NewRequest("PATCH", "/city/123", bytes.NewReader(dataupdate))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  204", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusNoContent)
			})
		})
		convey.Convey("when the request PATCH /city/1996 is hanled  by Router but the body is empty", func() {
			req, err := http.NewRequest("PATCH", "/city/1996", strings.NewReader(""))
			if err != nil {
				t.Error(err)
			}
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			convey.Convey("Then response should	be  400", func() {
				convey.So(w.Code, convey.ShouldEqual, http.StatusBadRequest)
			})
		})
	})
}
