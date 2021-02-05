package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var cursorIndex int
var requests int
var table *HashTable

func main() {
	cursorIndex = -1
	table = NewHashTable(10)
	fmt.Scan(&requests)
	var input string
	fmt.Scan(&input)
	list := CreateLinkedList()
	for i := 0; i < len(input); i++ {
		list.AddFront(string(input[i]))
	}
	//list.Display()
	start(list)
}

func start(list *LinkedList) {
	var inp1 string
	var inp2 string
	for i := 0; i < requests; i++ {
		//list.Display()
		fmt.Scan(&inp1)
		switch inp1 {
		case "<":
			if cursorIndex != -1 {
				cursorIndex -= 1
			}
			break
		case ">":
			if cursorIndex != list.Size {
				cursorIndex += 1
			}
			break
		case "+":
			fmt.Scan(&inp2)
			list.AddAfterIndex(inp2)
			if cursorIndex < list.Size {
				cursorIndex += 1
			}
			break
		case "-":
			list.DeleteBeforIndex()
			if cursorIndex > -1 {
				cursorIndex -= 1
			}
			break
		case "?":
			list.Display()
			break
		case "!":
			listString := list.getString()
			//fmt.Println(listString)
			tableResult := table.Search(listString)
			if tableResult != nil {
				fmt.Println(tableResult.Value)
			} else {
				//fmt.Println("calculating ... ")
				newList := InfToPost(list)
				res, _ := calculatePost(newList)
				table.Insert(listString, res)
				fmt.Println(res)
			}
			break
		}
	}
}

// a+b*(c^d-e)^(f+g*h)-i
// 1+2*(3^4-5)^(6+7*8)-9
// 1234^5-678*+^*+9-
// 1234^5-678*+^*+9-
// 14+2+3*5-9
func GetPre(c string) int {
	if strings.Compare(c, "^") == 0 {
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
		num := current.Key
		if err == nil {
			for err == nil && current.Next != nil {
				_, err = strconv.Atoi(current.Next.Key)
				if err == nil {
					num += current.Next.Key
					current = current.Next
					if current == nil {
						break
					}
				}
			}
			result.AddFront(num)
		} else if strings.Compare(current.Key, "(") == 0 {
			stack.Push(current.Key)
		} else if strings.Compare(current.Key, ")") == 0 {
			for ; !stack.IsEmpty() && strings.Compare(stack.Peak(), "(") != 0; {
				result.AddFront(stack.Pop())
			}
			if strings.Compare(stack.Peak(), "(") == 0 {
				stack.Pop()
			}
		} else {
			for ; !stack.IsEmpty() && GetPre(current.Key) <= GetPre(stack.Peak()); {
				result.AddFront(stack.Pop())
			}
			stack.Push(current.Key)
		}
		current = current.Next
	}
	for ; !stack.IsEmpty(); {
		result.AddFront(stack.Pop())
	}
	return result
}

func calculatePost(list *LinkedList) (int, error) {
	stack := CreateStack()
	// Scan all characters one by one
	current := list.Head
	for current != nil {
		// If the scanned character is an operand (number here),
		// push it to the stack.
		_, err := strconv.Atoi(current.Key)
		if err == nil {
			stack.Push(current.Key)
		} else { // If the scanned character is an operator, pop two
			// elements from stack apply the operator   else
			val1, err1 := strconv.Atoi(stack.Pop())
			val2, err2 := strconv.Atoi(stack.Pop())
			if err1 == nil && err2 == nil {
				switch current.Key {
				case "+":
					stack.Push(strconv.Itoa(val2 + val1))
					break
				case "-":
					stack.Push(strconv.Itoa(val2 - val1))
					break
				case "*":
					stack.Push(strconv.Itoa(val2 * val1))
					break
				case "/":
					stack.Push(strconv.Itoa(val2 / val1))
					break
				}
			}
		}
		current = current.Next
	}
	return strconv.Atoi(stack.Pop())
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

func
CreateLinkedList() *LinkedList {
	list := LinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	return &list
}

func (list *LinkedList) getString() string {
	result := ""
	temp := list.Head
	for temp != nil {
		//fmt.Println("tsss")
		result += temp.Key
		temp = temp.Next
	}
	return result
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
	list.Size = list.Size + 1
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
		//fmt.Print("@")
		fmt.Print(rooter.Key)
		//fmt.Print("@")
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

// hash table implementation :

type HashNode struct {
	Key   string
	Value int
}

type HashTable struct {
	Size     int
	Capacity int
	Array    []*HashNode
	Dummy    *HashNode
}

func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		Size:     0,
		Capacity: capacity,
		Array:    make([]*HashNode, capacity),
		Dummy: &HashNode{
			Key:   "",
			Value: -1,
		},
	}
}

func NewHashNode(key string, value int) *HashNode {
	return &HashNode{
		Key:   key,
		Value: value,
	}
}

/*func hashCode(key string) int {
	p := 31
	m := 1000000009
	power_of_p := 1
	hash_val := 0
	for i := 0; i < len(key); i++ {
		hash_val = (hash_val + (int(key[i])-'a'+1)*power_of_p) % m
		power_of_p = (power_of_p * p) % m
	}
	return hash_val
}*/
func (table *HashTable) hashCode(key string) int {
	h := 0
	o := 31415
	t := 27183
	for i := 0; i < len(key); i++ {
		h = (o*h + int(key[i])) % table.Capacity
		o = o * t % (table.Capacity - 1)
	}
	return h
}

func (table *HashTable) hashCodePrime(hashIndex int) int {
	hashIndex += 1
	hashIndex %= table.Capacity
	return hashIndex
}

func (table *HashTable) Insert(key string, value int) error {
	if table.Size == table.Capacity {
		return errors.New("hash table is full !")
	}
	node := NewHashNode(key, value)
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil && table.Array[hashIndex] != table.Dummy {
		hashIndex = table.hashCodePrime(hashIndex)
	}
	if table.Array[hashIndex] == nil || table.Array[hashIndex] == table.Dummy {
		table.Array[hashIndex] = node
		table.Size = table.Size + 1
		return nil
	} else {
		return errors.New("error in finding empty slot !!")
	}
}

func (table *HashTable) Search(key string) *HashNode {
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil {
		if table.Array[hashIndex].Key == key {
			return table.Array[hashIndex]
		}
		hashIndex = table.hashCodePrime(hashIndex)
	}
	return nil
}

func (table *HashTable) Delete(key string) error {
	hashIndex := table.hashCode(key)
	for table.Array[hashIndex] != nil {
		if table.Array[hashIndex].Key == key {
			table.Array[hashIndex] = table.Dummy
			return nil
		}
		hashIndex = table.hashCodePrime(hashIndex)
	}
	return errors.New("key not found ! ")
}

func (table *HashTable) Display() {
	for i := 0; i < table.Capacity; i++ {
		fmt.Print("for slot ")
		fmt.Print(i)
		fmt.Print(" : ")
		fmt.Println(table.Array[i])
		fmt.Println("###########################################")
	}
}
