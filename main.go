package main

import (
	"list/list"
	"strconv"
)

type Customer struct {
	firstName   string
	lastName    string
	age         int
	email       string
	phoneNumber string
}

type Test struct {
	data string
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
		firstName:   "Molly",
		lastName:    "Kirby",
		age:         2,
		email:       "Kirby.Molly@yahoo.com",
		phoneNumber: "(769) 257-8744",
	}
	customer3 := Customer{
		firstName:   "William",
		lastName:    "Wu",
		age:         3,
		email:       "Wu.William@yahoo.com",
		phoneNumber: "(769) 257-8744",
	}

	customerList := list.New[Customer]()
	customerList.Push(customer1)
	customerList.Push(customer2)
	customerList.Push(customer3)

	customerList2 := list.New[Customer]()
	customerList2.Push(customer1)
	customerList2.Push(customer2)
	customerList2.Push(customer3)

	newList := customerList.Concat(*customerList2)

	newList.Print()
}
