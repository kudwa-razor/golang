package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

//output:
//{1000000000 2} - i.e 1e9 = 1 followed by 9 zeros.