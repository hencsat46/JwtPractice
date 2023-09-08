package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString([]byte("sign string"))

	if err != nil {
		return "", err
	}

	return tokenStr, err
}

func someFunc(w http.ResponseWriter, r *http.Request) {
	newToken, err := CreateJWT()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintf(w, newToken)
}

func main() {
	http.HandleFunc("/api", someFunc)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":3000", nil)
}
