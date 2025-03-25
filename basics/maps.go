/*
Maps
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor can keys be added.

The make function returns a map of the given type, initialized and ready for use.
*/
package mapUtil

import "fmt"

type GeoVertex struct {
	Lat, Long float64
}

var m map[string]GeoVertex

func mapExample() {
	m = make(map[string]GeoVertex)
	m["Bell Labs"] = GeoVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
Map literals
Map literals are like struct literals, but the keys are required.
*/
type mapVertex struct {
	Lat, Long float64
}

var mp = map[string]mapVertex{
	"Bell Labs": mapVertex{
		40.68433, -74.39967,
	},
	"Google": mapVertex{
		37.42202, -122.08408,
	},
}

/*
Map literals continued
If the top-level type is just a type name, you can omit it from the elements of the literal.
*/
type mapLitralVertex struct {
	Lat, Long float64
}

var mapLitVer = map[string]mapLitralVertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func printMapLitVer() {
	fmt.Println(mapLitVer)
}

/*
Mutating Maps
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]
*/
