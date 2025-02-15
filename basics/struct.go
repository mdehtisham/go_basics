package mapUtil

import "fmt"

/*
In Go, a struct is a composite data type that groups together variables under a single name. These variables, known as fields, can have different types. Structs are used to create more complex data types that represent real-world entities.
*/

type Person struct {
	Name   string
	Age    int
	Gender string
}

func PrintPerson(p Person) {
	fmt.Printf("Name: %s, Age: %d, Gender: %s\n", p.Name, p.Age, p.Gender)
}

/*

How do you create an instance of a struct?
Answer: You can create an instance of a struct in several ways:

Using field names:
p := Person{Name: "Alice", Age: 25}


Without field names (positional):
p := Person{"Alice", 25}

Using the new keyword (returns a pointer):
p := new(Person)
p.Name = "Alice"
p.Age = 25

*/

/*
Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

*/
