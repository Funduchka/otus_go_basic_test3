package main

import "fmt"

func main() {

}

// явно определенная структура
type User struct {
	Name     string
	LastName string
	Age      int
	Role     struct {
		Name        string
		Permissions []string
	}
}

func NewUser() User {
	return User{
		Role: struct {
			Name        string
			Permissions []string
		}{
			Name:        "admin",
			Permissions: []string{"read", "write"},
		},
	}
}

func SaveUser(userToSave User) error {
	if mySuperError := ValidateUser(userToSave); mySuperError != nil {
		return mySuperError
	}

	// save user to DB
	return nil
}

func ValidateUser(me User) error {
	if me.Age > 100 {
		return fmt.Errorf("user is too old")
	}

	return nil
}

type Role struct {
	Name        string
	Permissions []string
}

func Anon() {
	user1 := User{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
	}
	fmt.Println(user1)
}
