package main

import "fmt"

func main() {
	// basics
	var whatToSay string // declare with zero value
	var i int            // declare with zero value

	whatToSay = "i is set to"
	i = 2

	fmt.Println(whatToSay, i)

	whatWasSaid, otherThing := saySomething()
	fmt.Println("The function returned", whatWasSaid, otherThing)

	// pointer
	var pointerTest = "aaa"
	fmt.Println(pointerTest)
	changeUsingPointer1(&pointerTest) // &variable is pointer
	fmt.Println(pointerTest)

	// struct
	user := User{Name: "foo", Age: 10, phoneNumber: "bar"}
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.phoneNumber)
}

func saySomething() (string, string) {
	return "say something", "2"
}

// *string / *s is pointer
func changeUsingPointer1(s *string) {
	newValue := "Red"
	*s = newValue
}

type User struct {
	Name        string // public field
	Age         int
	phoneNumber string // kind of private(only can use in same package, variable, function and other names is almost the same)
}
