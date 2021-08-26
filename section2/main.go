package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string // declare with zero value
	var i int            // declare with zero value

	whatToSay = "i is set to"
	i = 2

	fmt.Println(whatToSay, i)

	whatWasSaid, otherThing := saySomething()
	fmt.Println("The function returned", whatWasSaid, otherThing)
}

func saySomething() (string, string) {
	return "say something", "2"
}
