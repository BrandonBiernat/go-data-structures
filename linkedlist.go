package main

import (
	"strconv"
)

type Customer struct {
	firstName   string
	lastName    string
	age         int
	email       string
	phoneNumber string
}

func (customer Customer) String() string {
	return "Name:         " + customer.firstName + " " + customer.lastName +
		"\nAge:          " + strconv.Itoa(customer.age) +
		"\nEmail:        " + customer.email +
		"\nPhone Number: " + customer.phoneNumber
}

func newCustomer(
	firstName string,
	lastName string,
	age int,
	email string,
	phoneNumber string,
) Customer {
	return Customer{
		firstName:   firstName,
		lastName:    lastName,
		age:         age,
		email:       email,
		phoneNumber: phoneNumber,
	}
}

type Node struct {
	data Customer
	next *Node
}

func newNode(customer Customer, next *Node) *Node {
	return &Node{data: customer, next: next}
}

type List struct {
	headNode *Node
}

// Print contents of list
func (list *List) traverse() {
	temp := list.headNode
	for temp != nil {
		println(temp.data.String())
		temp = temp.next
	}
}

// Add new node to the end of the list
func (list *List) push(customer *Customer) {
	node := newNode(*customer, nil)

	// Handle empty list
	if list.headNode == nil {
		list.headNode = node
	} else {
		temp := list.headNode
		for temp != nil {
			temp = temp.next
			if temp.next == nil {
				temp.next = node
				break
			}
		}
	}
}

func main() {
	customer := newCustomer(
		"Brandon",
		"Biernat",
		22,
		"biernat.brandon@yahoo.com",
		"(769) 257-8744",
	)
	list := List{headNode: nil}
	list.push(&customer)
	list.traverse()
}
