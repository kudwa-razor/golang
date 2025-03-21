package main
import "fmt"

func zeroval(val int){
	val = 0
}

func zeroptr(valptr *int){
	*valptr = 0
}

func main(){
	i :=1
	fmt.Println("Initial: ", i)

	zeroval(i) //the entire value itself is passed, but the value chaged, doesnt effect the actual value.
	fmt.Println("After zeroval", i)

	zeroptr(&i) //&i syntax gives the memory address of i, i.e. a pointer to i.
	fmt.Println("After pointer func", i)

	fmt.Println("Getting the memory location ig:", &i)
}
/* output:
1
1
0
hexa-decimal location no ig /*