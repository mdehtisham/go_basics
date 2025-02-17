package mapUtil

import (
	"fmt"
	"runtime"
)

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

const Pi = 3.14

func getConstant() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/*
Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
*/

func SwitchExample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
