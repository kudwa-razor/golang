//1/17 and 2/17
/* 2/17- imports.go states that u can do multiple import packages under one bracket or write multiple
import statements.*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
}
