package main

import (
  "fmt"
)

func main() {
  p := newPerson("Joe", "Doe")
  fmt.Println(p.firstName, p.lastName)

  p.printName()
}
