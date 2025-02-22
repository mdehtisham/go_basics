package main

import (
	"fmt"
	mapUtil "main/basics"
)

/*
Imports
This code groups the imports into a parenthesized, "factored" import statement.

You can also write multiple import statements, like:

import "fmt"
import "math"
But it is good style to use the factored import statement.
*/

func main() {
	// fmt.Println("hello world")
	// valuesInGo()
	// forLoop()
	// mapUtil.MapExample()
	// mapUtil.CreateMixedMap()
	// mapUtil.PrintPerson(createPerson())
	// fmt.Println(mapUtil.Add(4, 7))
	// mapUtil.SwitchExample()
	mapUtil.WhenIsSaturday()
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
