package main

import (
	"fmt"
)

var input string = "Habi"

func sayBye(text string) {
	fmt.Println("Hello,", text)
}

func sayGreeting(text string) {
	fmt.Println("Good Morning,", text)
}

func main() {
	sayGreeting("everyone")
	sayBye(input)
}
