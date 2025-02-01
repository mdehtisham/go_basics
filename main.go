package main

import "fmt"

func main() {
	// fmt.Println("hello world")
	// valuesInGo()
	// forLoop()
	// mapExample()
	createMixedMap()
}

func valuesInGo() {
	fmt.Println("valuesInGoing live reload!")

	fmt.Println("go"+"lang", 3+7, "ji")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("index is : ", i)
	}
}

func mapExample() {
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

func createMixedMap() {
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
