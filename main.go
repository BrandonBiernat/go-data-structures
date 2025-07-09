package main

import (
	"fmt"
	"linkedlist/linkedlist"
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

func main() {
	customer1 := Customer{
		firstName:   "Brandon",
		lastName:    "Biernat",
		age:         1,
		email:       "biernat.brandon@yahoo.com",
		phoneNumber: "(769) 257-8744",
	}
	customer2 := Customer{
		firstName:   "Brandon",
		lastName:    "Biernat",
		age:         2,
		email:       "biernat.brandon@yahoo.com",
		phoneNumber: "(769) 257-8744",
	}
	customer3 := Customer{
		firstName:   "Brandon",
		lastName:    "Biernat",
		age:         3,
		email:       "biernat.brandon@yahoo.com",
		phoneNumber: "(769) 257-8744",
	}

	list := linkedlist.List[Customer]{}
	list.Push(customer1)
	list.Push(customer2)
	list.Print()

	list.Insert(customer3, 2)
	list.Print()
	message := fmt.Sprintf("Length: %d", list.Length())
	fmt.Println(message)
}
