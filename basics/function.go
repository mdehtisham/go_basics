package mapUtil

import (
	"fmt"
	"math"
)

/*
Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func FuncVal() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

/*
Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func adderMain() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

/*
Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
func fibMain() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
