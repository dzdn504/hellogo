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

	// fmt.Println("*********************************************")
	// salutations[0].RenameWithoutPointer("John")
	// salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)
	//
	// fmt.Println("*********************************************")
	// salutations[0].Rename("John")
	// salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)
	//
	// fmt.Println("*********************************************")
	// salutations[0] = salutations[0].RenameWithoutPointerReturn("Johnny")
	// salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)
	//
	// fmt.Println("*********************************************")
	// RenameToFrog(&salutations[0])
	// salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)

	// fmt.Println("*********************************************")
	// fmt.Fprintf(&salutations[0], "The count is %d", 10)
	//
	// done := make(chan bool, 2)
	// go func() {
	// 	salutations.Greet(greeting.CreatePrintFunction("<C>"), true, 5)
	// 	done <- true
	// 	time.Sleep(100 * time.Millisecond)
	// 	done <- true
	// 	fmt.Println("Done!")
	// }()

	// salutations.Greet(greeting.CreatePrintFunction("?"), true, 5)
	// <-done

	fmt.Println("*********************************************")
	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)

	// Call a goroutine that will fill the channel
	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}

		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
