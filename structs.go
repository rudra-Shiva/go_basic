package main

import (
	"fmt"

	"example.com/main/user"
)

func main() {

	fmt.Println("Structs Uses")
	userFirstName := getUserData("Enter firstName")
	userLastName := getUserData("Enter LastName")
	userBirthDate := getUserData("Enter UserBirthDate(MM/DD/YYYY)")

	var appUser *user.User
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now().Truncate(1), // remove the time part of a Time value
	// }
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Printf("Error creating user\n %v", err)
	}
	admin := user.NewAdmin("aa@a.in", "12346")
	admin.OutputUserDetailByAttachongMethodsInStructs()
	admin.ClearUserName()
	admin.OutputUserDetailByAttachongMethodsInStructs()
	//also create empty struc
	// appUser = user{}

	//define also like that
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthDate,
	// 	time.Now(),
	// }

	//define also like that
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	time.Now(),
	// }
	// outputUserDetail(appUser)
	// outputUserDetails(&appUser)
	appUser.OutputUserDetailByAttachongMethodsInStructs()
	appUser.ClearUserName()
	appUser.OutputUserDetailByAttachongMethodsInStructs()
}

// // by using pointer we can modify data from outside of function
// func outputUserDetails(u *user) {
// 	fmt.Println(u.firstName, " ", (*u).lastName) //technically correct way
// 	fmt.Println("Created in ", u.createdAt)
// 	fmt.Println("And his birthdate is ", u.birthDate)
// }
// func outputUserDetail(u user) {
// 	fmt.Println(u.firstName, " ", u.lastName)
// 	fmt.Println("Created in ", u.createdAt)
// 	fmt.Println("And his birthdate is ", u.birthDate)
// }

func getUserData(userInput string) string {
	fmt.Print("\n" + userInput + ": ")
	var userInputValue string
	_, err := fmt.Scanln(&userInputValue)
	if err != nil {
		fmt.Println(err)
	}
	return userInputValue

}
