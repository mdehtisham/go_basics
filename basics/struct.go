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
