package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"API_MVC/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var mykey = []byte("Fantatis Baby")

//GenerateJWT generate a token
func GenerateJWT(Acc models.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": Acc.Username,
		"password": Acc.Password,
		"exp":      time.Now().Add(time.Minute * 20).Unix(),
	})
	tokenString, error := token.SignedString(mykey)
	if error != nil {
		fmt.Println(error)
		return "", error
	}
	return tokenString, error
}

//Validate sdkjksdkj
func Validate(w http.ResponseWriter, r *http.Request, endpoint http.HandlerFunc) {
	if r.URL.Path != "/login" && r.URL.Path != "/signup" {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mykey, nil
			})
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Not Authorized")
		}
	} else {
		endpoint(w, r)
	}

}
