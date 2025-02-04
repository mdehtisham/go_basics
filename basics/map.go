/*
Packages
Every Go program is made up of packages.

Programs start running in package main.

This program is using the packages with import paths "fmt" and "math/rand".

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
*/
package mapUtil

import "fmt"

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
