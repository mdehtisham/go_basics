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

/*
You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/
//

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methodOnNonStruct() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
