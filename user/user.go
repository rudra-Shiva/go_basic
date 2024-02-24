package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (u *User) OutputUserDetailByAttachongMethodsInStructs() {
	fmt.Println(u.firstName, " ", u.lastName)
	fmt.Println("Created in ", u.createdAt)
	fmt.Println("And his birthdate is ", u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}

}

// one also create function with New() in Goloang it's a common pratice
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("all fields must be filled")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now().Truncate(1), // remove the time part of a Time value
	}, nil
}
