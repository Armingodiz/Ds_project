package main

import (
	"fmt"
)

var cursorIndex int

func main() {
	cursorIndex = -1
	var requests int
	fmt.Scan(&requests)
	var input string
	fmt.Scan(&input)
	list := CreateLinkedList()
	for i := 0; i < len(input); i++ {
		list.AddFront(string(input[i]))
	}
	list.Display()
	var inp1 string
	var inp2 string
	for i := 0; i < requests; i++ {
		fmt.Scan(&inp1)
		switch inp1 {
		case "<":
			if cursorIndex != list.Size-1 {
				cursorIndex += 1
			}
			break
		case ">":
			if cursorIndex != -1 {
				cursorIndex -= 1
			}
			break
		case "+":
			fmt.Scan(&inp2)
			break
		case "-":
			break
		case "?":
			break
		case "!":
			break

		}
	}
}
