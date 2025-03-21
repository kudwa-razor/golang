package main
import (
	"fmt"
	"time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {
	f("direct")

	go f("gouroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
    fmt.Println("done")
}
// the output from go f ("goroutine") and go func(msg string) will not be in order as both  are concurrently run on go runtime.