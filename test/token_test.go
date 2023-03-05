package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func Test_createtoken(*testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user{"123", "448", jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), Issuer: "test"}})
	signedString, err := token.SignedString([]byte("suhusduhhs65+978t$67"))
	if err != nil {
		return
	}
	fmt.Println(signedString)
}

func Test_parseToken(*testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjEyMyIsInBhc3N3b3JkIjoiNDQ4IiwiaXNzIjoidGVzdCIsImV4cCI6MTY3ODA5ODQ0OH0.n5llexjpe1rzHoph0XeN4BZIQraLUo88BAEQTq2ox1s"
	token, err := jwt.ParseWithClaims(tokenString, &user{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("suhusduhhs65+978t$67"), nil
	})
	if err != nil {
		return
	}
	if claims, ok := token.Claims.(*user); ok && token.Valid {
		fmt.Printf("%v %v %v", claims.Username, claims.Password, claims.RegisteredClaims.Issuer)
	} else {
		fmt.Println(err)
	}
}
