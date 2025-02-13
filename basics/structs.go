/*
What is a Struct?
A struct (short for "structure") is a user-defined type that groups together fields (variables) under a single name. Each field has a name and a type, and structs are used to represent complex data structures.

Why Use Structs?
Structs allow you to group related data together.

They are useful for representing real-world entities (e.g., a Person, Car, or Book).

Structs can have methods associated with them, making them similar to classes in other languages.

Defining a Struct
To define a struct, use the type and struct keywords.

type <StructName> struct {
    <Field1> <Type1>
    <Field2> <Type2>
    ...
}

*/

package mapUtil

import "fmt"

type Human struct {
	Name string
	Age  int
}

func createHuman() {
	p := Human{Name: "Alice", Age: 25}
	fmt.Println(p)

}
