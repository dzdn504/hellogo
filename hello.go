// from The Go Programming Language - Pluralsight course

package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string)

func Greet(saluation Salutation, do Printer) {
	message, alternate := CreateMessage(saluation.name, saluation.greeting, "yo")
	do(message)
	do(alternate)
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string) {
	fmt.Println(s + custom)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting[1] + " " + name
	alternate = "HEY!" + " " + name
	return
}

func main() {
	var s = Salutation{"Bob", "Hello"}

	Greet(s, CreatePrintFunction("?"))
}
