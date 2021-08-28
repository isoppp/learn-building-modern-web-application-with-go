package main

import (
	"fmt"
	"log"
	"sort"
)

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

	// receiver
	fmt.Printf("%v, %p\n", user.Name, &user.Name)
	user.printNameByPointer()
	user.printName()

	// map
	myMap := make(map[string]string)
	fmt.Println(myMap["foo"])
	myMap["foo"] = "bar"
	fmt.Println(myMap["foo"])

	// slice
	var mySlice []string
	mySlice = append(mySlice, "slice v2")
	mySlice = append(mySlice, "slice v1")
	mySlice = append(mySlice, "slice v5")
	log.Println(mySlice)
	log.Println(mySlice[1:3]) // got by index 1, 2, len = 2

	sort.Strings(mySlice)
	log.Println(mySlice)

	fmt.Printf("%v, %p\n", mySlice[0], &mySlice[0]) // referenced to the same memory address

	for _, v := range mySlice {
		// if use v, copied v
		fmt.Printf("%v, %p\n", v, &v) // referenced to the same memory address
	}

	for i2 := range mySlice {
		// if access by using index access the original variable memory address
		fmt.Printf("%v, %p\n", mySlice[i2], &mySlice[i2]) // referenced to the same memory address
	}
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

func (u *User) printNameByPointer() {
	fmt.Printf("%v, %p\n", u.Name, &u.Name) // referenced to the same memory address
}

func (u User) printName() {
	fmt.Printf("%v, %p\n", u.Name, &u.Name) // referenced to the copied u
}
