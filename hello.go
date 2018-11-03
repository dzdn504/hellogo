package main

import (
	"fmt"

	"github.com/dzdn504/hellogo/greeting"
)

func RenameToFrog(r greeting.Renamable) {
	r.Rename("Frog")
}

func main() {
	// var s = greeting.Salutation{"Bob", "Hello"}

	// var s []int
	// s = make([]int, 3)
	// s[0] = 1
	// s[1] = 10
	// s[2] = 12

	// v := []int{1, 10, 12}

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Joe", "Yo dawg"},
		{"Mary", "What's goood!?"},
	}

	salutations[0].RenameWithoutPointer("John")
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

	salutations[0].Rename("John")
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

	salutations[0] = salutations[0].RenameWithoutPointerReturn("Johnny")
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

	RenameToFrog(&salutations[0])
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

	fmt.Fprintf(&salutations[0], "The count is %d", 10)
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

}
