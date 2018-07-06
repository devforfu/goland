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

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x + y
	}
	return x
}

func main() {
	person := Person{"John", "Doe", 21}
	fmt.Printf("%v\n", person.Describe())
	println(fib(10))
}