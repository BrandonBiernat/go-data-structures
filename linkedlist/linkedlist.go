package linkedlist

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	data T
	next *node[T]
}

type List[T any] struct {
	head *node[T]
}

// Create an empty new list
func New[T any]() *List[T] {
	return &List[T]{}
}

// Get length of list
func (list *List[T]) Length() int {
	temp := list.head
	count := 0
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

// Print contents of list
func (list *List[T]) Print() {
	temp := list.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next

		if temp == nil {
			fmt.Println("-------------------END OF LIST-------------------")
		} else {
			fmt.Println("-------------------------------------------------")
		}
	}
}

// Add new node to the end of the list
func (list *List[T]) Push(value T) {
	node := &node[T]{data: value}

	if list.head == nil {
		list.head = node
	} else {
		temp := list.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
}

// Remove last node in the list
func (list *List[T]) Pop() error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.next == nil {
		list.head = nil
	} else {
		temp := list.head
		for temp.next.next != nil {
			temp = temp.next
		}
		temp.next = nil
	}

	return nil
}

// Insert node into the list at the given index
func (list *List[T]) Insert(value T, index int) error {
	if index < 1 {
		return errors.New("invalid index")
	}

	if list.head == nil {
		list.head = &node[T]{data: value}
	} else if index == 1 {
		node := &node[T]{data: value, next: list.head}
		list.head = node
	} else {
		curr := list.head
		count := 1
		for count < index-1 && curr != nil {
			curr = curr.next
			count++
		}

		if curr == nil {
			return errors.New("invalid node location")
		}

		node := &node[T]{data: value, next: curr.next}
		curr.next = node
	}

	return nil
}
