package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(saluation []Salutation, do Printer, isFormal bool, times int) {
	for _, s := range saluation {
		message, alternate := CreateMessage(s.Name, s.Greeting)

		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr. "
	case name == "Joe", name == "Amy":
		prefix = "Dr. "
	case name == "Mary":
		prefix = "Mrs. "
	default:
		prefix = "Dude "
	}
	return
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

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + " " + name
	return
}
