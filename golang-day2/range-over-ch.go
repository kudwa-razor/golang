package main
import "fmt"

func main(){
	queue := make(chan string, 2)
	queue <- "one" // sending value
	queue <- "two" //sending value
	close(queue)

	// here, range iterates over each element as it is received from queue.
	// since the queue is closed above, range ends iteration after 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
	
}