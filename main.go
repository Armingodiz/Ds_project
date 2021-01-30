package main

import "fmt"

func main() {
	link := CreateLinkedList()
	link.AddFront("head")
	link.AddBack("new head")
	link.AddBack("new new head")
	link.AddFront("tail")
	link.Display()
	fmt.Println("######################################################3")
	link.Delete(10)
	link.Display()
}
