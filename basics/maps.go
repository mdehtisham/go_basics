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
