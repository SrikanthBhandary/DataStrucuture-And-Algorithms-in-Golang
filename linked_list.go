package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	head *Node
}

var errNoNodesInList = errors.New("no nodes found")

func (l *LinkedList) Read() error {
	if l.head == nil {
		return errNoNodesInList
	}
	node := l.head
	for node != nil {
		fmt.Println("<", node.Data)
		node = node.Next
	}
	return nil
}

func (l *LinkedList) Delete(email string) error {
	if l.head == nil {
		return errNoNodesInList

	}
	node := l.head
	prevNode := node
	for node != nil {
		if node.Data == email {
			fmt.Println("item found")
			if node == l.head {
				l.head = node.Next
			} else {
				prevNode.Next = node.Next
			}
			break
		}
		prevNode = node
		node = node.Next
	}

	return nil

}

func main() {

	node1 := &Node{Data: "test1@test.com"}
	n := &Node{Data: "test@test.com", Next: node1}
	linkedList := LinkedList{head: n}
	err := linkedList.Read()
	fmt.Println(err)
	linkedList.Delete("test1@test.com")
	err = linkedList.Read()
	fmt.Println(err)
	linkedList.Delete("test@test.com")
	err = linkedList.Read()
	fmt.Println(err)
}
