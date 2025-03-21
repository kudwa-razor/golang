//12/17 - here printf is used instead of println - ***
// for printf we should use the %v for var int,float,bool and %q for string.
//println doesn't give anything for string - dk why doubt **

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s)
	fmt.Println(i, f, b, s)
}
