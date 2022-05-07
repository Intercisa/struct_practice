package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	zip   int
}

func newPerson(firstName, lastName string) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func newPersonWithContact(firstName string, lastName string, email string, zip int) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
		contact: contact{
			email: email,
			zip:   zip,
		},
	}
}

func (p person) printName() {
	fmt.Printf("This person's name is: %s %s\n", p.firstName, p.lastName)
}

func (p person) printWithContact() {
	p.printName()
	fmt.Printf("Contact info: %+v\n", p.contact)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
