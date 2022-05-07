package main

func main() {
	joe := newPerson("Joe", "Doe")

	joe.printName()

	jane := newPersonWithContact("Jane", "Doe", "janedoe@mail,com", 1123414)

	jane.printWithContact()
	jane.updateFirstName("Lisa")
	jane.printWithContact()

}
