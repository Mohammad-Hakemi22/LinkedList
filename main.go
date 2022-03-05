package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) InsertNode(data int) {
	n := Node{}
	n.data = data
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) InsertNodeAt(data, position int) {
	n := Node{}
	n.data = data
	if position < 0 || position > l.len {
		fmt.Println("Position is not valid!")
		return
	}
	if l.len == 0 {
		fmt.Println("Linked list is empty!")
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < position-1; i++ {
		ptr = ptr.next
	}
	n.next = ptr.next
	ptr.next = &n
	l.len++
}

func (l *LinkedList) Search(data int) (*Node, int, error) {
	if l.len == 0 {
		return nil, 0, errors.New("Linked List is empty")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.data == data {
			return ptr, i, nil
		}
		ptr = ptr.next
	}
	return nil, 0, errors.New(fmt.Sprintf("Not exist Node with data = %d", data))
}

func (l *LinkedList) DeleteNodeAt(position int) error {
	if position < 0 {
		return errors.New("Position is not valid!")
	}
	if l.len == 0 {
		return errors.New("Linked List is empty")
	}
	ptr := l.head
	for i := 0; i < position-2; i++ {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
	l.len--
	return nil
}

func (l *LinkedList) DeleteNodeByValue(data int) {
	_, pos, _ := l.Search(data)
	l.DeleteNodeAt(pos + 1)
}

func (l *LinkedList) PrintList() {
	if l.len == 0 {
		fmt.Println("Linked list is empty!")
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d. data: %d, next:%v\n", i+1, ptr.data, ptr.next)
		ptr = ptr.next
	}
}

func main() {
	l := LinkedList{}
	l.InsertNode(1)
	l.InsertNode(2)
	l.InsertNode(3)
	l.InsertNode(4)
	l.InsertNode(5)
	l.InsertNodeAt(10, 1)
	// n, err := l.Search(9)
	l.DeleteNodeByValue(3)
	l.PrintList()
}
