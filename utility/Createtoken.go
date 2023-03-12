package utility

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	*jwt.RegisteredClaims
}

var key = []byte("suhusduhhs65+978t$67")

func Createtoken(username string, password string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &user{username, password, &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), Issuer: "test"}})
	signedString, err := token.SignedString(key)
	if err != nil {
		return ""
	}
	return signedString
}

// 验证token
func ParseWithClaims(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &user{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims, ok := token.Claims.(*user); ok && token.Valid {
		fmt.Printf("%v %v %v", claims.Username, claims.Password, claims.RegisteredClaims.Issuer)
		fmt.Println(claims.Username)

	} else {
		fmt.Println(err)
	}

}
