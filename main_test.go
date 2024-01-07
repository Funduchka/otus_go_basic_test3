package main

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	user1 := User{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
	}
	fmt.Println(user1)
}
