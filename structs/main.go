package main

import "fmt"

// Go team prefers composition over inheritance.
// A struct is a type that contains named fields.

type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new Person instance.
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// ChangeName changes the name of the Person.
// The * in Go is used for references. If we don't pass by reference, the value will not change.
func (p *Person) ChangeName(newName string) {
	fmt.Println("Address of copy:", &p.Name)
	p.Name = newName
}

func main() {
	// myPerson := Person{
	// 	Name: "Aman",
	// 	Age:  20, // if we don't give value to age, it will default to 0
	// }

	myPerson := NewPerson("Billu", 20)
	myPerson.ChangeName("Ak Jain")

	// fmt.Println(myPerson)
	// fmt.Printf("This is my person %+v\n", myPerson) // +v is used for any type. It will print the value of the struct.

	// We can also use dot operator to access the value of the struct.
	// fmt.Println(myPerson.Name)
	// fmt.Println(myPerson.Age)

	// Changing value
	// myPerson.Name = "Aman Jain"

	fmt.Println(myPerson.Name)
	// fmt.Println("Address of allocated memory:", &myPerson.Name)
}
