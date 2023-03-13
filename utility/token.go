package utility

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type User struct {
	Indently string `json:"indently"`
	Username string `json:"username"`
	*jwt.RegisteredClaims
}

// Createtoken 生成token
func Createtoken(indently string, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &User{indently, username, &jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), Issuer: "test"}})
	signedString, err := token.SignedString(Key)
	if err != nil {
		return ""
	}
	return signedString
}

// ParseWithClaims 验证token
func ParseWithClaims(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("token不能为空！")
	}
	user := User{}
	token, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if claims, ok := token.Claims.(*User); ok && token.Valid {
		fmt.Printf("%v %v %v \n", claims.Indently, claims.Username, claims.RegisteredClaims.Issuer)
		//fmt.Println(claims.Username)

		return claims, nil

	} else {
		log.Println(err)
		return nil, err
	}

}
