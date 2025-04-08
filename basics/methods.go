/*
Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.
*/
package mapUtil

import (
	"fmt"
	"math"
)

type MethodsVertex struct {
	X, Y float64
}

func (v MethodsVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Methods are functions
Remember: a method is just a function with a receiver argument.

Here's Abs written as a regular function with no change in functionality..
*/
func useAbs() {
	v := MethodsVertex{3, 4}
	// Struct type method
	abs := v.Abs()
	fmt.Println(abs)
}
