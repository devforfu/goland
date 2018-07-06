package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

type Printable interface {
	Describe() string
}

func (p Person) Describe() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}


func main() {
	person := Person{"John", "Doe", 21}
	fmt.Printf("%v\n", person.Describe())
}