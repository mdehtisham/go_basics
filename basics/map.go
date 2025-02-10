/*
Packages
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
*/
package mapUtil

import "fmt"

/*
In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

pizza and pi do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

*/

/*
The var statement declares a list of variables; as in function argument lists, the type is last.

A var statement can be at package or function level. We see both in this example.
*/

func MapExample() {
	// Declare a map using make
	var identityMap = make(map[string]string)
	identityMap["name"] = "Ehtisham"
	identityMap["age"] = "30"
	identityMap["Student"] = "true"
	fmt.Println(identityMap)

	// Accessing elements
	fmt.Println("identityMap[\"name\"]:", identityMap["name"])

	// // Declare a map using map literal
	m2 := map[string]int{
		"apple":  5,
		"orange": 10,
	}
	fmt.Println(m2)
	// Deleting elements
	delete(m2, "apple")
	fmt.Println(m2)

	// checking if a key exists
	value, exist := m2["apple"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value not found")
	}

	// iterating over a map
	for key, value := range identityMap {
		fmt.Printf("key is %s and value is %s \n", key, value)
	}

}

func CreateMixedMap() {
	// Declare a map with string keys and interface{} values
	var mixedMap = make(map[string]interface{})

	// Adding elements of different types
	mixedMap["name"] = "John"
	mixedMap["age"] = 30
	mixedMap["isStudent"] = true
	mixedMap["grades"] = []int{90, 85, 88}

	// Accessing and printing elements
	// fmt.Println("Name:", mixedMap["name"])
	// fmt.Println("Age:", mixedMap["age"])
	// fmt.Println("Is Student:", mixedMap["isStudent"])
	// fmt.Println("Grades:", mixedMap["grades"])

	// Iterating over the map
	for key, value := range mixedMap {
		fmt.Printf("Key: %s, Value: %v and type is %T\n", key, value, value)
	}
}

/*
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
*/
func Add(x, y int) int {
	return x + y
}

/*
Multiple results
A function can return any number of results.

The swap function returns two strings.
*/
func swap(x, y string) (string, string) {
	return y, x
}
func UsingSwap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func VariablesDeclaration() {
	/*
		Variables with initializers
		A var declaration can include initializers, one per variable.
		If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	*/
	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)
}
