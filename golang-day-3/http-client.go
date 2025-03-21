package main
import (
	"fmt"
	"net/http"
	"bufio"
)

func main(){
	resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	//printing response.
	fmt.Println("Response status:", resp.Status)

	//printing 1st 5 lines .
	scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
