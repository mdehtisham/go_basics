package main

import "fmt"

func main() {
	// fmt.Println("hello world")
	// valuesInGo()
	// forLoop()
	mapExample()
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

	// // Declare a map using map literal
	m2 := map[string]int{
		"apple":  5,
		"orange": 10,
	}
	fmt.Println(m2)
	delete(m2, "apple")
	fmt.Println(m2)

}
