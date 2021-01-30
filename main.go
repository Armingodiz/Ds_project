package main

func main() {
	link := CreateLinkedList()
	link.AddFront("(")
	link.AddFront("2")
	link.AddFront("+")
	link.AddFront("3")
	link.AddFront(")")
	link.AddAfterIndex(-1, "5")
	link.Display()
}
