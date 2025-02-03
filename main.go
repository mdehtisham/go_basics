package main

import (
	"fmt"
	mapUtil "main/basics"
)

func main() {
	// fmt.Println("hello world")
	// valuesInGo()
	forLoop()
	// mapUtil.MapExample()
	// mapUtil.CreateMixedMap()
	// mapUtil.PrintPerson(createPerson())
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
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("index is : ", i)
	// }
	// range for loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("range", n)
	}
}

func createPerson() mapUtil.Person {
	var p1 mapUtil.Person
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p1.Name)
	fmt.Print("Enter Age: ")
	fmt.Scanln(&p1.Age)
	fmt.Print("Enter Gender: ")
	fmt.Scanln(&p1.Gender)
	return p1
}
