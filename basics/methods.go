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

/*
Pointer receivers
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.


*/

type PointerReceiverVertex struct {
	X, Y float64
}

func (v PointerReceiverVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *PointerReceiverVertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func PointerReceiverVertexMain() {
	v := PointerReceiverVertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}


/*
Pointers and functions
Here we see the Abs and Scale methods rewritten as functions.

Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

(If you're not sure, continue to the next page.)