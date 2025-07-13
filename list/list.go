package list

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
	tail *node[T]
}

// Create an empty new list
func New[T any]() *List[T] {
	return &List[T]{}
}

func NewFromList[T any](fromList List[T]) *List[T] {
	return &List[T]{
		head: fromList.head,
		tail: fromList.tail,
	}
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

// Add a new value to the end of the list
func (list *List[T]) Push(value T) {
	node := &node[T]{data: value}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}
}

// Remove last value in the list
func (list *List[T]) Pop() error {
	if list.head == nil {
		return errors.New("list is empty")
	}

	if list.head.next == nil {
		list.head = nil
		list.tail = nil
	} else {
		temp := list.head
		for temp.next.next != nil {
			temp = temp.next
		}
		temp.next = nil
		list.tail = temp
	}

	return nil
}

// Insert value into the list at the given index
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

// For every value execute the provided function
func (list *List[T]) ForEach(fn func(data T)) {
	for curr := list.head; curr != nil; curr = curr.next {
		fn(curr.data)
	}
}

// Concatenates two lists into one new list
func (list *List[T]) Concat(value List[T]) *List[T] {
	newList := New[T]()

	for curr := list.head; curr != nil; curr = curr.next {
		newList.Push(curr.data)
	}

	for curr := value.head; curr != nil; curr = curr.next {
		newList.Push(curr.data)
	}

	return newList
}

// For every value execute the provided function and return a value of the specified type
func Select[T any, R any](list *List[T], fn func(data T) R) *List[R] {
	newList := New[R]()

	for curr := list.head; curr != nil; curr = curr.next {
		newList.Push(fn(curr.data))
	}

	return newList
}
