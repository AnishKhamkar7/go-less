package main

import (
	"fmt"
)

type List struct {
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func (l *List) New() *List {
	fmt.Printf("Initial Call")
	return &List{head: nil}
}

func (l *List) get(idx int) int {
	n := l.head

	for i := 0; i < idx; i++ {
		if n == nil {
			return -1
		}
		n = n.next
	}

	if n == nil {
		return -1
	}

	return n.val

}

func (l *List) AddAtTail(val int) {
	n := l.head
	newNode := &Node{val: val}

	if l.head == nil {
		l.head = newNode
		return
	}

	for n.next != nil {
		n = n.next
	}

	n.next = newNode

}

func (l *List) AddAtHead(val int) {
	newNode := &Node{val: val, next: l.head}

	l.head = newNode
}

func (l *List) AddAtIndex(idx int, val int) {
	n := l.head

	if idx == 0 {
		l.AddAtHead(val)
		return
	}

	for i := 0; i < idx-1; i++ {
		if n == nil {
			return
		}
		n = n.next
	}

	if n == nil {
		return
	}

	newNode := &Node{val: val, next: n.next}

	n.next = newNode

}

func (l *List) DeleteAtIndex(idx int, val int) {
	n := l.head

	if idx == 0 {
		l.head = l.head.next
	}

	for i := 0; i < idx-1; i++ {
		if n == nil {
			return
		}
		n = n.next
	}

	if n.next == nil {
		return
	}

	n.next = n.next.next
}

func main() {
	l := &List{}

	l.AddAtHead(10)
	l.AddAtTail(20)
	l.AddAtTail(30)

	fmt.Println(l.get(0)) // 10
	fmt.Println(l.get(1)) // 20
	fmt.Println(l.get(2)) // 30

	l.DeleteAtIndex(1, 0)

	fmt.Println(l.get(1)) // should now be 30
}