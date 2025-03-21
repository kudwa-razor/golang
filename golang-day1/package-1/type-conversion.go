//13/17

package main

import (
	"fmt"
	"math"
)

func main() {
	//simple example first
	/*
		var i int =42
		var f float64 =float64(i)
		var u uint = uint(f)
	*/

	//line 12-14 can be simply typed as below:
	i := 42
	g := float64(i)
	u := uint(g)
	// so the code in the given drive states this:
	//Try removing the float64 or uint conversions in the example and see what happens. ---- no difference is seen idk.
	fmt.Println(i, g, u)

	//trying one more thing given in drive:

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

}
