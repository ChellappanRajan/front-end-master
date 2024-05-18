package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func changeName(person Person, newName string) Person {
	person.Name = newName
	return person
}

// this changenameV only exisit on our struct type
func (p *Person) changeNameV(newName string) {

	p.Name = newName
}

func main() {

	myPerson := Person{
		Name: "Chellappan",
	}

	myPerson.Name = "Chellappan V"

	fmt.Printf("%+v", myPerson)

	myPersonnew := newPerson("Chellappan", 30)

	//myPersonnew is new copy
	changeName(myPersonnew, "Chellappan V")

	myPerson.changeNameV("V")

	fmt.Printf("%+v", myPerson)

	fmt.Printf("%+v", myPersonnew)
}
