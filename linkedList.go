package main

import (
	"fmt"
)

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

func (list *LinkedList) AddAfterIndex(index int, key string) {
	if index == -1 {
		list.AddBack(key)

	} else {
		start := list.Search(index)
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
	rooter := list.Head
	fmt.Print("list is :   ")
	for rooter != nil {
		fmt.Print(rooter.Key)
		fmt.Print("  ")
		rooter = rooter.Next
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
