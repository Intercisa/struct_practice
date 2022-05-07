package main

import "fmt"

type person struct {
  firstName string
  lastName string
}

func newPerson(firstName, lastName string ) person {
  return person{
    firstName: firstName,
    lastName: lastName,
  }
}

func (p person) printName() {
  fmt.Printf("This person's name is: %s %s", p.firstName, p.lastName)
}

