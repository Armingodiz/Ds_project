package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cursorIndex int
var requests int

func main() {
	cursorIndex = -1
	fmt.Scan(&requests)
	var input string
	fmt.Scan(&input)
	list := CreateLinkedList()
	for i := 0; i < len(input); i++ {
		list.AddFront(string(input[i]))
	}
	list.Display()
	start(list)
}

func start(list *LinkedList) {
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
			list.AddAfterIndex(inp2)
			cursorIndex += 1
			break
		case "-":
			list.DeleteBeforIndex()
			cursorIndex -= 1
			break
		case "?":
			list.Display()
			break
		case "!":
			break
		}
	}
}
func GetPre(c string) int {
	if strings.EqualFold("^", c) {
		return 3
	} else if strings.EqualFold(c, "*") || strings.EqualFold(c, "/") {
		return 2
	} else if strings.EqualFold(c, "+") || strings.EqualFold(c, "-") {
		return 1
	} else {
		return -1
	}
}

func InfToPost(list *LinkedList) *LinkedList {
	stack := CreateStack()
	result := CreateLinkedList()
	current := list.Head
	for ; current != nil; {
		_, err := strconv.Atoi(current.Key)
		if err == nil {
			result.AddFront(current.Key)
		} else if strings.EqualFold(current.Key,"(") {
			stack.Push(current.Key)
		} else if strings.EqualFold(current.Key,")") {
			for ;!stack.IsEmpty() && strings.EqualFold(stack.Peak(),"("); {
				result.AddFront(stack.Pop())
			}
			if strings.EqualFold(stack.Peak(),"(") {
				stack.Pop()
			}
		} else{
			for;!stack.IsEmpty() && GetPre(current.Key) <= GetPre(stack.Peak());{
				result.AddFront(stack.Pop())
			}
			stack.Push(current.Key)
		}
		current = current.Next
	}
	for ;!stack.IsEmpty();{
		result.AddFront(stack.Pop())
	}
	return result
}

/// stack implementation :

type Stack struct {
	stack []string
	top   int
}

func CreateStack() *Stack {
	stack := &Stack{
		stack: []string{},
		top:   -1,
	}
	return stack
}

func (stack *Stack) IsEmpty() bool {
	if stack.top == -1 {
		fmt.Println("STACK IS EMPTY !")
		return true
	}
	return false
}

func (stack *Stack) Push(element string) {
	stack.stack = append(stack.stack, element)
	stack.top += 1
}

func (stack *Stack) Pop() string {
	if stack.IsEmpty() {
		return ""
	} else {
		popped := stack.stack[stack.top]
		stack.top -= 1
		stack.stack = stack.stack[:stack.top+1]
		return popped
	}
}

func (stack *Stack) Peak() string {
	if stack.IsEmpty() {
		return ""
	} else {
		return stack.stack[stack.top]
	}
}

/// linked list implementation :

type Node struct {
	Key  string
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func CreateLinkedList() *LinkedList {
	list := LinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	return &list
}

func (list *LinkedList) AddFront(key string) {
	if list.Size == 0 {
		list.firstInitialize(key)
	} else if list.Size == 1 {
		node := &Node{
			Key:  key,
			Prev: list.Head,
			Next: nil,
		}
		list.Head.Next = node
		list.Tail = node
	} else {
		node := &Node{
			Key:  key,
			Prev: list.Tail,
			Next: nil,
		}
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size = list.Size + 1
}

func (list *LinkedList) AddAfterIndex(key string) {
	if cursorIndex == -1 {
		list.AddBack(key)

	} else {
		start := list.Search(cursorIndex)
		node := &Node{
			Key:  key,
			Prev: nil,
			Next: list.Head,
		}
		node.Next = start.Next
		start.Next = node
		node.Prev = start
	}
}

func (list *LinkedList) DeleteBeforIndex() {
	if cursorIndex != -1 {
		node := list.Search(cursorIndex)
		if node.Next == nil && node.Prev == nil { // node is the only node
			list.Head = nil
		} else if node.Next == nil { //node is tail of list
			list.Tail = node.Prev
			node.Prev.Next = nil
		} else if node.Prev == nil { //nod is head of list
			list.Head = node.Next
			node.Next.Prev = nil
		} else {
			node.Next.Prev = node.Prev
			node.Prev.Next = node.Next
		}
		list.Size = list.Size - 1
	}
}

func (list *LinkedList) AddBack(key string) {
	if list.Size == 0 {
		list.firstInitialize(key)
	} else if list.Size == 1 {
		node := &Node{
			Key:  key,
			Prev: nil,
			Next: list.Head,
		}
		list.Head.Prev = node
		list.Tail = list.Head
		list.Head = node
	} else {
		node := &Node{
			Key:  key,
			Prev: nil,
			Next: list.Head,
		}
		list.Head.Prev = node
		list.Head = node
	}
	list.Size = list.Size + 1
}

// adding first element to linked list
func (list *LinkedList) firstInitialize(key string) {
	node := &Node{
		Key:  key,
		Prev: nil,
		Next: nil,
	}
	list.Head = node
}

func (list *LinkedList) Delete(index int) *Node {
	node := list.Search(index)
	if node == nil {
		return nil
	} else {
		if node.Next == nil && node.Prev == nil { // node is the only node
			list.Head = nil
		} else if node.Next == nil { //node is tail of list
			list.Tail = node.Prev
			node.Prev.Next = nil
		} else if node.Prev == nil { //nod is head of list
			list.Head = node.Next
			node.Next.Prev = nil
		} else {
			node.Next.Prev = node.Prev
			node.Prev.Next = node.Next
		}
		list.Size = list.Size - 1
		return node
	}
}

func (list *LinkedList) Display() {
	counter := 0
	printed := false
	if cursorIndex == -1 {
		fmt.Printf("|")
		printed = true
	}
	rooter := list.Head
	for rooter != nil {
		fmt.Print(rooter.Key)
		if counter == cursorIndex && !printed {
			fmt.Printf("|")
		}
		rooter = rooter.Next
		counter += 1
	}
	fmt.Println()
}

func (list *LinkedList) Search(index int) *Node {
	if index < 0 || index > list.Size {
		fmt.Println("INVALID INDEX !")
		return nil
	} else {
		node := list.Head
		for i := 0; i < index; i++ {
			node = node.Next
		}
		return node
	}
}
