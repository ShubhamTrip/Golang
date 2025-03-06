package main

import (
	"fmt"
)

type User struct {
	Name       string
	Age        int
	IsLoggedIn bool
	Email      string
}

// This is Method for Struct User
func (u User) getName() {
	fmt.Println("Name is: ", u.Name)
}

// Method with return type
func (u User) getAge() int {
	return u.Age
}

// Setter
func (u User) setEmail(email string) {
	u.Email = email
}

func main() {

	var user1 User = User{"Shubham", 21, true, "shubahm@gmial.com"}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Age)
	fmt.Println(user1.IsLoggedIn)

	user2 := User{Name: "Shubham", Age: 21, IsLoggedIn: true, Email: "ac@gmai.com"}

	fmt.Println(user2)

	user1.getName()

	user2.getName()

	fmt.Println("The age is: ", user1.getAge())

	user1.setEmail("sdfsfsd@gmail.com")

	fmt.Println(user1) // This doesnt change the email because we are not changing the value of the struct
	// To change the value of the struct we need to use pointers
	// We are currently passing the value of the struct to the method
	// What is happenng here is that the value of the struct is being copied to the method
	// So the changes made in the method are not reflected in the original struct

	/*
			The main thing happening

			When you define a method with a value receiver (func (u User) setEmail(email string) {}), Go creates a copy of the struct inside the function.

		     So When we call setEmail on a user instance, Go creates a copy of the user instance and passes it to the function. The function then modifies the copy, not the original instance.
	*/
}
