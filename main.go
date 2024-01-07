package main

import "fmt"

func main() {

}

// явно определенная структура
type User struct {
	Name     string
	LastName string
	Age      int
	Role     Role
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

func SaveUser(user User) error {
	if err := ValidateUser(user); err != nil {
		return err
	}

	// save user to DB
	return nil
}

func ValidateUser(user User) error {
	if user.Age > 100 {
		return fmt.Errorf("User is too old")
	}
	return nil
}
