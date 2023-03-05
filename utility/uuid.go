package utility

import (
	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	u2 := uuid.NewV4().String()
	return u2
}
