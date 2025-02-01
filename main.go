package main

import (
	"fmt"
	mapUtil "main/basics"
)

func main() {
	// fmt.Println("hello world")
	// valuesInGo()
	// forLoop()
	mapUtil.MapExample()
	mapUtil.CreateMixedMap()
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
