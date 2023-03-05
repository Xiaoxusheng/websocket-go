package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func Test_uuid(t *testing.T) {
	// Creating UUID Version 4
	// panic on error

	// or error handling
	u2 := uuid.NewV4().String()

	fmt.Println(u2)

	//// Parsing UUID from string input
	//u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	//if err != nil {
	//	fmt.Printf("Something went wrong: %s", err)
	//	return
	//}
	//fmt.Printf("Successfully parsed: %s", u2)
}
