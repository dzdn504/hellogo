package main

import (
	"github.com/dzdn504/hellogo/greeting"
)

func main() {
	// var s = greeting.Salutation{"Bob", "Hello"}

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Mary", "What's goood!?"},
		{"Joe", "Yo dawg"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 5)
}
