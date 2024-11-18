package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

// insert after last node
func (list *LinkedList) insertAtBack(data string) {
	newNode := &Node{data: data, next: nil}

	//jika head blm ada tambahkan node baru sebagai head
	if list.head == nil {
		list.head = newNode
		return
	}

	//jika head sudah ada current node adalah headnya
	current := list.head

	//looping current next yang tidak nil
	for current.next != nil {
		current = current.next
	}

	//insert current next yang nil dengan newNode
	current.next = newNode
}

// insert first node
func (list *LinkedList) insertAtFront(data string) {
	//jika headnya nil insert new Node sebagai head
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return
	}

	//jika sdh ada headnya set node baru sebagai head
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

func (list *LinkedList) insertAfterValue(afterVal string, data string) {
	newNode := &Node{data: data, next: nil}
	current := list.head

	for current != nil {
		if current.data == afterVal {
			newNode.next = current.next
			current.next = newNode
			return
		}

		current = current.next
	}

	fmt.Println("tidak bisa insert node")
}

func main() {
	link := LinkedList{}
	link.insertAtFront("danawan")
	link.insertAtBack("bimantoro")
	link.insertAfterValue("bimantoro", "putri")

	current := link.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next

	}
}
